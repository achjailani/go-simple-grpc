package pdfx

import (
	"github.com/unidoc/unipdf/v3/model"
	"log"
)

func Extract() {
	// Specify the path to the input PDF file
	inputFile := "input.pdf"

	// Open the input PDF
	pdfReader, err, _ := model.NewPdfReaderFromFile(inputFile, &model.ReaderOpts{
		Password: "",
		LazyLoad: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = pdfReader
	// Get the signature fields from the PDF

}
