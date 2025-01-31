package main
import "fmt"

//Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
// == igual     != diferente igual    <= menor igual 
// < imenor     >= maior igual         > maior
//Demonstre estes valores. 

func main(){
	a := (10 == 100)
	b := (10 != 100)
	c := (10 <= 100)
	d := (10 < 100)
	e := (10 >= 100)
	f := (10 > 100)

	fmt.Println(a, b, c, d, e, f)	

}