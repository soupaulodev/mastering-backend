package t1fundamentos

import (
	"fmt"
	"sync"
)

func TestSyncPackages() {
	mutexesImpl()
	waitGroupImpl()
	onceImpl()
	condImpl()
	poolImpl()
}

func mutexesImpl() {
	// Um mutex é um tipo de bloqueio utilizado para proteger recursos compartilhados, garantindo que apenas uma goroutine possa acessar um recurso de cada vez.

	var contador int
	var mutex sync.Mutex

	var incrementar = func() {
		mutex.Lock()         // Bloqueia o mutex
		contador++           // Acessa o recurso compartilhado
		mutex.Unlock()       // Libera o mutex
	}

	var wg sync.WaitGroup

	// Lançando 10 goroutines para incrementar o contador
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			incrementar()
		}()
	}

	wg.Wait()  // Espera todas as goroutines terminarem
	fmt.Println("Contador final usando mutex:", contador)
}

func waitGroupImpl() {
	// O sync.WaitGroup é usado para aguardar que um conjunto de goroutines termine suas execuções.
	// Ele permite que o programa principal ou qualquer outra goroutine aguarde que várias goroutines completem suas tarefas.
	var tarefa = func(id int, wg *sync.WaitGroup) {
		defer wg.Done()  // Chama Done quando a tarefa for concluída
		fmt.Printf("Tarefa %d iniciada\n", id)
	}
	
	var wg sync.WaitGroup

	// Lançando 5 goroutines
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrementa o contador de goroutines a serem aguardadas
		go tarefa(i, &wg)
	}

	wg.Wait()  // Espera todas as goroutines terminarem
	fmt.Println("Todas as tarefas foram concluídas!")
}

func onceImpl() {
	// A sync.Once garante que uma determinada ação seja executada apenas uma vez, independentemente de quantas vezes seja chamada.
	// Isso é útil para inicializações que só devem ocorrer uma vez, como configurar recursos globais, banco de dados, etc.
	var once sync.Once

	var inicializar = func() {
		fmt.Println("Inicialização realizada")
	}

	// Chama a inicialização várias vezes, mas ela ocorrerá apenas uma vez
	for i := 0; i < 3; i++ {
		once.Do(inicializar)
	}

	fmt.Println("Programa finalizado")
}

func condImpl() {
	// A sync.Cond é uma variável de condição que aguarda ou notifica outras goroutines.
	// Ela é útil para sincronizar a execução

	var cond = sync.NewCond(&sync.Mutex{})
	var contador int

	var aumentar = func() {
		cond.L.Lock()
		contador++
		if contador == 5 {
			cond.Signal()  // Notifica a goroutine que está esperando
		}
		cond.L.Unlock()
	}

	var esperar =func() {
		cond.L.Lock()
		for contador < 5 {
			cond.Wait()  // Espera a condição ser sinalizada
		}
		fmt.Println("Contador atingiu 5!")
		cond.L.Unlock()
	}


	go aumentar()
	go aumentar()
	go aumentar()
	go aumentar()
	go aumentar()

	esperar()
}

func poolImpl() {
	/*
		O sync.Pool é uma estrutura fornecida pelo pacote sync do Go que permite o reaproveitamento de objetos para otimizar a alocação de memória e reduzir a pressão sobre o garbage collector (coletor de lixo). Ele é útil quando você tem objetos caros para criar ou objetos que são frequentemente reutilizados em um sistema concorrente. O sync.Pool oferece um mecanismo de "pool" de objetos, ou seja, você pode armazenar objetos em um pool e reutilizá-los, em vez de alocá-los e desalocá-los repetidamente.
	*/

	// Criando um pool com objetos do tipo string
	var pool = sync.Pool{
		New: func() interface{} {
			// Função que cria um novo objeto quando o pool está vazio
			obj := "Novo objeto"
			return &obj  // Retorna um ponteiro para string
		},
	}

	// Pegando um objeto do pool e usando como ponteiro
	obj1 := pool.Get().(*string)
	fmt.Println("Primeiro objeto:", *obj1) // Exibe "Novo objeto"

	// Preparando um objeto para ser colocado de volta no pool
	obj := "Objeto reutilizado"
	pool.Put(&obj)  // Coloca um ponteiro para string no pool

	// Pegando um objeto reutilizado do pool
	obj2 := pool.Get().(*string)
	fmt.Println("Segundo objeto:", *obj2) // Exibe "Objeto reutilizado"

	// Pegando novamente do pool (podendo ser reutilizado ou o novo)
	obj3 := pool.Get().(*string)
	fmt.Println("Terceiro objeto:", *obj3) // Exibe "Novo objeto"
}