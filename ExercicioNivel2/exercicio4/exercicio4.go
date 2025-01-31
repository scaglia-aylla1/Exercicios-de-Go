package main

import "fmt"

// Crie um programa que:
// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal


func main(){
	
	num := 200
	fmt.Printf("%d\t, %b\t, %#x\n", num, num, num)
	y := num << 1
	fmt.Printf("%d\t, %b\t, %#x", y, y, y)

}