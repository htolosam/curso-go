package main

import "fmt"

// ===== DEFINIR STRUCTS =====

// Struct simple
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

// Struct con campos anónimos (embebidos)
type Direccion struct {
	Calle  string
	Ciudad string
	CP     string
}

// Struct con otro struct embebido
type Empleado struct {
	Nombre    string
	Puesto    string
	Salario   float64
	Direccion Direccion
}

// Struct con composición (embedding)
type Estudiante struct {
	Persona // Campo anónimo - hereda los campos de Persona
	Carrera string
	Año     int
}

// ===== MÉTODOS =====

// Método con receptor por valor
func (p Persona) Saludar() string {
	return fmt.Sprintf("Hola, soy %s y tengo %d años", p.Nombre, p.Edad)
}

// Método con receptor por puntero (puede modificar el struct)
func (p *Persona) CumplirAños() {
	p.Edad++
}

// Método que retorna valor
func (e Empleado) SalarioAnual() float64 {
	return e.Salario * 12
}

// ===== FUNCIONES CONSTRUCTORAS =====

func NuevaPersona(nombre string, edad int, email string) *Persona {
	return &Persona{
		Nombre: nombre,
		Edad:   edad,
		Email:  email,
	}
}

func main() {
	// ===== CREAR STRUCTS =====
	fmt.Println("=== CREAR STRUCTS ===")

	// Forma 1: Declaración con valores cero
	var p1 Persona
	fmt.Printf("Persona vacía: %+v\n", p1)

	// Forma 2: Literal con nombres de campos
	p2 := Persona{
		Nombre: "Ana",
		Edad:   25,
		Email:  "ana@email.com",
	}
	fmt.Printf("Ana: %+v\n", p2)

	// Forma 3: Literal sin nombres (orden importa)
	p3 := Persona{"Carlos", 30, "carlos@email.com"}
	fmt.Printf("Carlos: %+v\n", p3)

	// Forma 4: Con función constructora
	p4 := NuevaPersona("María", 28, "maria@email.com")
	fmt.Printf("María: %+v\n", p4)

	// ===== ACCEDER A CAMPOS =====
	fmt.Println("\n=== ACCESO A CAMPOS ===")
	fmt.Printf("Nombre: %s\n", p2.Nombre)
	fmt.Printf("Edad: %d\n", p2.Edad)

	// Modificar campos
	p2.Edad = 26
	fmt.Printf("Nueva edad de Ana: %d\n", p2.Edad)

	// ===== PUNTEROS A STRUCTS =====
	fmt.Println("\n=== PUNTEROS ===")
	p5 := &Persona{
		Nombre: "Luis",
		Edad:   35,
		Email:  "luis@email.com",
	}
	fmt.Printf("Luis (puntero): %+v\n", p5)
	// Go permite acceder a campos directamente sin desreferenciar
	fmt.Printf("Nombre: %s\n", p5.Nombre)

	// ===== STRUCTS ANIDADOS =====
	fmt.Println("\n=== STRUCTS ANIDADOS ===")

	empleado := Empleado{
		Nombre:  "Pedro",
		Puesto:  "Desarrollador",
		Salario: 3000,
		Direccion: Direccion{
			Calle:  "Calle Mayor 10",
			Ciudad: "Madrid",
			CP:     "28001",
		},
	}

	fmt.Printf("Empleado: %+v\n", empleado)
	fmt.Printf("Ciudad: %s\n", empleado.Direccion.Ciudad)

	// ===== COMPOSICIÓN (EMBEDDING) =====
	fmt.Println("\n=== COMPOSICIÓN ===")

	estudiante := Estudiante{
		Persona: Persona{
			Nombre: "Laura",
			Edad:   20,
			Email:  "laura@uni.com",
		},
		Carrera: "Ingeniería Informática",
		Año:     2,
	}

	fmt.Printf("Estudiante: %+v\n", estudiante)
	// Acceso directo a campos embebidos
	fmt.Printf("Nombre: %s\n", estudiante.Nombre)
	fmt.Printf("Carrera: %s\n", estudiante.Carrera)

	// ===== MÉTODOS =====
	fmt.Println("\n=== MÉTODOS ===")

	persona := Persona{
		Nombre: "Jorge",
		Edad:   40,
		Email:  "jorge@email.com",
	}

	// Llamar método
	saludo := persona.Saludar()
	fmt.Println(saludo)

	// Método que modifica el struct
	fmt.Printf("Edad antes: %d\n", persona.Edad)
	persona.CumplirAños()
	fmt.Printf("Edad después: %d\n", persona.Edad)

	// Método de struct anidado
	fmt.Printf("Salario anual de Pedro: %.2f\n", empleado.SalarioAnual())

	// ===== COMPARACIÓN DE STRUCTS =====
	fmt.Println("\n=== COMPARACIÓN ===")

	a := Persona{"Ana", 25, "ana@email.com"}
	b := Persona{"Ana", 25, "ana@email.com"}
	c := Persona{"Carlos", 30, "carlos@email.com"}

	fmt.Printf("a == b: %t\n", a == b) // true
	fmt.Printf("a == c: %t\n", a == c) // false

	// ===== STRUCT VACÍO =====
	fmt.Println("\n=== STRUCT VACÍO ===")

	// Struct sin campos (útil para sets o señales)
	type Vacio struct{}
	v := Vacio{}
	fmt.Printf("Struct vacío: %+v, tamaño: %d bytes\n", v, 0)

	// Set usando map con struct vacío
	set := make(map[string]struct{})
	set["a"] = struct{}{}
	set["b"] = struct{}{}
	fmt.Printf("Set: %v\n", set)

	// ===== TAGS DE STRUCTS =====
	fmt.Println("\n=== TAGS ===")

	type Usuario struct {
		ID       int    `json:"id"`
		Nombre   string `json:"nombre"`
		Email    string `json:"email,omitempty"`
		Password string `json:"-"` // No se serializa
	}

	u := Usuario{ID: 1, Nombre: "Ana", Email: "ana@test.com", Password: "secret"}
	fmt.Printf("Usuario con tags: %+v\n", u)
	// Las tags se usan con paquetes como encoding/json
}
