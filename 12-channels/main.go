package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== CHANNELS BÁSICOS =====

func enviarMensaje(ch chan string) {
	ch <- "¡Hola desde goroutine!"
}

func productor(ch chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Produciendo: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Cerrar canal cuando termine
}

func consumidor(ch chan int) {
	for valor := range ch { // Itera hasta que se cierre el canal
		fmt.Printf("Consumiendo: %d\n", valor)
		time.Sleep(150 * time.Millisecond)
	}
}

// ===== CHANNELS DIRECCIONALES =====

// Solo puede enviar
func soloEnviar(ch chan<- string) {
	ch <- "Mensaje enviado"
}

// Solo puede recibir
func soloRecibir(ch <-chan string) {
	mensaje := <-ch
	fmt.Println("Recibido:", mensaje)
}

// ===== CANAL BUFFERED =====

func demoBuffered() {
	// Canal con buffer de tamaño 3
	ch := make(chan int, 3)

	// Podemos enviar 3 valores sin bloquear
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("Enviados 3 valores al buffer")

	// Recibir valores
	fmt.Println("Recibiendo:", <-ch)
	fmt.Println("Recibiendo:", <-ch)
	fmt.Println("Recibiendo:", <-ch)
}

// ===== SELECT =====

func servidor1(ch chan string) {
	time.Sleep(200 * time.Millisecond)
	ch <- "Respuesta del servidor 1"
}

func servidor2(ch chan string) {
	time.Sleep(100 * time.Millisecond)
	ch <- "Respuesta del servidor 2"
}

func demoSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go servidor1(ch1)
	go servidor2(ch2)

	// Select espera el primero que responda
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}
}

// ===== SELECT CON TIMEOUT =====

func demoSelectConTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Respuesta tardía"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: operación tardó demasiado")
	}
}

// ===== SELECT CON DEFAULT =====

func demoSelectDefault() {
	ch := make(chan string, 1)

	// Intenta recibir sin bloquear
	select {
	case msg := <-ch:
		fmt.Println("Recibido:", msg)
	default:
		fmt.Println("No hay datos disponibles")
	}

	// Enviar y luego recibir
	ch <- "Ahora hay datos"

	select {
	case msg := <-ch:
		fmt.Println("Recibido:", msg)
	default:
		fmt.Println("No hay datos disponibles")
	}
}

// ===== MÚLTIPLES GOROUTINES CON CHANNELS =====

func trabajador(id int, tareas <-chan int, resultados chan<- int) {
	for tarea := range tareas {
		fmt.Printf("Trabajador %d procesando tarea %d\n", id, tarea)
		time.Sleep(100 * time.Millisecond)
		resultados <- tarea * 2
	}
}

func demoPoolTrabajadores() {
	const numTrabajadores = 3
	const numTareas = 9

	tareas := make(chan int, numTareas)
	resultados := make(chan int, numTareas)

	// Iniciar trabajadores
	for w := 1; w <= numTrabajadores; w++ {
		go trabajador(w, tareas, resultados)
	}

	// Enviar tareas
	for t := 1; t <= numTareas; t++ {
		tareas <- t
	}
	close(tareas)

	// Recolectar resultados
	for r := 1; r <= numTareas; r++ {
		resultado := <-resultados
		fmt.Printf("Resultado: %d\n", resultado)
	}
}

// ===== FAN-OUT FAN-IN =====

func generador(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func cuadrado(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	
	// Función para procesar cada canal
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}
	
	// Lanzar una goroutine por cada canal
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	
	// Cerrar el canal de salida cuando todos terminen
	go func() {
		wg.Wait()
		close(out)
	}()
	
	return out
}

// ===== DONE CHANNEL (PATRÓN DE CANCELACIÓN) =====

func demoDoneChannel() {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Goroutine cancelada")
				return
			default:
				fmt.Println("Trabajando...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	time.Sleep(600 * time.Millisecond)
	done <- true
	time.Sleep(100 * time.Millisecond)
}

func main() {
	// ===== CHANNEL BÁSICO =====
	fmt.Println("=== CHANNEL BÁSICO ===")

	ch := make(chan string)
	go enviarMensaje(ch)
	mensaje := <-ch // Recibir del canal (bloquea hasta que haya datos)
	fmt.Println("Recibido:", mensaje)

	// ===== PRODUCTOR-CONSUMIDOR =====
	fmt.Println("\n=== PRODUCTOR-CONSUMIDOR ===")

	canal := make(chan int)
	go productor(canal)
	consumidor(canal)

	// ===== CHANNELS DIRECCIONALES =====
	fmt.Println("\n=== CHANNELS DIRECCIONALES ===")

	ch2 := make(chan string, 1)
	go soloEnviar(ch2)
	soloRecibir(ch2)

	// ===== CANAL BUFFERED =====
	fmt.Println("\n=== CANAL BUFFERED ===")
	demoBuffered()

	// ===== SELECT =====
	fmt.Println("\n=== SELECT ===")
	demoSelect()

	// ===== SELECT CON TIMEOUT =====
	fmt.Println("\n=== SELECT CON TIMEOUT ===")
	demoSelectConTimeout()

	// ===== SELECT CON DEFAULT =====
	fmt.Println("\n=== SELECT CON DEFAULT ===")
	demoSelectDefault()

	// ===== POOL DE TRABAJADORES =====
	fmt.Println("\n=== POOL DE TRABAJADORES ===")
	demoPoolTrabajadores()

	// ===== FAN-OUT FAN-IN =====
	fmt.Println("\n=== FAN-OUT FAN-IN ===")

	// Generar números
	nums := generador(1, 2, 3, 4, 5)

	// Fan-out: múltiples goroutines procesan
	out1 := cuadrado(nums)

	// Fan-in: combinar resultados
	for n := range out1 {
		fmt.Printf("Cuadrado: %d\n", n)
	}

	// ===== DONE CHANNEL =====
	fmt.Println("\n=== DONE CHANNEL (CANCELACIÓN) ===")
	demoDoneChannel()

	// ===== CHANNEL CERRADO =====
	fmt.Println("\n=== CHANNEL CERRADO ===")

	ch3 := make(chan int, 2)
	ch3 <- 1
	ch3 <- 2
	close(ch3)

	// Recibir hasta que se cierre
	for valor := range ch3 {
		fmt.Println("Valor:", valor)
	}

	// Verificar si está cerrado
	ch4 := make(chan int, 1)
	ch4 <- 42
	close(ch4)

	valor, ok := <-ch4
	fmt.Printf("Valor: %d, Canal abierto: %t\n", valor, ok)

	valor, ok = <-ch4
	fmt.Printf("Valor: %d, Canal abierto: %t (cerrado)\n", valor, ok)

	// ===== BUENAS PRÁCTICAS =====
	fmt.Println("\n=== BUENAS PRÁCTICAS ===")
	fmt.Println("1. Solo el emisor debe cerrar el canal")
	fmt.Println("2. Usar range para recibir hasta que se cierre")
	fmt.Println("3. Usar select para manejar múltiples channels")
	fmt.Println("4. Usar canales buffered para evitar bloqueos")
	fmt.Println("5. Canales direccionales para mayor seguridad")
	fmt.Println("6. Usar done channel para cancelación")
	fmt.Println("7. Cerrar canales cuando no se necesiten más")
}
