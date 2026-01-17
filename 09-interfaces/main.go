package main

import (
	"fmt"
	"math"
)

// ===== DEFINIR INTERFACES =====

// Interface simple
type Hablador interface {
	Hablar() string
}

// Interface con múltiples métodos
type Forma interface {
	Area() float64
	Perimetro() float64
}

// ===== IMPLEMENTAR INTERFACES =====

// Persona implementa Hablador
type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Hablar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

// Perro también implementa Hablador
type Perro struct {
	Nombre string
	Raza   string
}

func (p Perro) Hablar() string {
	return fmt.Sprintf("¡Guau! Soy %s, un %s", p.Nombre, p.Raza)
}

// Robot también implementa Hablador
type Robot struct {
	Modelo string
}

func (r Robot) Hablar() string {
	return fmt.Sprintf("Procesando... Soy el modelo %s", r.Modelo)
}

// ===== IMPLEMENTAR INTERFACE FORMA =====

type Rectangulo struct {
	Ancho  float64
	Alto   float64
}

func (r Rectangulo) Area() float64 {
	return r.Ancho * r.Alto
}

func (r Rectangulo) Perimetro() float64 {
	return 2 * (r.Ancho + r.Alto)
}

type Circulo struct {
	Radio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Radio * c.Radio
}

func (c Circulo) Perimetro() float64 {
	return 2 * math.Pi * c.Radio
}

// ===== FUNCIONES QUE USAN INTERFACES =====

func hacerHablar(h Hablador) {
	fmt.Println(h.Hablar())
}

func imprimirInfo(f Forma) {
	fmt.Printf("Área: %.2f, Perímetro: %.2f\n", f.Area(), f.Perimetro())
}

// ===== INTERFACE VACÍA =====

func imprimirCualquierCosa(i interface{}) {
	fmt.Printf("Valor: %v, Tipo: %T\n", i, i)
}

// ===== TYPE ASSERTION =====

func procesarValor(i interface{}) {
	// Type assertion básica
	if str, ok := i.(string); ok {
		fmt.Printf("Es un string: %s (longitud: %d)\n", str, len(str))
		return
	}

	if num, ok := i.(int); ok {
		fmt.Printf("Es un int: %d (doble: %d)\n", num, num*2)
		return
	}

	fmt.Println("Tipo no reconocido")
}

// ===== TYPE SWITCH =====

func identificarTipo(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Printf("String: %s\n", v)
	case int:
		fmt.Printf("Int: %d\n", v)
	case float64:
		fmt.Printf("Float64: %.2f\n", v)
	case bool:
		fmt.Printf("Bool: %t\n", v)
	case Persona:
		fmt.Printf("Persona: %s\n", v.Nombre)
	default:
		fmt.Printf("Tipo desconocido: %T\n", v)
	}
}

// ===== COMPOSICIÓN DE INTERFACES =====

type Lector interface {
	Leer() string
}

type Escritor interface {
	Escribir(string)
}

// Interface compuesta
type LectorEscritor interface {
	Lector
	Escritor
}

type Archivo struct {
	Contenido string
}

func (a *Archivo) Leer() string {
	return a.Contenido
}

func (a *Archivo) Escribir(texto string) {
	a.Contenido = texto
}

// ===== INTERFACE STRINGER =====

type Producto struct {
	Nombre string
	Precio float64
}

// Implementar la interface Stringer de fmt
func (p Producto) String() string {
	return fmt.Sprintf("%s - $%.2f", p.Nombre, p.Precio)
}

func main() {
	// ===== USAR INTERFACES =====
	fmt.Println("=== INTERFACES BÁSICAS ===")

	persona := Persona{Nombre: "Ana", Edad: 25}
	perro := Perro{Nombre: "Max", Raza: "Labrador"}
	robot := Robot{Modelo: "X-2000"}

	// Todos implementan Hablador
	hacerHablar(persona)
	hacerHablar(perro)
	hacerHablar(robot)

	// ===== SLICE DE INTERFACES =====
	fmt.Println("\n=== SLICE DE INTERFACES ===")

	habladores := []Hablador{persona, perro, robot}
	for i, h := range habladores {
		fmt.Printf("%d: %s\n", i+1, h.Hablar())
	}

	// ===== INTERFACE FORMA =====
	fmt.Println("\n=== INTERFACE FORMA ===")

	rect := Rectangulo{Ancho: 5, Alto: 3}
	circ := Circulo{Radio: 4}

	fmt.Println("Rectángulo:")
	imprimirInfo(rect)

	fmt.Println("\nCírculo:")
	imprimirInfo(circ)

	formas := []Forma{rect, circ}
	fmt.Println("\nTodas las formas:")
	for _, f := range formas {
		imprimirInfo(f)
	}

	// ===== INTERFACE VACÍA =====
	fmt.Println("\n=== INTERFACE VACÍA ===")

	imprimirCualquierCosa(42)
	imprimirCualquierCosa("Hola")
	imprimirCualquierCosa(3.14)
	imprimirCualquierCosa(true)
	imprimirCualquierCosa(persona)

	// ===== TYPE ASSERTION =====
	fmt.Println("\n=== TYPE ASSERTION ===")

	var i interface{} = "Hola Mundo"
	procesarValor(i)

	i = 42
	procesarValor(i)

	i = 3.14
	procesarValor(i)

	// ===== TYPE SWITCH =====
	fmt.Println("\n=== TYPE SWITCH ===")

	valores := []interface{}{
		"texto",
		123,
		45.67,
		true,
		Persona{Nombre: "Carlos", Edad: 30},
		[]int{1, 2, 3},
	}

	for _, v := range valores {
		identificarTipo(v)
	}

	// ===== COMPOSICIÓN DE INTERFACES =====
	fmt.Println("\n=== COMPOSICIÓN DE INTERFACES ===")

	archivo := &Archivo{}
	archivo.Escribir("Contenido del archivo")
	fmt.Printf("Leer: %s\n", archivo.Leer())

	// archivo implementa LectorEscritor
	var le LectorEscritor = archivo
	le.Escribir("Nuevo contenido")
	fmt.Printf("Leer: %s\n", le.Leer())

	// ===== INTERFACE STRINGER =====
	fmt.Println("\n=== INTERFACE STRINGER ===")

	prod1 := Producto{Nombre: "Laptop", Precio: 999.99}
	prod2 := Producto{Nombre: "Mouse", Precio: 29.99}

	// fmt.Println usa automáticamente String() si está implementado
	fmt.Println(prod1)
	fmt.Println(prod2)

	// ===== INTERFACE NIL =====
	fmt.Println("\n=== INTERFACE NIL ===")

	var h Hablador
	fmt.Printf("Interface nil: %v, es nil: %t\n", h, h == nil)

	// Asignar valor
	h = persona
	fmt.Printf("Después de asignar: %v, es nil: %t\n", h, h == nil)

	// ===== VERIFICAR IMPLEMENTACIÓN =====
	fmt.Println("\n=== VERIFICAR IMPLEMENTACIÓN ===")

	// Esto se compila solo si Persona implementa Hablador
	var _ Hablador = Persona{}
	var _ Forma = Rectangulo{}
	fmt.Println("Todas las interfaces están implementadas correctamente")
}
