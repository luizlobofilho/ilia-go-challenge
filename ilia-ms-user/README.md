# Users Microservice Boilerplate (Go + Gin + Postgres)

Este é um boilerplate mínimo do Microserviço Users, baseado no padrão do projeto Wallet.

## Como rodar o serviço

### Pré-requisitos
- Docker e Docker Compose instalados
- Go 1.25+ instalado (opcional, para desenvolvimento local)

### Subindo com Docker Compose
No diretório `ilia-ms-user` execute:

```bash
docker-compose up --build
```

O serviço será exposto por padrão na porta `3002`.

### Testes
Para rodar os testes unitários locais:

```bash
cd ilia-ms-user
go test ./...
```

### Migrations
Este boilerplate inclui um exemplo de migration para a tabela `users` em `migrations/`.
Usamos o formato compatível com `golang-migrate`.

Exemplo para aplicar as migrations localmente (ajuste a DSN conforme seu ambiente):

```bash
migrate -path ./migrations -database "postgres://usersuser:userspass@localhost:5434/usersdb?sslmode=disable" up
```

### Variáveis de ambiente principais
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Configuração do Postgres
- `JWT_SECRET`: Chave do JWT (use `ILIACHALLENGE` para desenvolvimento)
- `PORT`: Porta do serviço (default: `3002`)

### Endpoint disponível (inicial)
- `GET /users/:id` (protegido por JWT)

Exemplo de requisição curl:

```bash
curl -H "Authorization: Bearer <token>" http://localhost:3002/users/1
```

### Gerar token JWT de teste (exemplo em Go)

```go
package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func main() {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": "1"})
	t, _ := token.SignedString([]byte("ILIACHALLENGE"))
	fmt.Println(t)
}
```

### Observações
- A autenticação JWT espera o token assinado com a chave definida em `JWT_SECRET`.
- O repositório usa GORM para acesso ao banco; a implementação atual usa Postgres via DSN montada a partir das variáveis de ambiente.

---
