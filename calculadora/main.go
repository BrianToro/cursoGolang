package main

import (
	"strconv"
	"strings"
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Inserte la operacion (numero operador numero) Ejem: 2 + 2")
	scanner.Scan()
	operacion := scanner.Text()
	valores := separar(operacion)
	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[2])
	if err1 != nil || err2 != nil {
		controlDeErrores()
	} else {
		resultado := definirOperacion(operador1, operador2, valores[1])
		fmt.Println(resultado)
	}
}

func controlDeErrores(){
	fmt.Println("Operacion invalida")
}

func definirOperacion(num1 int, num2 int, operador string)(int){
	var resultado int
	switch operador {
	case "+":
		resultado = num1 + num2
		break
	case "-":
		resultado = num1 - num2
		break
	case "*", "x":
		resultado = num1 * num2
		break
	case "/":
		resultado = num1 / num2
		break
	default:
		resultado = num1 + num2
	}
	return resultado
}

func separar(cadena string)([]string){
	valores :=  strings.Split(cadena, " ")
	return valores
}
