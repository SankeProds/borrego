package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"text/template"
)

func GetTransDataFromFile(filepath string) (Transaction, error) {
	var transData Transaction
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return transData, err
	}

	err = json.Unmarshal(file, &transData)
	if err != nil {
		return transData, err
	}
	return transData, nil
}

func createTransHandler(transData Transaction, tmplPath string, handlerFileFormat string) error {
	file, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		fmt.Printf("Error loading Template: %v\n", err)
		return err
	}
	tmpl, err := template.New("Trans").Parse(string(file))
	if err != nil {
		return err
	}
	fout, err := os.Create(fmt.Sprintf(handlerFileFormat, transData.Name))
	if err != nil {
		return err
	}
	defer fout.Close()
	err = tmpl.Execute(fout, transData)
	if err != nil {
		return err
	}
	return nil
}

func GenHandler(transactionFilePath string) {
	transData, err := GetTransDataFromFile(transactionFilePath)
	if err != nil {
		fmt.Printf("Error loading Transaction %s: %v\n", transactionFilePath, err)
		os.Exit(1)
	}

	// LOADS TRANSACTION HANDLER TEMPLATE AND RENDERS IT WITH LOADED 
	// TRANS DATA
	err = createTransHandler(transData, "generators/handler.tmpl", "handlers/handler%s.go")
	if err != nil {
		fmt.Printf("Error creating handler for Transaction %s: %v\n", transactionFilePath, err)
		os.Exit(1)
	}
}
