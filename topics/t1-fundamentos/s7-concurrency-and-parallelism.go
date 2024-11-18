package t1fundamentos

import (
	"fmt"
	"runtime"
	"time"
)

func TestConcorrenciaEParalelismo() {
	concorrencia()
	paralelismo()
}

func concorrencia() {
	/*
		Concorrência é a capacidade de gerenciar múltiplas tarefas ao mesmo tempo
		(mas não necessariamente ao mesmo instante),
		alternando entre elas para que todas possam progredir.

		Exemplo: Você está assistindo a um filme e, ao mesmo tempo,
		está conversando com um amigo no WhatsApp. Você não consegue fazer as duas coisas
		ao mesmo tempo, mas consegue alternar entre elas rapidamente.
	*/

	// Concorrência em Go com Goroutines (sem paralelismo explícito):
	var tarefa = func (nome string) {
		for i := 0; i < 3; i++ {
			fmt.Println(nome, "executando tarefa concorrente", i)
			time.Sleep(time.Second)
		}
	}

	go tarefa("Tarefa 1")  // Goroutine concorrente
	go tarefa("Tarefa 2")  // Outra goroutine concorrente

	time.Sleep(5 * time.Second) // Espera o tempo necessário para as goroutines terminarem
}

func paralelismo() {
	/*
		O paralelismo é sobre executar múltiplas tarefas de forma simultânea.

		Exemplo: Você está assistindo a um filme e, ao mesmo tempo, seu amigo está
		assistindo a outro filme em uma tela diferente. Ambos estão assistindo a filmes
		ao mesmo tempo.
	*/

	// Paralelismo em Go com Goroutines (se múltiplos núcleos estiverem disponíveis):
	var tarefa = func (nome string) {
		for i := 0; i < 3; i++ {
			fmt.Println(nome, "executando tarefa paralela", i)
			time.Sleep(time.Second)
		}
	}

	runtime.GOMAXPROCS(2) // Define o número máximo de núcleos de CPU para a execução

	go tarefa("Tarefa 1")
	go tarefa("Tarefa 2")

	time.Sleep(5 * time.Second) // Espera o tempo necessário para as goroutines terminarem
}