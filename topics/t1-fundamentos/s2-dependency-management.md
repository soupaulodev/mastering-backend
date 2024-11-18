# Gerenciamento de Dependências

## Iniciar um módulo

Para criar um novo projeto, pode iniciar um módulo no diretório atual com o comando:

```sh
go mod init github.com/usuario/meu-projeto
```

## Adicionar dependências

Para utilizar um pacote externo, o Go Modules automaticamente busca e adiciona essa dependência ao arquivo go.mod.

```go
import "github.com/gorilla/mux"
```

Para instalar:

```go
go get github.com/gorilla/mux
```

Isso adiciona a dependência ao go.mod com a versão apropriada e atualiza o go.sum.

## Atualizar dependências

Para atualizar uma dependência específica:

```go
go get github.com/gorilla/mux@v1.8.0
```

Para atualizar todas as dependências:

```go
go get -u ./...
```

## Verificar dependências

Para verificar as dependências do projeto:

```go
go list -m all
```

## Remover dependências não usadas

```go
go mod tidy
```

Isso remove dependências que não estão sendo utilizadas no projeto e organiza os arquivos go.mod e go.sum.

## Baixar dependências

O comando `go mod` download baixa todas as dependências listadas no go.mod para o cache local.

As dependências baixadas são armazenadas em um cache local localizado no diretório:

- Linux/macOS: $GOPATH/pkg/mod
- Windows: %GOPATH%\pkg\mod

## Ambientes sem internet: Modo Vendoring

Para projetos que precisam funcionar sem acesso à internet, o Go permite o uso de vendor, um diretório que contém cópias das dependências. Para criar um diretório vendor, use:

```go
go mod vendor
```

Depois, ao compilar o projeto, use o comando:

```go
go build -mod=vendor
```
