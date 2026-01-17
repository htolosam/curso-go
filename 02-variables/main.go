package main

import "fmt"

func main() {
	// Declaración de variables con var
	var nombre string = "Juan"
	var edad int = 30
	var altura float64 = 1.75
	var activo bool = true

	fmt.Println("=== Declaración con var ===")
	fmt.Printf("Nombre: %s, Edad: %d, Altura: %.2f, Activo: %t\n", nombre, edad, altura, activo)

	// Declaración corta con :=
	ciudad := "Madrid"
	temperatura := 25.5
	fmt.Println("\n=== Declaración corta := ===")
	fmt.Printf("Ciudad: %s, Temperatura: %.1f°C\n", ciudad, temperatura)

	// Múltiples variables
	var (
		pais      = "España"
		poblacion = 47000000
	)
	fmt.Println("\n=== Múltiples variables ===")
	fmt.Printf("País: %s, Población: %d\n", pais, poblacion)

	// Constantes
	const PI = 3.14159
	const NOMBRE_APP = "Curso Go"
	fmt.Println("\n=== Constantes ===")
	fmt.Printf("PI: %.5f, App: %s\n", PI, NOMBRE_APP)

	// Tipos de datos básicos
	fmt.Println("\n=== Tipos de datos ===")
	var entero int = 42
	var entero8 int8 = 127
	var entero64 int64 = 9223372036854775807
	var flotante32 float32 = 3.14
	var flotante64 float64 = 3.141592653589793
	var complejo complex64 = 1 + 2i
	var booleano bool = true
	var caracter rune = 'A' // rune es un alias de int32
	var cadena string = "Hola"

	fmt.Printf("int: %d\n", entero)
	fmt.Printf("int8: %d\n", entero8)
	fmt.Printf("int64: %d\n", entero64)
	fmt.Printf("float32: %.2f\n", flotante32)
	fmt.Printf("float64: %.15f\n", flotante64)
	fmt.Printf("complex64: %v\n", complejo)
	fmt.Printf("bool: %t\n", booleano)
	fmt.Printf("rune: %c (código: %d)\n", caracter, caracter)
	fmt.Printf("string: %s\n", cadena)

	// Zero values (valores por defecto)
	fmt.Println("\n=== Zero values ===")
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("int: %d, float64: %f, bool: %t, string: '%s'\n", i, f, b, s)
}
