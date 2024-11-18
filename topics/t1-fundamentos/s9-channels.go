package t1fundamentos

import (
	"fmt"
	"time"
)

func TestChannels() {
	canalSimples()
	canalComBuffer()
	fechandoCanais()
	canalComSelect()
}

func canalSimples() {
	// Nesse caso, a goroutine envia uma mensagem para o canal e a função principal recebe a mensagem.
	// Porém, a goroutine fica bloqueada até que a mensagem seja enviada.

	var goroutineEnvia = func (ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "Mensagem enviada"
	}
	
		ch := make(chan string)
	
		go goroutineEnvia(ch) // Lança uma goroutine que envia dados para o canal
	
		// Recebe dados do canal
		mensagem := <-ch
		fmt.Println("Recebido no canal simples:", mensagem)
}

func canalComBuffer() {
	// Nesse caso, a goroutine não será bloqueada até que o buffer do canal esteja cheio.

	var goroutineEnvia = func(ch chan string) {
		for i := 1; i <= 3; i++ {
			msg := fmt.Sprintf("Mensagem %d", i)
			ch <- msg // Envia a mensagem para o canal
			fmt.Println("Enviado:", msg)
			time.Sleep(1 * time.Second)
		}
	}

	ch := make(chan string, 2) // Canal com buffer de tamanho 2

	go goroutineEnvia(ch) // Lança a goroutine que envia dados para o canal

	// Recebe os dados enviados
	for i := 1; i <= 3; i++ {
		msg := <-ch
		fmt.Println("Recebido no canal com buffer:", msg)
	}
}

func fechandoCanais() {
	// Canais podem ser fechados para indicar que não haverá mais envio de dados.
	// A função close() fecha o canal.

	var goroutineEnvia = func(ch chan string) {
		for i := 1; i <= 3; i++ {
			msg := fmt.Sprintf("Mensagem %d", i)
			ch <- msg // Envia a mensagem para o canal
			fmt.Println("Enviado:", msg)
			time.Sleep(1 * time.Second)
		}
		close(ch) // Fecha o canal após enviar todas as mensagens
	}

	ch := make(chan string, 2) // Canal com buffer de tamanho 2

	go goroutineEnvia(ch) // Lança a goroutine que envia dados para o canal

	// Recebe os dados enviados
	for msg := range ch {
		fmt.Println("Recebido no canal com fechamento:", msg)
	}
}

func canalComSelect() {
	// O select é usado para lidar com múltiplos canais.
	// Ele bloqueia até que um dos seus casos esteja pronto
	// e então executa o caso correspondente.

	var tarefa = func(ch chan string) {
		ch <- "Mensagem da tarefa"
	}

	ch1 := make(chan string)
	ch2 := make(chan string)

	go tarefa(ch1)
	go tarefa(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println("Recebido do canal 1 com select:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Recebido do canal 1 com select:", msg2)
	
	}
}