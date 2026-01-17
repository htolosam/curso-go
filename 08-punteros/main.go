package main

import "fmt"

// ===== CONCEPTOS BÁSICOS =====

func duplicar(x int) {
	x = x * 2
	fmt.Printf("Dentro de duplicar: x = %d\n", x)
}

func duplicarConPuntero(x *int) {
	*x = *x * 2
	fmt.Printf("Dentro de duplicarConPuntero: *x = %d\n", *x)
}

// ===== STRUCTS Y PUNTEROS =====

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) CumplirAñosPorValor() {
	p.Edad++
	fmt.Printf("  Dentro método por valor: edad = %d\n", p.Edad)
}

func (p *Persona) CumplirAñosPorPuntero() {
	p.Edad++
	fmt.Printf("  Dentro método por puntero: edad = %d\n", p.Edad)
}

func modificarPersona(p *Persona, nuevoNombre string) {
	p.Nombre = nuevoNombre
}

// ===== SLICES Y PUNTEROS =====

func agregarElemento(nums []int, valor int) []int {
	return append(nums, valor)
}

func modificarSlice(nums []int) {
	if len(nums) > 0 {
		nums[0] = 999
	}
}

func main() {
	// ===== INTRODUCCIÓN A PUNTEROS =====
	fmt.Println("=== INTRODUCCIÓN A PUNTEROS ===")

	x := 10
	fmt.Printf("Valor de x: %d\n", x)
	fmt.Printf("Dirección de x: %p\n", &x)

	// Crear puntero
	var p *int = &x
	fmt.Printf("Puntero p apunta a: %p\n", p)
	fmt.Printf("Valor al que apunta p: %d\n", *p)

	// Modificar a través del puntero
	*p = 20
	fmt.Printf("Después de *p = 20, x = %d\n", x)

	// ===== PASO POR VALOR VS PUNTERO =====
	fmt.Println("\n=== PASO POR VALOR VS PUNTERO ===")

	y := 5
	fmt.Printf("Antes de duplicar: y = %d\n", y)
	duplicar(y)
	fmt.Printf("Después de duplicar: y = %d (no cambió)\n", y)

	z := 5
	fmt.Printf("\nAntes de duplicarConPuntero: z = %d\n", z)
	duplicarConPuntero(&z)
	fmt.Printf("Después de duplicarConPuntero: z = %d (cambió!)\n", z)

	// ===== NEW =====
	fmt.Println("\n=== NEW ===")

	// new crea un puntero a un valor zero
	ptr := new(int)
	fmt.Printf("ptr: %p, *ptr: %d\n", ptr, *ptr)
	*ptr = 100
	fmt.Printf("Después de asignar: *ptr: %d\n", *ptr)

	// ===== PUNTEROS CON STRUCTS =====
	fmt.Println("\n=== PUNTEROS CON STRUCTS ===")

	persona1 := Persona{Nombre: "Ana", Edad: 25}
	fmt.Printf("Persona1: %+v\n", persona1)

	// Método por valor (no modifica el original)
	fmt.Println("\nLlamando método por valor:")
	persona1.CumplirAñosPorValor()
	fmt.Printf("Después del método: %+v (no cambió)\n", persona1)

	// Método por puntero (modifica el original)
	fmt.Println("\nLlamando método por puntero:")
	persona1.CumplirAñosPorPuntero()
	fmt.Printf("Después del método: %+v (cambió!)\n", persona1)

	// ===== CREAR STRUCTS CON PUNTEROS =====
	fmt.Println("\n=== CREAR STRUCTS CON PUNTEROS ===")

	// Forma 1: Crear y obtener puntero
	persona2 := Persona{Nombre: "Carlos", Edad: 30}
	ptrPersona2 := &persona2
	fmt.Printf("persona2: %+v\n", persona2)
	fmt.Printf("ptrPersona2: %+v\n", ptrPersona2)

	// Forma 2: Crear directamente como puntero
	persona3 := &Persona{Nombre: "María", Edad: 28}
	fmt.Printf("persona3: %+v\n", persona3)

	// Modificar a través de puntero
	modificarPersona(persona3, "María González")
	fmt.Printf("Después de modificar: %+v\n", persona3)

	// ===== NIL POINTERS =====
	fmt.Println("\n=== NIL POINTERS ===")

	var ptrNil *Persona
	fmt.Printf("ptrNil: %v, es nil: %t\n", ptrNil, ptrNil == nil)

	// Verificar antes de desreferenciar
	if ptrNil != nil {
		fmt.Println(ptrNil.Nombre)
	} else {
		fmt.Println("El puntero es nil, no se puede desreferenciar")
	}

	// ===== PUNTEROS EN SLICES =====
	fmt.Println("\n=== PUNTEROS EN SLICES ===")

	// Los slices son referencias, pero veamos el comportamiento
	numeros := []int{1, 2, 3}
	fmt.Printf("Original: %v\n", numeros)

	// Modificar contenido del slice
	modificarSlice(numeros)
	fmt.Printf("Después de modificarSlice: %v (cambió!)\n", numeros)

	// Append no modifica el original (puede crear nuevo array subyacente)
	numeros2 := []int{1, 2, 3}
	fmt.Printf("\nOriginal: %v\n", numeros2)
	nuevosNumeros := agregarElemento(numeros2, 4)
	fmt.Printf("Original después de agregarElemento: %v\n", numeros2)
	fmt.Printf("Nuevo slice: %v\n", nuevosNumeros)

	// ===== SLICE DE PUNTEROS =====
	fmt.Println("\n=== SLICE DE PUNTEROS A STRUCTS ===")

	personas := []*Persona{
		{Nombre: "Juan", Edad: 30},
		{Nombre: "Laura", Edad: 25},
		{Nombre: "Pedro", Edad: 35},
	}

	fmt.Println("Personas:")
	for i, p := range personas {
		fmt.Printf("%d: %+v\n", i, p)
	}

	// Modificar a través de puntero en el slice
	personas[0].Edad = 31
	fmt.Printf("\nDespués de modificar: %+v\n", personas[0])

	// ===== COMPARACIÓN DE PUNTEROS =====
	fmt.Println("\n=== COMPARACIÓN DE PUNTEROS ===")

	a := 10
	b := 10
	ptrA := &a
	ptrB := &b
	ptrA2 := &a

	fmt.Printf("ptrA == ptrB: %t (apuntan a diferentes variables)\n", ptrA == ptrB)
	fmt.Printf("ptrA == ptrA2: %t (apuntan a la misma variable)\n", ptrA == ptrA2)
	fmt.Printf("*ptrA == *ptrB: %t (los valores son iguales)\n", *ptrA == *ptrB)

	// ===== CUÁNDO USAR PUNTEROS =====
	fmt.Println("\n=== CUÁNDO USAR PUNTEROS ===")
	fmt.Println("Usa punteros cuando:")
	fmt.Println("1. Necesitas modificar el valor original")
	fmt.Println("2. El struct es grande y copiar sería costoso")
	fmt.Println("3. Quieres compartir datos entre funciones")
	fmt.Println("4. Trabajas con interfaces que modifican estado")
}
