package main

import (
	"fmt"

	"./crypto"
)

/*
var (
	language = "Go"
	company  = "google"
	year     = 2022
)*/

func main() {
	//var language = "go"
	fmt.Printf("Aprendiendo %s con %s en el año %d", language, company, year)
	crypto.ToBitcoin(10000)
	crypto.ToEthereum(10000)
	crypto.ToLitecoin(10000)
}
