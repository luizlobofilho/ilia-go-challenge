# Wallet Microservice Boilerplate (Go + Gin + Postgres + Migrations)

## Como rodar o projeto

### Pré-requisitos
- Docker e Docker Compose instalados
- Go 1.25+ instalado (opcional, para desenvolvimento local)

### Subindo com Docker Compose
```bash
docker-compose up --build
```

### Migrations
O projeto usa [golang-migrate](https://github.com/golang-migrate/migrate) para versionamento do banco.

Para rodar as migrations manualmente:
```bash
# Exemplo (ajuste o caminho conforme necessário)
migrate -path ./migrations -database "postgres://walletuser:walletpass@localhost:5433/walletdb?sslmode=disable" up
```

### Variáveis de ambiente principais
- `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`: Configuração do Postgres
- `JWT_SECRET`: Chave do JWT (use ILIACHALLENGE)
- `PORT`: Porta do serviço (default: 3001)

### Endpoints
- `GET /transactions` (protegido por JWT)

---
