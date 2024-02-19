package db

import (
	"bytes"
	mdl "distr-calc/model"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

var exprs = map[string]mdl.Expression{
	"1a8dbfc9-d089-4544-a2e9-8093c9012fd9": {Uuid: "1a8dbfc9-d089-4544-a2e9-8093c9012fd9", Status: "Ready", Value: "2+2*2", Result: "6"},
	//"d4295216-be86-409f-b0d8-4be26301fa15": {Uuid: "d4295216-be86-409f-b0d8-4be26301fa15", Status: "Calculating", Value: "2+2*2", Result: "?"},
	"c7c41838-f305-41e4-8fff-d10c56e48d92": {Uuid: "c7c41838-f305-41e4-8fff-d10c56e48d92", Status: "Error", Value: "2+2*2", Result: "error"},
}

func InitExpressions() {
	exs := map[string]mdl.Expression{}

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

	err = json.Unmarshal(byteValue, &exs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(exs)
	exprs = exs
}

func UpdateExpression(expr mdl.Expression) {
	save(expr)
}

func GetExpressions() map[string][]mdl.Expression {
	es := maps.Values(exprs)
	r := map[string][]mdl.Expression{
		"Expressions": es,
	}
	return r
}

func SaveExpressions(value string, status string, result string) mdl.Expression {
	id := uuid.New().String()

	expression := mdl.Expression{Uuid: id, Status: status, Value: value, Result: result}

	save(expression)

	return expression
}

func save(expr mdl.Expression) {
	exprs[expr.Uuid] = expr

	bn, err := writeDataToFileAsJSON(exprs, "expressions.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DC printed ", bn, " bytes to ", "expressions.json")
}

func writeDataToFileAsJSON(data interface{}, filedir string) (int, error) {
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
