package t1fundamentos

import (
	"fmt"
	"time"
)

func tarefa(nome string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(nome, "executando a tarefa", i)
		time.Sleep(1 * time.Second) // Simula tempo de execução
	}
}

func TestGourotines() {
	// Lançando duas goroutines
	go tarefa("Goroutine 1") // Lança a goroutine para "Goroutine 1"
	go tarefa("Goroutine 2") // Lança a goroutine para "Goroutine 2"
	
	// O programa principal vai esperar um pouco para as goroutines terminarem
	time.Sleep(4 * time.Second)
	fmt.Println("Execução principal concluída.")
}