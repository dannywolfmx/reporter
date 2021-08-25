package reporter

import (
	"bytes"
	"fmt"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

type PDF interface{}

type pdfReporter struct{
	moroto pdf.Maroto 

}

type HeadersParams struct{
	MinimumDiscount string
	InvoiceNumber string
	Name, LastName, SecondLastName string
	ActualDate string
}

func NewPDFReporter() *pdfReporter{
	moroto := pdf.NewMaroto(consts.Portrait, consts.A4)
	return &pdfReporter{
		moroto: moroto,	
	}
}

func (p *pdfReporter) Generate(headersParams HeadersParams, images []string, sumaryData, calculationData [][]string) (bytes.Buffer, error){
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 15, 10)
	//m.SetBorder(true)

	p.header(headersParams)
	m.Row(4, func() {})

	p.summaryTable(sumaryData)
	m.Row(4, func() {})

	p.calculationTable(calculationData)
	m.Row(4, func() {})

	p.printCharts(images)

	return p.moroto.Output() 
}

func (p *pdfReporter) printCharts(images []string) {
	for _, image := range images {
		p.moroto.Row(10, func() {})
		p.moroto.Row(90, func() {
			p.moroto.Col(12, func() {
				p.moroto.Base64Image(image, consts.Png, props.Rect{
					Left:   5,
					Top:    5,
					Center: true,
				})
			})
		})
	}

}

func (p *pdfReporter) header(params HeadersParams){
	p.moroto.RegisterHeader(func() {
		p.moroto.SetBackgroundColor(darPurpleColor)
		p.moroto.Row(12, func() {
			p.moroto.Col(12, func() {
				p.moroto.Text("Estrategia de liquidación de deuda ", props.Text{
					Color:  color.NewWhite(),
					Top:    4,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   13,
					Align:  consts.Center,
				})
			})
		})

		p.moroto.Row(12, func() {
			p.moroto.ColSpace(1)

			p.moroto.Col(4, func() {
				p.moroto.Text(fmt.Sprintf("Folio: %s", params.InvoiceNumber), props.Text{
					Color:  color.NewWhite(),
					Top:    4,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   13,
					Align:  consts.Left,
				})
			})

			p.moroto.Col(7, func() {
				descuento := fmt.Sprintf("Descuento minimo esperado: %s", params.MinimumDiscount)
				p.moroto.Text(descuento, props.Text{
					Color:  color.NewWhite(),
					Top:    4,
					Style:  consts.Bold,
					Family: consts.Courier,
					Size:   13,
					Align:  consts.Left,
				})
			})

		})

		p.moroto.SetBackgroundColor(color.NewWhite())
		p.moroto.Row(11, func() {
			p.moroto.Col(9, func() {
				p.moroto.Text(fmt.Sprintf("%s %s %s", params.Name, params.LastName, params.SecondLastName), props.Text{
					Top:   2,
					Size:  11,
					Style: consts.Bold,
					Align: consts.Center,
				})
			})

			p.moroto.Col(3, func() {
				p.moroto.Text(params.ActualDate, props.Text{
					Top:   2,
					Size:  11,
					Style: consts.Bold,
					Align: consts.Center,
				})
			})
		})
	})
}


func SummaryTable(m pdf.Maroto, data [][]string) {
	headers := []string{
		"Numero de cuentas",
		"Ahorro Mensual",
		"Meses en el programa",
		"Porcentaje de descuento",
	}

	m.TableList(headers, data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{3, 3, 3, 3},
			Family:    consts.Courier,
			Style:     consts.Bold,
		},
		ContentProp: props.TableListContent{
			Size:      11,
			GridSizes: []uint{3, 3, 3, 3},
			Family:    consts.Courier,
			Style:     consts.Bold,
		},
		Align:              consts.Center,
		HeaderContentSpace: 1,
		Line:               false,
	})
}

func CalculationTable(p pdf.Maroto, data [][]string) {
	headers := []string{
		"",
		"Banco",
		"Numero de cuenta",
		"Deuda inicial",
		"Pago al banco",
		"Comisión",
		"Mes de Liquidación",
	}

	p.SetBackgroundColor(tealColor)
	p.Row(10, func() {
		p.Col(12, func() {
			p.Text("Transacciones", props.Text{
				Top:    2,
				Size:   13,
				Style:  consts.Bold,
				Align:  consts.Center,
				Family: consts.Courier,
				Color:  color.NewWhite(),
			})
		})
	})
	p.SetBackgroundColor(color.NewWhite())

	p.TableList(headers, data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 2, 2, 2, 2, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 2, 2, 2, 2, 2, 1},
		},
		Align:              consts.Center,
		HeaderContentSpace: 1,
		Line:               false,
		AlternatedBackground: &color.Color{
			Red:   210,
			Green: 200,
			Blue:  260,
		},
	})
}

func (p *pdfReporter) summaryTable(data [][]string) {
	headers := []string{
		"Numero de cuentas",
		"Ahorro Mensual",
		"Meses en el programa",
		"Porcentaje de descuento",
	}

	p.moroto.TableList(headers, data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      10,
			GridSizes: []uint{3, 3, 3, 3},
			Family:    consts.Courier,
			Style:     consts.Bold,
		},
		ContentProp: props.TableListContent{
			Size:      11,
			GridSizes: []uint{3, 3, 3, 3},
			Family:    consts.Courier,
			Style:     consts.Bold,
		},
		Align:              consts.Center,
		HeaderContentSpace: 1,
		Line:               false,
	})
}

func (p *pdfReporter) calculationTable(data [][]string) {
	headers := []string{
		"",
		"Banco",
		"Numero de cuenta",
		"Deuda inicial",
		"Pago al banco",
		"Comisión",
		"Mes de Liquidación",
	}

	p.moroto.SetBackgroundColor(tealColor)
	p.moroto.Row(10, func() {
		p.moroto.Col(12, func() {
			p.moroto.Text("Transacciones", props.Text{
				Top:    2,
				Size:   13,
				Style:  consts.Bold,
				Align:  consts.Center,
				Family: consts.Courier,
				Color:  color.NewWhite(),
			})
		})
	})
	p.moroto.SetBackgroundColor(color.NewWhite())

	p.moroto.TableList(headers, data, props.TableList{
		HeaderProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 2, 2, 2, 2, 2, 1},
		},
		ContentProp: props.TableListContent{
			Size:      8,
			GridSizes: []uint{1, 2, 2, 2, 2, 2, 1},
		},
		Align:              consts.Center,
		HeaderContentSpace: 1,
		Line:               false,
		AlternatedBackground: &color.Color{
			Red:   210,
			Green: 200,
			Blue:  260,
		},
	})
}