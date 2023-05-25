package pdfx_test

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestConvertDocToPDF(t *testing.T) {
	pwd, _ := os.Getwd()
	docPath := fmt.Sprintf("%s/%s", pwd, "file-0001.docx")
	pdfPath := fmt.Sprintf("%s/%s", pwd, "file-0001.pdf")

	if _, err := os.Stat(pdfPath); err == nil {
		fmt.Println("File already exists")
		return
	}

	cmd := exec.Command("soffice", "--headless", "--convert-to", "pdf", pdfPath, docPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conversion completed successfully!")
}

func TestOverridePDFMetadata(t *testing.T) {
	pwd, _ := os.Getwd()
	inPdfPath := fmt.Sprintf("%s/%s", pwd, "file-0001.pdf")
	outPdfPath := fmt.Sprintf("%s/%s", pwd, "file-0001-overrided.pdf")

	// check if file exist
	if _, err := os.Stat(inPdfPath); err != nil {
		fmt.Println("File not found")
		return
	}

	// remove existing out pdf file
	if _, err := os.Stat(outPdfPath); err == nil {
		_ = os.Remove(outPdfPath)
	}

	// read file
	inPdf, err := pdfcpu.ReadFile(inPdfPath, model.NewDefaultConfiguration())
	if err != nil {
		log.Fatal(err)
	}

	inPdf.Title = "PDF Overriding Metadata"
	inPdf.Author = "John Snow"
	inPdf.Subject = "Just Simple Simulation"
	inPdf.Keywords = "Pdf Operation, Privy"

	// Write the modified PDF file with overridden metadata.
	err = api.WriteContextFile(inPdf, outPdfPath)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PDF metadata overridden successfully!")
}
