package main

import (
	"text/template"
	"os"
	"fmt"
	"io/ioutil"
)

type MainData struct {
	Transactions []Transaction
}

func createMain(transactions []Transaction, tmplPath string, mainPath string) error{
	file, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		fmt.Printf("Error loading Template: %v\n", err)
		return err
	}
	tmpl, err := template.New("Main").Parse(string(file))
	if err != nil {
		return err
	}
	fout, err := os.Create(mainPath)
	if err != nil {
		return err
	}
	defer fout.Close()
	err = tmpl.Execute(fout, MainData{transactions})
	if err != nil {
		return err
	}
	return nil
}

func GenMain(transactionFolder string){
	files, _ := ioutil.ReadDir(transactionFolder)
	var transactions []Transaction
	for _, f := range files {
		t, err := GetTransDataFromFile(fmt.Sprintf("%s/%s", transactionFolder, f.Name()))
		if err != nil {
			fmt.Printf("ERROR: %v\n", err)
			continue
		}
		transactions = append(transactions, t)
	}
	createMain(transactions, "generators/main.tmpl", "borre.go")
}
