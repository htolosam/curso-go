package main

import "fmt"

func main() {
	// ===== IF / ELSE =====
	fmt.Println("=== IF / ELSE ===")
	edad := 18

	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}

	// If con inicialización
	if numero := 10; numero%2 == 0 {
		fmt.Printf("%d es par\n", numero)
	} else {
		fmt.Printf("%d es impar\n", numero)
	}

	// If anidado
	nota := 85
	if nota >= 90 {
		fmt.Println("Calificación: A")
	} else if nota >= 80 {
		fmt.Println("Calificación: B")
	} else if nota >= 70 {
		fmt.Println("Calificación: C")
	} else {
		fmt.Println("Calificación: D")
	}

	// ===== SWITCH =====
	fmt.Println("\n=== SWITCH ===")
	dia := 3

	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	default:
		fmt.Println("Fin de semana")
	}

	// Switch con múltiples valores
	mes := "Enero"
	switch mes {
	case "Diciembre", "Enero", "Febrero":
		fmt.Println("Invierno")
	case "Marzo", "Abril", "Mayo":
		fmt.Println("Primavera")
	case "Junio", "Julio", "Agosto":
		fmt.Println("Verano")
	case "Septiembre", "Octubre", "Noviembre":
		fmt.Println("Otoño")
	}

	// Switch sin condición (como if-else)
	temperatura := 25
	switch {
	case temperatura < 0:
		fmt.Println("Hace mucho frío")
	case temperatura < 15:
		fmt.Println("Hace frío")
	case temperatura < 25:
		fmt.Println("Temperatura agradable")
	default:
		fmt.Println("Hace calor")
	}

	// ===== FOR (el único bucle en Go) =====
	fmt.Println("\n=== FOR ===")

	// For clásico
	fmt.Println("Contador 1-5:")
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For como while
	fmt.Println("\nCuenta regresiva:")
	contador := 5
	for contador > 0 {
		fmt.Printf("%d... ", contador)
		contador--
	}
	fmt.Println("¡Despegue!")

	// For infinito (con break)
	fmt.Println("\nBúsqueda de número:")
	numero := 0
	for {
		numero++
		if numero == 7 {
			fmt.Printf("Encontré el 7 en la iteración %d\n", numero)
			break
		}
	}

	// For con continue
	fmt.Println("\nNúmeros pares del 1 al 10:")
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue // Salta los impares
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// For con range (para colecciones)
	fmt.Println("\nIterando sobre un slice:")
	numeros := []int{10, 20, 30, 40, 50}
	for indice, valor := range numeros {
		fmt.Printf("Índice: %d, Valor: %d\n", indice, valor)
	}

	// Range solo con valores
	fmt.Println("\nSolo valores:")
	for _, valor := range numeros {
		fmt.Printf("%d ", valor)
	}
	fmt.Println()

	// Range con string
	fmt.Println("\nIterando sobre string:")
	texto := "Hola"
	for i, char := range texto {
		fmt.Printf("Posición %d: %c\n", i, char)
	}

	// ===== DEFER =====
	fmt.Println("\n=== DEFER ===")
	fmt.Println("Inicio")
	defer fmt.Println("Esto se ejecuta al final (defer 1)")
	defer fmt.Println("Esto se ejecuta antes (defer 2)")
	defer fmt.Println("Esto se ejecuta primero (defer 3)")
	fmt.Println("Fin")
	// Los defer se ejecutan en orden LIFO (Last In, First Out)
}
