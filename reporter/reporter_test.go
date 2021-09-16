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
	if err != nil {
		t.Fatalf("Error al abrir archivo imagen de prueba %s", err)
	}

	images := []string{
		base64.StdEncoding.EncodeToString(fileContent),
		base64.StdEncoding.EncodeToString(fileContent),
		base64.StdEncoding.EncodeToString(fileContent),
	}

	sumaryData := [][]string{
		{"2", "$6,182.43", "18", "25.81%", "0.05%", "$19,913.01"},
	}

	dataRow := []string{
		"1", "BANCO", "123123123", "$50,000.00", "$28,0000.00", "$2,1001.81", "6",
	}

	calculationData := [][]string{}

	for i := 0; i <= 15; i++ {
		calculationData = append(calculationData, dataRow)
	}

	buffer, err := reporter.Generate(headers, images, sumaryData, calculationData)

	if err != nil {
		t.Fatalf("Error al generar el buffer %s", err)
	}

	f, err := os.Create("prueba.pdf")

	defer func() {
		err = f.Close()
		if err != nil {
			t.Fatalf("Error al cerrar el archivo %s", err)
		}
	}()

	//Error handler of os.Create
	if err != nil {
		t.Fatalf("Error al crear el archivo %s", err)
	}

	f.Write(buffer.Bytes())

}
