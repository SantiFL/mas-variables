package main

import "fmt"

func main() {
	var nombre string = " santiago"
	var apellido string = "ferreyra"
	var edad int = 22
	const dni float64 = 40.835432

	fmt.Println("Este es un programa echo en go mas completo que el Hola mundo")
    fmt.Println("mis datos son :" + nombre + apellido)
    fmt.Println(edad)
    fmt.Println(dni)
}
