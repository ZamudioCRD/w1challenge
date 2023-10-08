package main

import (
	"fmt"
	"os"
	"w1challenge/evenodd"
	"w1challenge/ohm"
	"w1challenge/foobar"
	"w1challenge/bmi"
	"w1challenge/mario"
)

func main() {
	args := os.Args
	functionName := args[1]
	switch functionName{
	case "evenodd":
		var num int
		fmt.Println("Numero a evaluar")
		fmt.Scan(&num)
		fmt.Println(evenodd.EvenOrOdd(num))
	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Ingresa valores de v, r e i")
		fmt.Scan(&v,&r,&i)
		fmt.Println(ohm.Ohm(v,r,i))
	case "foobar":
		var num int
		fmt.Println("Ingresa numero deseado")
		fmt.Scan(&num)
		fmt.Println(foobar.Foobar(num))
	case "bmi":
		var peso, altura float32
		fmt.Println("¿Cuánto pesas? (no mientas)")
		fmt.Scan(&peso)
		fmt.Println("¿Cuánto mides? (sin zapatos)")
		fmt.Scan(&altura)
		var imc float32 = bmi.CalcBmi(peso, altura)
		fmt.Println("Right now your BMI is", imc)
		var mensaje string = bmi.BmiMensaje(imc)
		fmt.Println(mensaje)
	case "mario":
		var altura int
		for {
			fmt.Print("Pyramid height: ")
			var num int
			fmt.Scan(&num)
			if altura < 1 || altura > 8 {
				continue
			} else {
				break
			}
		}	
		mario.Piramide(altura)
	}
}
