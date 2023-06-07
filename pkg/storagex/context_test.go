package storagex_test

import (
	"fmt"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"github/achjailani/go-simple-grpc/pkg/storagex"
	"github/achjailani/go-simple-grpc/tests"
	"os"
	"strings"
	"testing"
)

func TestNewStorageX(t *testing.T) {
	textbox := tests.Init()
	cfg := textbox.Cfg
	ctx := textbox.Ctx

	minio, err := storagex.NewMinio(cfg)

	f := faker.New()

	t.Run("it should not error connect to minio", func(t *testing.T) {
		assert.NoError(t, err)
		assert.NotNil(t, minio)
	})

	strx := storagex.NewStorageX(minio)

	t.Run("it should not be nil storagex", func(t *testing.T) {
		assert.NotNil(t, strx)
	})

	pwd, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/%s", getFilePath(pwd), "pdfx/file-0001.pdf")

	file, err := os.Open(filePath)

	t.Run("It should not be error open file", func(t *testing.T) {
		assert.NoError(t, err)
	})

	t.Run("it should be valid store file", func(t *testing.T) {
		err = strx.Store(ctx, &storagex.FileUploadObject{
			File:     file,
			FileName: f.UUID().V4(),
		})

		assert.NoError(t, err)
	})
}

func getFilePath(pwd string) string {
	fls := strings.Split(pwd, "/")
	return strings.Join(fls[:len(fls)-1], "/")
}
