package reporter

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {

	reporter := NewPDFReporter()

	headers := HeadersParams{}

	fileContent, err := ioutil.ReadFile("test.png")
	if err != nil{
		t.Fatalf("Error al abrir archivo imagen de prueba %s", err)
	}


	images := []string{
		base64.StdEncoding.EncodeToString(fileContent),
		base64.StdEncoding.EncodeToString(fileContent),
		base64.StdEncoding.EncodeToString(fileContent),
	}

	sumaryData := [][]string{
		{"hdhd", "3","3", "3"},
	}

	calculationData := [][]string{
		{"hdhd", "3","3", "3", "2", "1", "3"},
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
