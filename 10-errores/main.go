package main

import (
	"errors"
	"fmt"
	"strconv"
)

// ===== CREAR ERRORES =====

// Error simple con errors.New
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

// Error con fmt.Errorf (permite formateo)
func validarEdad(edad int) error {
	if edad < 0 {
		return fmt.Errorf("edad inválida: %d (debe ser positiva)", edad)
	}
	if edad < 18 {
		return fmt.Errorf("edad insuficiente: %d (debe ser >= 18)", edad)
	}
	return nil
}

// ===== ERRORES PERSONALIZADOS =====

// Struct de error personalizado
type ErrorValidacion struct {
	Campo   string
	Valor   interface{}
	Mensaje string
}

func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("error en campo '%s' (valor: %v): %s", e.Campo, e.Valor, e.Mensaje)
}

func validarUsuario(nombre string, edad int) error {
	if nombre == "" {
		return &ErrorValidacion{
			Campo:   "nombre",
			Valor:   nombre,
			Mensaje: "no puede estar vacío",
		}
	}
	if edad < 0 {
		return &ErrorValidacion{
			Campo:   "edad",
			Valor:   edad,
			Mensaje: "debe ser positiva",
		}
	}
	return nil
}

// ===== MÚLTIPLES ERRORES =====

type ErrorMultiple struct {
	Errores []error
}

func (e *ErrorMultiple) Error() string {
	if len(e.Errores) == 0 {
		return "sin errores"
	}
	msg := fmt.Sprintf("se encontraron %d errores:\n", len(e.Errores))
	for i, err := range e.Errores {
		msg += fmt.Sprintf("  %d. %s\n", i+1, err.Error())
	}
	return msg
}

func validarFormulario(nombre string, email string, edad int) error {
	var errores []error

	if nombre == "" {
		errores = append(errores, errors.New("nombre requerido"))
	}
	if email == "" {
		errores = append(errores, errors.New("email requerido"))
	}
	if edad < 18 {
		errores = append(errores, errors.New("debe ser mayor de edad"))
	}

	if len(errores) > 0 {
		return &ErrorMultiple{Errores: errores}
	}
	return nil
}

// ===== WRAPPING ERRORS (Go 1.13+) =====

func procesarArchivo(nombre string) error {
	// Simular lectura de archivo
	err := leerArchivo(nombre)
	if err != nil {
		// Envolver el error con contexto adicional
		return fmt.Errorf("error al procesar archivo %s: %w", nombre, err)
	}
	return nil
}

func leerArchivo(nombre string) error {
	if nombre == "" {
		return errors.New("nombre de archivo vacío")
	}
	// Simular que el archivo no existe
	return errors.New("archivo no encontrado")
}

// ===== SENTINEL ERRORS =====

var (
	ErrNoEncontrado   = errors.New("recurso no encontrado")
	ErrNoAutorizado   = errors.New("acceso no autorizado")
	ErrDatosInvalidos = errors.New("datos inválidos")
)

func buscarUsuario(id int) (string, error) {
	if id <= 0 {
		return "", ErrDatosInvalidos
	}
	if id == 999 {
		return "", ErrNoAutorizado
	}
	// Simular que no se encuentra
	return "", ErrNoEncontrado
}

// ===== PANIC Y RECOVER =====

func usarPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperado de panic: %v\n", r)
		}
	}()

	fmt.Println("Antes del panic")
	panic("¡Algo salió muy mal!")
	fmt.Println("Esto nunca se ejecuta")
}

func dividirConPanic(a, b int) int {
	if b == 0 {
		panic("división por cero no permitida")
	}
	return a / b
}

func usarDividirConPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error capturado: %v\n", r)
		}
	}()

	resultado := dividirConPanic(10, 2)
	fmt.Printf("10 / 2 = %d\n", resultado)

	resultado = dividirConPanic(10, 0) // Esto causa panic
	fmt.Println("Esta línea no se ejecuta")
}

func main() {
	// ===== ERRORES BÁSICOS =====
	fmt.Println("=== ERRORES BÁSICOS ===")

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", resultado)
	}

	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Resultado: %.2f\n", resultado)
	}

	// ===== VALIDACIÓN =====
	fmt.Println("\n=== VALIDACIÓN ===")

	err = validarEdad(25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Edad válida")
	}

	err = validarEdad(15)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = validarEdad(-5)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// ===== ERRORES PERSONALIZADOS =====
	fmt.Println("\n=== ERRORES PERSONALIZADOS ===")

	err = validarUsuario("Ana", 25)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Usuario válido")
	}

	err = validarUsuario("", 25)
	if err != nil {
		fmt.Println("Error:", err)
		// Type assertion para obtener más información
		if ve, ok := err.(*ErrorValidacion); ok {
			fmt.Printf("  Campo problemático: %s\n", ve.Campo)
		}
	}

	// ===== MÚLTIPLES ERRORES =====
	fmt.Println("\n=== MÚLTIPLES ERRORES ===")

	err = validarFormulario("Juan", "juan@email.com", 25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Formulario válido")
	}

	err = validarFormulario("", "", 15)
	if err != nil {
		fmt.Println(err)
	}

	// ===== WRAPPING ERRORS =====
	fmt.Println("\n=== WRAPPING ERRORS ===")

	err = procesarArchivo("datos.txt")
	if err != nil {
		fmt.Println("Error:", err)
		
		// Unwrap para obtener el error original
		if errors.Is(err, errors.New("archivo no encontrado")) {
			fmt.Println("Es un error de archivo no encontrado")
		}
	}

	// ===== SENTINEL ERRORS =====
	fmt.Println("\n=== SENTINEL ERRORS ===")

	usuario, err := buscarUsuario(1)
	if err != nil {
		// Comparar con errores conocidos
		if errors.Is(err, ErrNoEncontrado) {
			fmt.Println("Usuario no encontrado")
		} else if errors.Is(err, ErrNoAutorizado) {
			fmt.Println("Acceso denegado")
		} else if errors.Is(err, ErrDatosInvalidos) {
			fmt.Println("ID inválido")
		}
	} else {
		fmt.Println("Usuario:", usuario)
	}

	// ===== CONVERTIR STRING A NÚMERO =====
	fmt.Println("\n=== MANEJO DE ERRORES EN CONVERSIÓN ===")

	numero, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error al convertir:", err)
	} else {
		fmt.Printf("Número convertido: %d\n", numero)
	}

	numero, err = strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Error al convertir:", err)
	}

	// ===== PANIC Y RECOVER =====
	fmt.Println("\n=== PANIC Y RECOVER ===")

	usarPanic()
	fmt.Println("El programa continúa después del recover")

	usarDividirConPanic()
	fmt.Println("El programa continúa después del segundo recover")

	// ===== MEJORES PRÁCTICAS =====
	fmt.Println("\n=== MEJORES PRÁCTICAS ===")
	fmt.Println("1. Siempre verifica los errores")
	fmt.Println("2. Proporciona contexto útil en los mensajes de error")
	fmt.Println("3. Usa errores personalizados para casos complejos")
	fmt.Println("4. No uses panic para control de flujo normal")
	fmt.Println("5. Usa sentinel errors para errores conocidos")
	fmt.Println("6. Envuelve errores con fmt.Errorf y %w")
}
