package main
import "fmt"

// Crie constantes tipadas e não-tipadas.
//Demonstre seus valores.
func main()  {
	const idade int = 32 //tipada
	const name = "Ana"       // não tipada
	
	fmt.Println(name)
	fmt.Println(idade)
}