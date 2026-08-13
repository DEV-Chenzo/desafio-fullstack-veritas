# Server — Veritas Mini Kanban

API REST em Go responsável por validar, persistir e disponibilizar tarefas do quadro Kanban.

## Arquitetura

```text
server/
├── main.go
└── internal/
    ├── config/       # Carrega variáveis de ambiente e valores padrão
    ├── database/     # Cria o pool pgx e aplica a migração idempotente
    ├── httpapi/      # Registra rotas, handlers, respostas e middlewares
    └── task/         # Modelo, regras de negócio e repositório PostgreSQL
```

O fluxo de uma requisição é:

```text
HTTP request → handler → service → repository → PostgreSQL
```

- O **handler** interpreta HTTP, JSON e parâmetros de rota.
- O **service** normaliza e valida as regras de negócio.
- O **repository** contém somente as consultas de persistência.

## Execução

Pré-requisito: o PostgreSQL do Docker em execução na raiz do repositório.

```bash
cd server
go mod download
go run .
```

A API inicia em `http://localhost:8080/api`.

## Variáveis de ambiente

| Variável | Padrão | Descrição |
| --- | --- | --- |
| `DATABASE_URL` | `postgres://kanban:kanban@127.0.0.1:5433/kanban?sslmode=disable` | URL do PostgreSQL. |
| `PORT` | `8080` | Porta HTTP da API. |
| `CORS_ORIGIN` | `http://localhost:5173` | Origem permitida para o frontend. |

A porta externa `5433` evita conflito com uma instalação local do PostgreSQL na porta padrão `5432`.

## Rotas

| Método | Rota | Resposta de sucesso |
| --- | --- | --- |
| `GET` | `/api/health` | `200` com status da API e banco. |
| `GET` | `/api/tasks` | `200` com a lista de tarefas. |
| `POST` | `/api/tasks` | `201` com a tarefa criada. |
| `GET` | `/api/tasks/{id}` | `200` com a tarefa. |
| `PUT` | `/api/tasks/{id}` | `200` com a tarefa atualizada. |
| `DELETE` | `/api/tasks/{id}` | `204` sem corpo. |

### Corpo para criar ou atualizar

```json
{
  "title": "Preparar apresentação",
  "description": "Consolidar os resultados da semana",
  "status": "todo"
}
```

`title` é obrigatório; `description` é opcional; `status` pode ser `todo`, `doing` ou `done`. Erros de validação retornam `400` e recursos inexistentes retornam `404`, sempre com JSON no formato `{"error":"mensagem"}`.

## Banco de dados

Na inicialização, a aplicação cria a tabela `tasks` caso ela ainda não exista. A operação é idempotente e adequada ao escopo local do desafio.

```sql
tasks (
  id BIGSERIAL PRIMARY KEY,
  title VARCHAR(120) NOT NULL,
  description TEXT NOT NULL,
  status VARCHAR(10) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
)
```

## Verificação

```bash
cd server
go test ./...
go vet ./...
```
