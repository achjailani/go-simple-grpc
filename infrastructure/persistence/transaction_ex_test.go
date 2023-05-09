package persistence_test

import (
	"database/sql"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/tests"
	"strconv"
	"strings"
	"testing"
	"time"
)

// NOTE:
// There 4 isolation levels which are (READ UNCOMMITTED, READ COMMITTED, REPEATABLE READ, SERIALIZABLE)
// Postgres has 3 (READ UNCOMMITTED and READ COMMITTED are equal)

func Test_Gorm_ReadCommitted_Or_UnCommitted_ReadOp(t *testing.T) {
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

func Test_Gorm_RepeatableRead_WriteOp(t *testing.T) {
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

	time.Sleep(3 * time.Second)

	err := tx.Commit().Error
	assert.NoError(t, err)
}
