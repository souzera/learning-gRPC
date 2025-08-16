# learning-gRPC

- Google Remote Procedure Call
- é um framework open-source de RPC
- comunicação eficiente e de alta performance
- utiliza p *HTTP/2.0* para transportar "protocol buffers"
- fornece suporte nativo para várias linguagens de programação

## Preparando o Ambiente no GO

Instale os plugins de compilação

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Atualize o seu `path` para o compilador encontrar os plugins

```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

## Como Usar

1. Escreva o contrato do `*.proto`
2. Execute o comando

```bash
protoc --go_out=. --go-grpc_out=. ./*.proto
```
para gerar o 
