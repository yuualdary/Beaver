package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// fmt.Println(custompkg.Is(5, 1))

	// fmt.Printf("Right now is %s\n", carbon.Now().DateTimeString())
	// today, _ := carbon.NowInLocation("Asia/Jakarta") // _ = berfungsi sebagai penampung data yang ingin balikin function yang dipakai cmn gamau dipakai, misal name, _= "john","Doe"  fmt.Println(name) outputny jadi john, karena Doe sudah dimasukkan kedalam tempat tak terpakai dan tempat tak terpakai harus dibelakang ","
	// fmt.Printf("Right now in Jakarta is %s\n", today)

	// fmt.Print(rand.Intn(100))

	// a := 100
	// b := 200

	// var test bool = a > b

	// fmt.Println(test)

	// var intA int32 = 128
	// var intB int8 = int8(intA)

	// fmt.Println(intA)
	// fmt.Println(intB)

	// var x float64 = 26.9

	// y := int(x)
	// fmt.Println(y)
	// fmt.Printf("y data type : %T\n", y) //untuk cek data type

	// age := "dua puluh enam"

	// result, err := strconv.Atoi(age)

	// if err != nil {

	// 	log.Fatal(err)
	// }

	// for i := 1; i <= 1000000; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println(result)

	test := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan Teks : ")

	value, _ := test.ReadString('\n')
	fmt.Println("Hallo  " + value)

}
