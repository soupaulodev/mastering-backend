# Compilação e Execução de Programas Go

## Preparando o Ambiente

Para instalar o Go verifique o [site oficial](https://golang.org/dl/) e siga as instruções de acordo com seu sistema operacional.

Após a instalação, verifique a versão instalada e a configuração com:

```go
go version
go env
```

## Executar diretamente o programa

Go permite executar programas diretamente, sem a necessidade de compilá-los manualmente. Para isso, use o comando go run:

```go
go run nome-do-arquivo.go
```

## Compilação do programa

Para compilar um programa Go em um executável binário:

```go
go build nome-do-arquivo.go
```

Após esse comando, um arquivo executável será gerado no mesmo diretório, com o mesmo nome do arquivo Go (sem extensão). No caso do exemplo, o executável será:

- hello no Linux/macOS.
- hello.exe no Windows.

Para compilar o programa e especificar um nome para o binário:

```go
go build -o nome-desejado nome-do-arquivo.go
```

Após a compilação, você pode executar o binário diretamente:

```sh
./hello # Linux/macOS
hello.exe # Windows 5. Compilar e executar múltiplos arquivos
```

Se seu projeto tem múltiplos arquivos Go, o Go automaticamente os compila juntos. Por exemplo:

```go
go build
```

O comando acima compila todos os arquivos .go no diretório atual, criando um único executável.

## Compilação cruzada (Cross-Compilation)

Go suporta a compilação para diferentes sistemas operacionais e arquiteturas sem precisar de ferramentas adicionais. Isso é útil para criar binários que rodem em sistemas diferentes do seu ambiente de desenvolvimento.

Especificando o sistema operacional e arquitetura:
Antes de compilar, defina as variáveis de ambiente GOOS (sistema operacional) e GOARCH (arquitetura). Por exemplo:

```sh
GOOS=windows GOARCH=amd64 go build
```

Esse comando compila o programa para Windows 64 bits, mesmo que você esteja usando Linux ou macOS.

Valores comuns:

- GOOS: windows, linux, darwin (macOS), freebsd, etc.
- GOARCH: amd64 (64 bits), 386 (32 bits), arm, etc. 7. Executar testes

Go facilita a execução de testes com o comando go test:

```go
go test ./...
```

Esse comando encontra e executa todos os testes no projeto (arquivos que terminam com \_test.go).
