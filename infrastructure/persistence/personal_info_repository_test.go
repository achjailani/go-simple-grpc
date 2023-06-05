package persistence_test

import (
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/domain/entity"
	"github/achjailani/go-simple-grpc/pkg/cryptox"
	"github/achjailani/go-simple-grpc/tests"
	"testing"
)

func TestPersonalInfoRepo_Create(t *testing.T) {
	testbox := tests.Init()
	ctx := testbox.Ctx

	repo := testbox.Repo

	f := faker.New()

	t.Run("it should be valid create", func(t *testing.T) {
		err := repo.PersonalInfo.Create(ctx, &entity.PersonalInfo{
			Name:  f.Person().Name(),
			Email: f.Internet().Email(),
			Phone: f.Phone().Number(),
		})

		assert.NoError(t, err)
	})
}

func TestPersonalInfoRepo_FindByEmail(t *testing.T) {
	testbox := tests.Init()
	ctx := testbox.Ctx

	repo := testbox.Repo

	f := faker.New()

	var (
		email = f.Internet().Email()
	)

	t.Run("it should be valid create", func(t *testing.T) {
		err := repo.PersonalInfo.Create(ctx, &entity.PersonalInfo{
			Name:  f.Person().Name(),
			Email: email,
			Phone: f.Phone().Number(),
		})

		assert.NoError(t, err)
	})

	t.Run("it should be valid find by email", func(t *testing.T) {
		blindIndex, _ := cryptox.MakeBlindIndex(email)

		r, err := repo.PersonalInfo.FindByEmail(ctx, blindIndex)

		assert.NoError(t, err)
		assert.NotNil(t, r)
		fmt.Println(r.String())
	})
}
