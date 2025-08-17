# learning-gRPC

- Google Remote Procedure Call
- é um framework open-source de RPC
- comunicação eficiente e de alta performance
- utiliza p *HTTP/2.0* para transportar "protocol buffers"
- fornece suporte nativo para várias linguagens de programação

### HTTP/2.0

O **HTTP/2.0** é a evolução do **HTTP/1.1**, foi projetado para ser mais rápido, eficiente e compatível com aplicações modernas. 

No HTTP/1.1 existem algumas limitações...
1. *HEAD-of-Line Blocking*: uma requisição lenta bloqueia as outras
2. *Multiplas conexões TCP*: cada conexão TCP só pode processar uma única requisição/resposta por vez, então faz-se necessário a abertura de múltiplas conexões TCPs paralelas
3. *Headers Duplicados*: desperdício de largura de banda
4. *Sem server push*: o cliente sempre inicia a comunicação

Então, no HTTP/2.0 possui propostas para solucionar esses problemas.
1. **Multiplexação**: uma única conexão TCP pode carregar múltiplas conexões, eliminando o *HEAD-of-Line Blocking*, reduzindo a latência, otimização da conexão
2. **Protocolo Binário**:

No HTTP/1.1
```text
GET /api/users HTTP/1.1
Host: example.com
Accept: application/json
User-Agent: MyApp/1.0
```

No HTTP/2.0
```text
┌─────────────┬────────────┬──────────────┐
│   Header    │   Method   │     Path     │
│   Block     │    GET     │  /api/users  │
└─────────────┴────────────┴──────────────┘
```

O protocolo binário é um formato mais compacto, tem o parsing mais rápido e menos propenso a erros.

3. **Header Compression(HPACK)**:

HTTP/1.1
```
Request 1: 
Host: api.example.com 
User-Agent: MyApp/1.0 
Accept: application/json 

Request 2: 
Host: api.example.com ← Duplicado! 
User-Agent: MyApp/1.0 ← Duplicado! 
Accept: application/json ← Duplicado!
```

HTTP/2 com HPACK:
```
Request 1: [Headers completos]
Request 2: [Referências]: :method=GET :path=/new-endpoint
```

Reduz drasticamente a quantidade de bytes dos headers.

4. **Server Push**: O servidor pode enviar recursos antes do cliente pedir
5. **Stream Priorization**: o cliente pode definir prioridades

Com HTTP/1.1, cada chamada gRPC precisaria de uma conexão TCP separada. Com HTTP/2, todas as chamadas compartilham uma única conexão, sendo muito mais eficiente!

### Protocol Buffers

[Documentação](https://protobuf.dev/)

*O que são os Protocol Buffers?*

Protocol Buffers é uma tecnologia desenvolvida pelo Google para serializar e desserializar dados estruturados independente de linguagem e versão.

**Comparação com outros formatos**

| Formato      | Tamanho | Velocidade | Tipagem | Legibilidade |
| ------------ | ------- | ---------- | ------- | ------------ |
| **JSON**     | Grande  | Lenta      | Fraca   | Alta         |
| **XML**      | Maior   | Mais Lenta | Fraca   | Alta         |
| **Protobuf** | Pequeno | Rápido     | Forte   | Baixa        |

- tipos de dados
- tipos de streaming
- versionamento
- campos opcionais

#### Pontos importantes

**Tags Numéricas**: São a chave da compatibilidade, **nunca as reutilize**

```protobuf
message User { 
	string name = 1; // Tag 1 será sempre "name" 
	int32 age = 2; // Tag 2 será sempre "age" 
	// string email = 3; // ❌ NUNCA reutilize tag 3! 
}
```

## Arquivo .proto

Estrutura do *.proto

```proto
syntax = "proto3";

package example;

service ExampleService {
	rpc GetExample (ExampleRequest) returns (ExampleResponse)
}

message ExampleRequest {
	string id = 1;
}

message ExampleResponse {
	string message = 1;
}
```

campo option:
```go
option go_package = "/example";
```

para gerar
```bash
protoc --go_out=. --go-grpc_out=. ./example.proto
```


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

## Executando gRPC

1. Escreva o contrato do `*.proto`
2. Execute o comando

```bash
protoc --go_out=. --go-grpc_out=. ./*.proto
```
para gerar os arquivos `*_grpc.pb.go` e `*.pb.go`

3. Execute o servidor

```bash
go run server/main.go
```

4. Em outro terminal o cliente

```bash
go run client/main.go
```


