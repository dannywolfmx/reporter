package reporter

import (
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {

	reporter := NewPDFReporter()

	headers := HeadersParams{}

	images := []string{}

	sumaryData := [][]string{
		[]string{"hdhd", "3","3", "3"},
	}

	calculationData := [][]string{
		[]string{"hdhd", "3","3", "3", "2", "1", "3", "1"},
	}

	buffer, err := reporter.Generate(headers, images, sumaryData, calculationData)

	if err != nil{
		t.Fatalf("Error al generar el buffer %s", err)
	}


	f, err := os.Create("prueba.pdf")

	defer func(){
		err = f.Close()
		if err != nil{
			t.Fatalf("Error al cerrar el archivo %s", err)
		}
	}()

	//Error handler of os.Create
	if err != nil{
		t.Fatalf("Error al crear el archivo %s", err)
	}
	
	f.Write(buffer.Bytes())

}
