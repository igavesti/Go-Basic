package main

import "fmt"

func main() {
	var nilai32 int32 = 123
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

	nilai32 = 1
	fmt.Println(nilai32)
	fmt.Println(nilai64)

	fmt.Print(string("Budi"[1]))
}
