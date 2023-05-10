package persistence_test

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/tests"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"testing"
	"time"
)

// NOTE:
// There 4 isolation levels which are (READ UNCOMMITTED, READ COMMITTED, REPEATABLE READ, SERIALIZABLE)
// Postgres has 3 (READ UNCOMMITTED and READ COMMITTED are equal)

func Test_IsoLevel_ReadCommitted_Or_UnCommitted_ReadOp(t *testing.T) {
	box := tests.Init()
	ctx := box.Ctx

	f := faker.New()

	db := box.Repo.DB
	db2 := box.Repo.DB

	tx := db.Begin(&sql.TxOptions{
		Isolation: sql.LevelReadUncommitted, // equal to LevelReadCommitted
	})

	const (
		id = 1
	)

	t.Run("It should return 0: number of users before inserting", func(t *testing.T) {
		var number int64
		err := tx.WithContext(ctx).
			Model(&entity.User{}).
			Count(&number).
			Error

		assert.NoError(t, err)
		assert.Equal(t, int64(0), number)
	})

	t.Run("it should valid inserting data", func(t *testing.T) {
		err := tx.WithContext(ctx).Create(&entity.User{
			ID:       id,
			Name:     f.Person().Name(),
			Username: strings.Join(strings.Split(strings.ToLower(f.Person().Name()), " "), "_"),
			Password: strings.Join(append(strings.Split(f.Person().Name(), " "), strconv.Itoa(f.RandomNumber(3))), ""),
		}).Error

		assert.NoError(t, err)
	})

	t.Run("It should return 0 from another transaction", func(t *testing.T) {
		var number int64
		err := db2.WithContext(ctx).
			Model(&entity.User{}).
			Count(&number).
			Error

		assert.NoError(t, err)
		assert.Equal(t, int64(0), number)
	})

	t.Run("It should return 1 from main transaction", func(t *testing.T) {
		var number int64
		err := tx.WithContext(ctx).
			Model(&entity.User{}).
			Count(&number).
			Error

		assert.NoError(t, err)
		assert.Equal(t, int64(1), number)
	})

	err := tx.Commit().Error
	assert.NoError(t, err)
}

func Test_IsoLevel_RepeatableRead_WriteOp(t *testing.T) {
	box := tests.Init()
	ctx := box.Ctx

	f := faker.New()

	db := box.Repo.DB
	db2 := box.Repo.DB

	tx := db.Begin(&sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	const (
		id       = 1
		name     = "Sansa Stark"
		username = "sansastart"
		password = "sansastark124"
	)

	var numberOfUsers int64
	t.Run("It should return 0: number of users before inserting", func(t *testing.T) {
		err := tx.WithContext(ctx).
			Model(&entity.User{}).
			Count(&numberOfUsers).
			Error

		assert.NoError(t, err)
		assert.Equal(t, int64(0), numberOfUsers)
	})

	t.Run("it should valid inserting data", func(t *testing.T) {
		err := tx.WithContext(ctx).Create(&entity.User{
			ID:       id,
			Name:     name,
			Username: username,
			Password: password,
		}).Error

		assert.NoError(t, err)
	})

	t.Run("It should return 1: number of users after inserting", func(t *testing.T) {
		err := tx.WithContext(ctx).
			Model(&entity.User{}).
			Count(&numberOfUsers).
			Error

		assert.NoError(t, err)
		assert.Equal(t, int64(1), numberOfUsers)
	})

	t.Run("It should no error when updating, but no effect", func(t *testing.T) {
		err := db2.WithContext(ctx).
			Model(&entity.User{}).
			Where("id = ?", id).
			Update("name", f.Person().Name()).
			Error

		assert.NoError(t, err)

		var usr entity.User
		err = tx.WithContext(ctx).
			Model(&entity.User{}).
			Where("id = ?", id).
			Take(&usr).
			Error

		assert.NoError(t, err)
		assert.Equal(t, name, usr.Name)
	})

	t.Run("It should be valid inserting http logs: from another transaction", func(t *testing.T) {
		err := db2.WithContext(ctx).
			Create(&entity.HttpLog{
				Ip:     f.Internet().Ipv4(),
				Path:   f.Internet().User(),
				Method: f.Internet().HTTPMethod(),
			}).Error

		assert.NoError(t, err)
	})

	time.Sleep(3 * time.Second)

	err := tx.Commit().Error
	assert.NoError(t, err)
}

func Test_IsoLevel_Serializable_WriteOp(t *testing.T) {
	box := tests.Init()
	ctx := box.Ctx

	db := box.Repo.DB
	db2 := box.Repo.DB

	ids := []uint{1, 2, 3, 4, 5}

	tx := db.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	tx2 := db2.Begin(&sql.TxOptions{
		Isolation: sql.LevelSerializable,
	})

	var (
		latestName     = "Jacop"
		latestUsername = "jacop123"

		latestName2     = "Jackpad"
		latestUsername2 = "jackpad123"
	)

	t.Run("it should be valid running seed", func(t *testing.T) {
		err := seedTxUsers(ids, db)

		assert.NoError(t, err)
	})

	t.Run("it should be valid", func(t *testing.T) {
		var r []*entity.User

		err := tx.WithContext(ctx).Find(&r, ids).Error

		assert.NoError(t, err)
		assert.NotEmpty(t, r)
		assert.Equal(t, len(ids), len(r))
	})

	t.Run("it should be valid when updating from another transaction", func(t *testing.T) {
		err := tx2.WithContext(ctx).Where("id = ?", ids[1]).Updates(&entity.User{
			Name:     latestName,
			Username: latestUsername,
		}).Error

		assert.NoError(t, err)

		err = tx2.Commit().Error

		assert.NoError(t, err)
	})

	t.Run("it should return err, updating from main transaction", func(t *testing.T) {
		err := tx.WithContext(ctx).Where("id = ?", ids[4]).Updates(&entity.User{
			Name:     latestName2,
			Username: latestUsername2,
		}).Error

		assert.Error(t, err)
	})

	t.Run("it should not be valid commit", func(t *testing.T) {
		err := tx.Commit().Error

		assert.Error(t, err)
	})
}

func seedTxUsers(ids []uint, db *gorm.DB) error {

	f := faker.New()
	var preps []*entity.User

	for idx, id := range ids {
		preps = append(preps, &entity.User{
			ID:       id,
			Name:     f.Person().Name(),
			Username: strings.Join(strings.Split(strings.ToLower(f.Person().Name()), " "), "_"),
			Password: strings.Join(append(strings.Split(f.Person().Name(), " "), strconv.Itoa(f.RandomNumber(3))), ""),
		})

		err := db.WithContext(context.Background()).Create(preps[idx]).Error
		if err != nil {
			return err
		}
	}

	for _, r := range preps {
		fmt.Printf("ID: %d, NAME: %s, USERNAME: %s\n", r.ID, r.Name, r.Username)
	}

	return nil
}
