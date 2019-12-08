package main

import "fmt"
import "unsafe"

func type_test1() {
	a := 100
	b := 35.465

	fmt.Println("Size of int :", unsafe.Sizeof(a))
	fmt.Println("Size of float :", unsafe.Sizeof(b))
}

func type_test2() {
	var b byte = 0x9C
	var r1 rune = '민'
	var r2 rune = '\uAD7F'

	fmt.Println("Size of byte :", unsafe.Sizeof(b))
	fmt.Println("Size of rune :", unsafe.Sizeof(r1), unsafe.Sizeof(r2))
}

func type_test3() {
	var com64 complex64 = 5.54 + 2.71i
	var com128 complex128 = 7.01 + 3.4512e-10i
	var com64_2 complex64 = complex(5.54, 2.71)			// complex 함수를 이용해 선언 가능
	var com128_2 complex128 = complex(7.01, 3.4512e-10)	// complex 함수를 이용해 선언 가능

	com64_real := real(com64)
	com64_imag := imag(com64)
	com128_real := real(com128)
	com128_imag := imag(com128)

	fmt.Println("Complex64 :", unsafe.Sizeof(com64), unsafe.Sizeof(com64_2))
	fmt.Println("Complex128 :", unsafe.Sizeof(com128), unsafe.Sizeof(com128_2))
	fmt.Println("Complex64 real, imag :", unsafe.Sizeof(com64_real), unsafe.Sizeof(com64_imag))
	fmt.Println("Complex128 real, imag :", unsafe.Sizeof(com128_real), unsafe.Sizeof(com128_imag))
}

func main() {
	type_test1()
	type_test2()
	type_test3()
}
