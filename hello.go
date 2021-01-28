package main

import (
	"fmt"
	"hello/operations"

	"github.com/iancoleman/strcase"
	"github.com/leekchan/accounting"
)

func main() {

	ac := accounting.Accounting{Symbol: "Rp", Precision: 3}
	fmt.Println(ac.FormatMoney(123456789.213123)) // "$123,456,789.21"
	fmt.Println("Hello World")
	fmt.Println(operations.Addition(4, 8))
	fmt.Println(operations.Substraction(4, 8)) //kalau mau panggil package harus huruf besar gabisa huruf kecil
	var message = "AnyKind of_string"
	fmt.Println(strcase.ToCamel(message))
}
