package main

import "fmt"

func main() {

	students := [...]string{"Yudistiro", "Wahyu", "Aldary"}

	for _, student := range students { //kenapa dikosongin "_" karena nampung variable yang ga kepake

		fmt.Println(student)
	}
}
