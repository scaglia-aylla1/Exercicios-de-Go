package main

import "fmt"
//Escreva um programa que mostre um número em decimal, binário, e hexadecimal.


func main(){
	x := 10
	fmt.Printf("%d, %#x, %b", x, x, x)

}
//solução => %d representa o numero em decimal
//%#x => representa o numero em hexadecimal
//%b => representa o numero binário