package main

import "fmt"

func main() {
	// ===== ARRAYS =====
	fmt.Println("=== ARRAYS ===")

	// Declaración de array (tamaño fijo)
	var numeros [5]int
	fmt.Println("Array vacío:", numeros) // [0 0 0 0 0]

	// Array con valores iniciales
	dias := [7]string{"Lun", "Mar", "Mié", "Jue", "Vie", "Sáb", "Dom"}
	fmt.Println("Días:", dias)

	// Array con tamaño inferido
	colores := [...]string{"Rojo", "Verde", "Azul"}
	fmt.Println("Colores:", colores)

	// Acceso a elementos
	fmt.Printf("Primer día: %s\n", dias[0])
	fmt.Printf("Último día: %s\n", dias[6])

	// Modificar elementos
	numeros[0] = 10
	numeros[1] = 20
	fmt.Println("Números modificados:", numeros)

	// Longitud del array
	fmt.Printf("Longitud del array: %d\n", len(dias))

	// ===== SLICES =====
	fmt.Println("\n=== SLICES ===")

	// Crear slice desde array
	todosLosDias := dias[:]
	fmt.Println("Slice de todos los días:", todosLosDias)

	// Slice de una porción
	diasLaborales := dias[0:5] // Del índice 0 al 4
	fmt.Println("Días laborales:", diasLaborales)

	finDeSemana := dias[5:] // Desde índice 5 hasta el final
	fmt.Println("Fin de semana:", finDeSemana)

	// Crear slice con make
	slice1 := make([]int, 5)    // longitud 5, capacidad 5
	slice2 := make([]int, 3, 5) // longitud 3, capacidad 5
	fmt.Printf("slice1: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	// Slice literal
	frutas := []string{"Manzana", "Banana", "Naranja"}
	fmt.Println("Frutas:", frutas)

	// Append (agregar elementos)
	frutas = append(frutas, "Pera")
	fmt.Println("Frutas después de append:", frutas)

	// Append múltiples elementos
	frutas = append(frutas, "Uva", "Sandía")
	fmt.Println("Más frutas:", frutas)

	// Append de otro slice
	masFrutas := []string{"Kiwi", "Mango"}
	frutas = append(frutas, masFrutas...)
	fmt.Println("Todas las frutas:", frutas)

	// Copy
	copiaFrutas := make([]string, len(frutas))
	copy(copiaFrutas, frutas)
	fmt.Println("Copia de frutas:", copiaFrutas)

	// Modificar copia no afecta original
	copiaFrutas[0] = "Fresa"
	fmt.Println("Original:", frutas)
	fmt.Println("Copia modificada:", copiaFrutas)

	// Eliminar elemento (reconstruir slice)
	fmt.Println("\n=== Eliminar elementos ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original:", nums)

	// Eliminar elemento en índice 2 (el 3)
	indice := 2
	nums = append(nums[:indice], nums[indice+1:]...)
	fmt.Println("Sin el elemento en índice 2:", nums)

	// Slice multidimensional
	fmt.Println("\n=== Slice 2D (Matriz) ===")
	matriz := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("Matriz:")
	for i, fila := range matriz {
		fmt.Printf("Fila %d: %v\n", i, fila)
	}

	// Iterar sobre slice
	fmt.Println("\n=== Iteración ===")
	for i, fruta := range frutas {
		fmt.Printf("%d: %s\n", i, fruta)
	}

	// Slice vacío vs nil
	fmt.Println("\n=== Slice nil vs vacío ===")
	var sliceNil []int
	sliceVacio := []int{}
	fmt.Printf("sliceNil: %v, es nil: %t, len: %d\n", sliceNil, sliceNil == nil, len(sliceNil))
	fmt.Printf("sliceVacio: %v, es nil: %t, len: %d\n", sliceVacio, sliceVacio == nil, len(sliceVacio))
}
