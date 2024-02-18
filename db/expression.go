package db

import (
	"bytes"
	mdl "distr-calc/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

var expressions = map[string][]mdl.Expression{
	"Expressions": {
		{Uuid: "1", Status: "Ready", Value: "2+2*2", Result: "6"},
		{Uuid: "2", Status: "Calculating", Value: "2+2*2", Result: "?"},
		{Uuid: "3", Status: "Error", Value: "2+2*2", Result: "error"},
	},
}

func InitExpressions() {
	es := []mdl.Expression{}

	jsonFile, err := os.Open("expressions.json")
	if err != nil {
		fmt.Printf("Open file error: %s\n", err)
		return
	}

	fmt.Println("Successfully Opened expressions.json")
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("ReadAll error: %s\n", err)
		return
	}

	err = json.Unmarshal(byteValue, &es)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(es)
	expressions["Expressions"] = es
}

func GetExpressions() map[string][]mdl.Expression {
	return expressions
}

func SaveExpressions(e mdl.Expression) {
	es := expressions["Expressions"]
	es = append(es, e)
	expressions["Expressions"] = es

	n, err := WriteDataToFileAsJSON(expressions["Expressions"], "expressions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DC printed ", n, " bytes to ", "expressions.json")
}

func WriteDataToFileAsJSON(data interface{}, filedir string) (int, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return 0, err
	}
	file, err := os.OpenFile(filedir, os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return 0, err
	}
	n, err := file.Write(buffer.Bytes())
	if err != nil {
		return 0, err
	}

	return n, nil
}
