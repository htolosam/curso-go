package main

import "fmt"

func main() {
	// ===== CREAR MAPS =====
	fmt.Println("=== CREAR MAPS ===")

	// Map vacío con make
	edades := make(map[string]int)
	fmt.Println("Map vacío:", edades)

	// Map literal
	capitales := map[string]string{
		"España":    "Madrid",
		"Francia":   "París",
		"Italia":    "Roma",
		"Alemania":  "Berlín",
		"Portugal":  "Lisboa",
	}
	fmt.Println("Capitales:", capitales)

	// ===== OPERACIONES BÁSICAS =====
	fmt.Println("\n=== OPERACIONES ===")

	// Agregar/Modificar elementos
	edades["Ana"] = 25
	edades["Carlos"] = 30
	edades["María"] = 28
	fmt.Println("Edades:", edades)

	// Modificar valor existente
	edades["Ana"] = 26
	fmt.Println("Edad de Ana actualizada:", edades["Ana"])

	// Leer valor
	edad := edades["Carlos"]
	fmt.Printf("Carlos tiene %d años\n", edad)

	// Leer valor que no existe (retorna zero value)
	edadPedro := edades["Pedro"]
	fmt.Printf("Pedro tiene %d años (no existe)\n", edadPedro)

	// Verificar si existe una clave
	edadLuis, existe := edades["Luis"]
	if existe {
		fmt.Printf("Luis tiene %d años\n", edadLuis)
	} else {
		fmt.Println("Luis no está en el map")
	}

	// Forma idiomática de verificar y usar
	if edad, existe := edades["María"]; existe {
		fmt.Printf("María tiene %d años\n", edad)
	}

	// Eliminar elemento
	delete(edades, "Carlos")
	fmt.Println("Después de eliminar a Carlos:", edades)

	// Longitud del map
	fmt.Printf("Número de personas: %d\n", len(edades))

	// ===== ITERAR SOBRE MAP =====
	fmt.Println("\n=== ITERACIÓN ===")

	puntuaciones := map[string]int{
		"Matemáticas": 95,
		"Historia":    87,
		"Ciencias":    92,
		"Literatura":  88,
	}

	fmt.Println("Puntuaciones:")
	for asignatura, puntuacion := range puntuaciones {
		fmt.Printf("  %s: %d\n", asignatura, puntuacion)
	}

	// Solo claves
	fmt.Println("\nAsignaturas:")
	for asignatura := range puntuaciones {
		fmt.Printf("  - %s\n", asignatura)
	}

	// ===== MAPS ANIDADOS =====
	fmt.Println("\n=== MAPS ANIDADOS ===")

	usuarios := map[string]map[string]string{
		"user1": {
			"nombre":   "Juan",
			"email":    "juan@email.com",
			"ciudad":   "Madrid",
		},
		"user2": {
			"nombre":   "Ana",
			"email":    "ana@email.com",
			"ciudad":   "Barcelona",
		},
	}

	fmt.Println("Usuarios:")
	for id, datos := range usuarios {
		fmt.Printf("\n%s:\n", id)
		for campo, valor := range datos {
			fmt.Printf("  %s: %s\n", campo, valor)
		}
	}

	// ===== MAPS CON STRUCTS =====
	fmt.Println("\n=== MAPS CON STRUCTS ===")

	type Producto struct {
		Nombre string
		Precio float64
		Stock  int
	}

	inventario := make(map[string]Producto)
	
	inventario["001"] = Producto{
		Nombre: "Laptop",
		Precio: 899.99,
		Stock:  10,
	}
	
	inventario["002"] = Producto{
		Nombre: "Mouse",
		Precio: 19.99,
		Stock:  50,
	}

	fmt.Println("Inventario:")
	for codigo, producto := range inventario {
		fmt.Printf("%s: %s - $%.2f (Stock: %d)\n",
			codigo, producto.Nombre, producto.Precio, producto.Stock)
	}

	// ===== MAPS DE SLICES =====
	fmt.Println("\n=== MAPS DE SLICES ===")

	equipos := map[string][]string{
		"Desarrollo": {"Ana", "Carlos", "Luis"},
		"Diseño":     {"María", "Pedro"},
		"Marketing":  {"Laura", "Jorge", "Elena", "Pablo"},
	}

	fmt.Println("Equipos:")
	for departamento, miembros := range equipos {
		fmt.Printf("%s (%d miembros): %v\n",
			departamento, len(miembros), miembros)
	}

	// ===== MAP NIL =====
	fmt.Println("\n=== MAP NIL ===")
	var mapNil map[string]int
	fmt.Printf("mapNil: %v, es nil: %t\n", mapNil, mapNil == nil)
	// fmt.Println(mapNil["test"]) // Leer de map nil es seguro (retorna zero value)
	// mapNil["test"] = 1 // ERROR: asignación a map nil causa panic
}
