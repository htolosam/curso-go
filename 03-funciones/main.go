package main

import "fmt"

// Función simple sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola desde una función!")
}

// Función con parámetros
func saludarPersona(nombre string) {
	fmt.Printf("¡Hola, %s!\n", nombre)
}

// Función con retorno
func sumar(a int, b int) int {
	return a + b
}

// Parámetros del mismo tipo (sintaxis corta)
func multiplicar(a, b int) int {
	return a * b
}

// Múltiples valores de retorno
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return a / b, nil
}

// Retorno nombrado
func operaciones(a, b int) (suma, resta int) {
	suma = a + b
	resta = a - b
	return // retorno implícito
}

// Funciones variádicas (número variable de argumentos)
func sumarTodos(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}

// Función que retorna una función (closure)
func multiplicador(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Función recursiva
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("=== Funciones básicas ===")
	saludar()
	saludarPersona("María")

	fmt.Println("\n=== Funciones con retorno ===")
	resultado := sumar(5, 3)
	fmt.Printf("5 + 3 = %d\n", resultado)
	fmt.Printf("4 * 6 = %d\n", multiplicar(4, 6))

	fmt.Println("\n=== Múltiples retornos ===")
	cociente, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", cociente)
	}

	_, err2 := dividir(10, 0)
	if err2 != nil {
		fmt.Println("Error esperado:", err2)
	}

	fmt.Println("\n=== Retornos nombrados ===")
	s, r := operaciones(10, 3)
	fmt.Printf("suma: %d, resta: %d\n", s, r)

	fmt.Println("\n=== Funciones variádicas ===")
	fmt.Printf("Suma de 1, 2, 3: %d\n", sumarTodos(1, 2, 3))
	fmt.Printf("Suma de 1, 2, 3, 4, 5: %d\n", sumarTodos(1, 2, 3, 4, 5))

	numeros := []int{10, 20, 30}
	fmt.Printf("Suma de slice: %d\n", sumarTodos(numeros...))

	fmt.Println("\n=== Closures ===")
	duplicar := multiplicador(2)
	triplicar := multiplicador(3)
	fmt.Printf("Duplicar 5: %d\n", duplicar(5))
	fmt.Printf("Triplicar 5: %d\n", triplicar(5))

	fmt.Println("\n=== Recursión ===")
	fmt.Printf("Factorial de 5: %d\n", factorial(5))
	fmt.Printf("Factorial de 7: %d\n", factorial(7))
}
