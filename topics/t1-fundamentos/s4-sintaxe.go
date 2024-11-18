package t1fundamentos

import (
	"fmt"
)

func TestSintaxe() {
	fmt.Println("\n\nSintaxe")
	tiposPrimitivos()
	variaveis()
	operadores()
	estruturasDeDados()
	controleDeFluxo()
	loops()
	funVisivel()
	funNaoVisivel()
	funcaoSemParametroERetorno()
	funcaoComParametroEUmRetorno(1)
	funcaoComMultiplosParametrosERetorno(1, 2, 3.0)
	funcIgnorandoValoresRetornadosNaChamada()
	funcaoComParametrosVariadicos(1, 2, 3, 4, 5)
	funcaoComFuncaoComoParametro(funcaoAnonima)
	funcaoAnonima()
	funcaoRetornandoFuncao()()
	funcaoAnonimaExecutadaAposDeclaracao()
	funcoesClosure()
}

func tiposPrimitivos() {
	fmt.Println("Tipos Primitivos")

	// inteiros
	var i int = 1
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1

	// Inteiros sem sinal
	var ui uint = 1
	var ui8 uint8 = 1
	var ui16 uint16 = 1
	var ui32 uint32 = 1
	var ui64 uint64 = 1
	var uiPtr uintptr = 1

	// Decimais - Ponto flutuante
	var f32 float32 = 1.0
	var f64 float64 = 1.0

	// Complexos - Formado por Parte real e imaginária
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i

	// Booleano
	var b bool = true

	// Strings - Cadeia de caracteres
	var s string = "Hello, World!"

	// Caractere Unicode (internamente é um inteiro de 32 bits)
	var c rune = 'a'

	fmt.Println(i, i8, i16, i32, i64)
	fmt.Println(ui, ui8, ui16, ui32, ui64, uiPtr)
	fmt.Println(f32, f64)
	fmt.Println(c64, c128)
	fmt.Println(b)
	fmt.Println(s)
	fmt.Println(c)
}

// Declaração de variavel explicita
var veglobal int = 1
func variaveis() {
	fmt.Println("Variáveis")
	// Declaração de variável local explicita
	var velocal int = 1

	// Declaração de variável local implicita (Inferência de tipo)
	vilocal := 2

	// Atribuição de valor a variável
	veglobal = 3

	// Declaração de Variável com múltiplos valores
	v1, v2 := 1, 2

	// Declaração de variável com múltiplos valores de tipos diferentes
	v3, v4 := 1, "Hello, World!"

	// Declaração de Constante
	const c1 int = 1


	fmt.Println(velocal, vilocal, veglobal)
	fmt.Println(v1, v2)
	fmt.Println(v3, v4)
	fmt.Println(c1)

	// Variáveis não inicializadas recebem um valor padrão, chamado zero value:
	// int: 0
	// float: 0.0
	// string: "" (string vazia)
	// bool: false
}

func operadores() {
	fmt.Println("operadores")

	// Aritméticos
	fmt.Println("Aritméticos")
	var a, b int = 1, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)

	// Atribuição
	fmt.Println("Atribuição")
	var c int = 1
	c += 2
	fmt.Println(c)

	// Comparação
	fmt.Println("Comparação")
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)

	// Lógicos
	fmt.Println("Lógicos")
	var d, e bool = true, false
	fmt.Println(d && e)
	fmt.Println(d || e)
	fmt.Println(!d)

	// Bitwise
	fmt.Println("Bitwise")
	var f, g int = 1, 2
	fmt.Println(f & g)
	fmt.Println(f | g)
	fmt.Println(f ^ g)
	fmt.Println(f << 1)
	fmt.Println(f >> 1)
}

func estruturasDeDados() {
	fmt.Println("Estruturas de Dados")

	// Arrays - Coleção fixa de elementos do mesmo tipo
	fmt.Println("Arrays")
	var a [5]int
	fmt.Println(a)

	// Slices - Coleção dinâmica baseada em arrays
	fmt.Println("Slices")
	var s []int
	fmt.Println(s)
	// Implicitamente
	s2 := make([]int, 5)
	fmt.Println(s2)

	// Maps - Coleção de pares chave-valor
	fmt.Println("Maps")
	var m map[string]int
	fmt.Println(m)
	// Implicitamente
	m2 := make(map[string]int)
	fmt.Println(m2)

	// Structs - Agrupamento de dados em uma estrutura personalizada.
	fmt.Println("Structs")
	type Person struct {
		Name string
		Age  int
	}
	var p Person
	fmt.Println(p)
}

func controleDeFluxo() {
	fmt.Println("Controle de Fluxo")

	// If
	fmt.Println("If")
	var a, b int = 1, 2
	if a > b {
		fmt.Println("a é maior que b")
	} else if (a == b) {
		fmt.Println("a é igual a b")
	} else {
		fmt.Println("b é maior que a")
	}

	// Switch
	fmt.Println("Switch")
	var c int = 1
	switch c {
	case 1:
		fmt.Println("c é igual a 1")
	case 2:
		fmt.Println("c é igual a 2")
	default:
		fmt.Println("c é diferente de 1 e 2")
	}
}

func loops() {
	// For
	fmt.Println("For")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// For com range
	fmt.Println("For com range")
	var s = []int{1, 2, 3, 4, 5}
	for i, v := range s {
		fmt.Println(i, v)
	}

	// For (While)
	fmt.Println("For (While)")
	var j int = 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// For (Do While)
	fmt.Println("For (Do While)")
	var k int = 0
	for {
		fmt.Println(k)
		k++
		if k >= 5 {
			break
		}
	}

	// Iterando sobre um Slice ou Map
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// For infinito
	fmt.Println("For infinito")
	for {
		fmt.Println("Loop infinito")
		break
	}
}

// Visibilidade
// - Inicia com letra maiúscula: Pública
// - Inicia com letra minúscula: Privada
var VariavelVisivelExt int = 1
var variavelNaoVisivelExt int = 1
func funVisivel() {
	fmt.Println(variavelNaoVisivelExt) // Utilizando apenas para evitar erro de compilação
	fmt.Println("Função visível externamente")
}
func funNaoVisivel() {
	fmt.Println("Função não visível externamente")
}

func funcaoSemParametroERetorno() {
	fmt.Println("Função sem parâmetro e retorno")
}

func funcaoComParametroEUmRetorno(a int) int {
	fmt.Println("Função com parâmetro e retorno")
	return a
}

func funcaoComMultiplosParametrosERetorno(a, b int, c float32) (int, int, float32) {
	fmt.Println("Função com múltiplos parâmetros e retorno")
	return a + b, a - b, c
}

func funcIgnorandoValoresRetornadosNaChamada() {
	fmt.Println("Função ignorando valores retornados na chamada")
	_, b, _ := funcaoComMultiplosParametrosERetorno(1, 2, 3.0)
	fmt.Println(b)
}

func funcaoComParametrosVariadicos(a ...int) {
	fmt.Println("Função com parâmetros variádicos")
	for i, v := range a {
		fmt.Println(i, v)
	}
}

func funcaoComFuncaoComoParametro(f func()) {
	fmt.Println("Função com função como parâmetro")
	f()
}

var funcaoAnonima = func() {
	fmt.Println("Função anônima")
}

func funcaoRetornandoFuncao() func() {
	fmt.Println("Função anônima")
	return func() {
		fmt.Println("Função anônima executada")
	}
}

func funcaoAnonimaExecutadaAposDeclaracao() {
	func () {
		fmt.Println("Função anônima executada após declaração")
	}()
}

func funcoesClosure() {
	fmt.Println("Funções Closure")
	var a int = 1
	var f = func() {
		fmt.Println(a)
	}
	f()
}



