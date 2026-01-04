package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== GOROUTINES BÁSICAS =====

func decirHola() {
	fmt.Println("¡Hola desde una goroutine!")
}

func contar(nombre string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%s: %d\n", nombre, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ===== GOROUTINES CON PARÁMETROS =====

func saludar(nombre string, veces int) {
	for i := 0; i < veces; i++ {
		fmt.Printf("¡Hola %s! (%d/%d)\n", nombre, i+1, veces)
		time.Sleep(200 * time.Millisecond)
	}
}

// ===== FUNCIÓN CON RETORNO USANDO CHANNEL =====

func calcularCuadrado(numero int) int {
	time.Sleep(100 * time.Millisecond) // Simular trabajo
	return numero * numero
}

// ===== WAIT GROUP =====

func trabajador(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca que terminó al salir
	fmt.Printf("Trabajador %d iniciando\n", id)
	time.Sleep(time.Duration(id*100) * time.Millisecond)
	fmt.Printf("Trabajador %d terminado\n", id)
}

// ===== GOROUTINES CON CLOSURE =====

func contarConClosure(nombre string) {
	// Goroutine con función anónima
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("%s: %d\n", nombre, i)
			time.Sleep(150 * time.Millisecond)
		}
	}()
}

// ===== POOL DE TRABAJADORES =====

func procesarTarea(id int, tarea int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Trabajador %d procesando tarea %d\n", id, tarea)
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Trabajador %d completó tarea %d\n", id, tarea)
}

func main() {
	// ===== GOROUTINE BÁSICA =====
	fmt.Println("=== GOROUTINE BÁSICA ===")
	
	// Ejecutar en background
	go decirHola()
	
	// Dar tiempo a que se ejecute
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Función main continúa...")

	// ===== MÚLTIPLES GOROUTINES =====
	fmt.Println("\n=== MÚLTIPLES GOROUTINES ===")
	
	go contar("Goroutine 1")
	go contar("Goroutine 2")
	go contar("Goroutine 3")
	
	// Esperar a que terminen
	time.Sleep(600 * time.Millisecond)

	// ===== GOROUTINES CON PARÁMETROS =====
	fmt.Println("\n=== GOROUTINES CON PARÁMETROS ===")
	
	go saludar("Ana", 3)
	go saludar("Carlos", 2)
	
	time.Sleep(1 * time.Second)

	// ===== WAIT GROUP =====
	fmt.Println("\n=== WAIT GROUP ===")
	
	var wg sync.WaitGroup
	
	// Lanzar 5 trabajadores
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementar contador
		go trabajador(i, &wg)
	}
	
	// Esperar a que todos terminen
	wg.Wait()
	fmt.Println("Todos los trabajadores terminaron")

	// ===== GOROUTINES CON CLOSURE =====
	fmt.Println("\n=== GOROUTINES CON CLOSURE ===")
	
	contarConClosure("Closure 1")
	contarConClosure("Closure 2")
	
	time.Sleep(600 * time.Millisecond)

	// ===== GOROUTINES EN BUCLE (CUIDADO CON CLOSURES) =====
	fmt.Println("\n=== GOROUTINES EN BUCLE ===")
	
	// Forma INCORRECTA (todas las goroutines verían i=5)
	fmt.Println("Forma incorrecta:")
	for i := 1; i <= 3; i++ {
		go func() {
			// time.Sleep(10 * time.Millisecond)
			// fmt.Printf("Valor incorrecto: %d\n", i) // Todas podrían mostrar 4
		}()
	}
	
	// Forma CORRECTA (pasar el valor como parámetro)
	fmt.Println("Forma correcta:")
	for i := 1; i <= 3; i++ {
		go func(numero int) {
			fmt.Printf("Valor correcto: %d\n", numero)
		}(i) // Pasar i como argumento
	}
	
	time.Sleep(100 * time.Millisecond)

	// ===== GOROUTINES CON RETORNO USANDO CHANNELS =====
	fmt.Println("\n=== GOROUTINES CON CANALES ===")
	
	resultados := make(chan int, 5)
	
	// Lanzar cálculos en paralelo
	numeros := []int{1, 2, 3, 4, 5}
	for _, num := range numeros {
		go func(n int) {
			resultados <- calcularCuadrado(n)
		}(num)
	}
	
	// Recolectar resultados
	fmt.Println("Cuadrados:")
	for i := 0; i < len(numeros); i++ {
		fmt.Printf("%d ", <-resultados)
	}
	fmt.Println()
	close(resultados)

	// ===== POOL DE TRABAJADORES =====
	fmt.Println("\n=== POOL DE TRABAJADORES ===")
	
	var wg2 sync.WaitGroup
	numTrabajadores := 3
	tareas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	// Canal para las tareas
	canalTareas := make(chan int, len(tareas))
	
	// Llenar canal con tareas
	for _, tarea := range tareas {
		canalTareas <- tarea
	}
	close(canalTareas)
	
	// Crear trabajadores
	for i := 1; i <= numTrabajadores; i++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			for tarea := range canalTareas {
				fmt.Printf("Trabajador %d procesando tarea %d\n", id, tarea)
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}
	
	wg2.Wait()
	fmt.Println("Todas las tareas procesadas")

	// ===== GOROUTINE ANÓNIMA =====
	fmt.Println("\n=== GOROUTINE ANÓNIMA ===")
	
	mensaje := "¡Hola desde goroutine anónima!"
	go func() {
		fmt.Println(mensaje)
	}()
	
	time.Sleep(100 * time.Millisecond)

	// ===== BUENAS PRÁCTICAS =====
	fmt.Println("\n=== BUENAS PRÁCTICAS ===")
	fmt.Println("1. Usar WaitGroup para esperar goroutines")
	fmt.Println("2. Pasar variables de bucle como parámetros")
	fmt.Println("3. Usar channels para comunicación")
	fmt.Println("4. Evitar compartir memoria; usar channels")
	fmt.Println("5. Tener cuidado con closures y variables capturadas")
	fmt.Println("6. Siempre asegurar que las goroutines terminen")
}
