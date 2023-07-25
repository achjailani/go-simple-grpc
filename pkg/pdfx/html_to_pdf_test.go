package pdfx_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestHtmlToPdf(t *testing.T) {
	pwd, _ := os.Getwd()
	inFile := fmt.Sprintf("%s/%s", pwd, "html-file.html")
	outFile := fmt.Sprintf("%s/%s", pwd, "pdf-file.pdf")

	if _, err := os.Stat(outFile); err == nil {
		fmt.Println("File already exists")
		return
	}

	cmd := exec.Command("wkhtmltopdf", inFile, outFile)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	assert.True(t, true)
}
