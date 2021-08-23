package reporter

import (
	"bytes"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
)

type PDF interface{}

type pdfReporter struct{
	moroto pdf.Maroto 
}


func NewPDFReporter() *pdfReporter{
	moroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	return &pdfReporter{
		moroto: moroto,	
	}
}


func (p *pdfReporter) Generate() (bytes.Buffer, error){
	return p.moroto.Output() 
}