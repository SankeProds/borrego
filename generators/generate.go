package main

import (
	"os"
)

func main(){
	if os.Args[1] == "Handler" {
		GenHandler(os.Args[2]) // PATH TO TRANSACTION FILE
	}
	if os.Args[1] == "Main" {
		GenMain(os.Args[2]) // PATH TO TRANSACTIONS FOLDER
	}
}
