# Clean Architecture - Golang Expert

Para rodar essa aplicação siga os seguintes passos:

1. Na pasta `/cmd/ordersystem` tem um arquivo `.env` com as variaveis de ambiente.

2. Para subir o banco de dados mysql:

```
docker-compose up -d
```
3. Baixe todas as dependências do projeto rodando:

```
go mod tidy
```

4. Para rodar as migrations no banco de dados:

```
migrate -path=internal/infra/database/sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

5. Para subir a aplicação navegue até `/cmd/ordersystem` e execute

```
go run main.go wire_gen.go
```

Após subir o serviço, você poderá acessar as seguintes interfaces:

- **API REST**: http://localhost:8000
- **Servidor GraphQL**: http://localhost:8080
- **Serviço gRPC**: Porta 50051

## Documentação da API REST

A documentação das rotas do servidor HTTP está disponível no arquivo `./api/apis.http`.
