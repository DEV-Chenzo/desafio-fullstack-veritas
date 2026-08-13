# Veritas Mini Kanban

Aplicação fullstack para organizar tarefas em um fluxo Kanban. O projeto entrega CRUD completo, persistência em PostgreSQL, movimentação entre colunas e uma interface responsiva feita com React, Vite e Tailwind CSS.

## Demonstração do fluxo

O diagrama completo está em [docs/user-flow.md](docs/user-flow.md). Em resumo: o usuário cria uma tarefa, acompanha-a nas colunas **A Fazer**, **Em Progresso** e **Concluído**, podendo editar, mover e excluir a qualquer momento.

## Tecnologias

| Camada | Escolha | Motivo |
| --- | --- | --- |
| Client | React + Vite + Tailwind CSS | Componentização, feedback imediato e interface responsiva. |
| API | Go (net/http) | API leve, explícita e com poucas dependências. |
| Dados | PostgreSQL 16 | Persistência relacional confiável e adequada ao domínio. |
| Ambiente | Docker Compose | Banco reproduzível com um único comando. |

## Estrutura

```text
.
├── client/             # Interface React (detalhes em docs/client.md)
├── server/             # API REST em Go (detalhes em docs/server.md)
│   └── internal/       # Configuração, banco, HTTP e domínio de tarefas
├── docs/               # Documentação complementar
│   ├── client.md       # Client React
│   ├── server.md       # API Go e PostgreSQL
│   └── user-flow.md    # Diagrama e decisões de experiência
├── docker-compose.yml  # PostgreSQL local
└── .env.example        # Variáveis da API
```

## Como executar

Pré-requisitos: Docker Desktop, Go 1.22+ e Node.js 18+.

1. Inicie o banco de dados na raiz do projeto:

```bash
docker compose up -d
```

2. Em outro terminal, inicie a API:

```bash
cd server
go mod download
go run .
```

3. Em um terceiro terminal, inicie o cliente:

```bash
cd client
npm install
npm run dev
```

Abra `http://localhost:5173`. A API estará em `http://localhost:8080/api`; na inicialização ela cria a tabela `tasks` automaticamente, de forma idempotente.

### Configuração opcional

O backend usa por padrão `postgres://kanban:kanban@127.0.0.1:5433/kanban?sslmode=disable`. A porta `5433` evita conflito com uma instalação PostgreSQL local que possa estar usando a porta padrão `5432`. Consulte também [docs/server.md](docs/server.md) para o detalhamento da API. Para alterar, copie `.env.example` e defina as variáveis no ambiente antes de executar o Go:

- `DATABASE_URL`: conexão PostgreSQL.
- `PORT`: porta da API (padrão `8080`).
- `CORS_ORIGIN`: origem permitida para o cliente (padrão `http://localhost:5173`).

Para apontar o client para outra API, crie `client/.env.local` com `VITE_API_URL=http://localhost:8080/api`.

## Funcionalidades

- Criar tarefas com título obrigatório, descrição opcional e coluna inicial.
- Listar tarefas persistidas após recarregar a página ou reiniciar a API.
- Editar título e descrição diretamente no card.
- Mover tarefas entre os estágios pelos controles de direção.
- Excluir tarefas.
- Indicador de quantidade por coluna, estados vazios e mensagens de erro.

## API REST

Documentação específica do backend: [docs/server.md](docs/server.md).

| Método | Rota | Descrição |
| --- | --- | --- |
| `GET` | `/api/health` | Verifica conexão da API e do banco. |
| `GET` | `/api/tasks` | Lista tarefas por data de criação. |
| `POST` | `/api/tasks` | Cria uma tarefa. |
| `GET` | `/api/tasks/{id}` | Busca uma tarefa. |
| `PUT` | `/api/tasks/{id}` | Atualiza título, descrição e status. |
| `DELETE` | `/api/tasks/{id}` | Remove uma tarefa. |

Exemplo de criação:

```json
{
  "title": "Preparar apresentação",
  "description": "Consolidar os resultados da semana",
  "status": "todo"
}
```

Os valores aceitos de `status` são `todo`, `doing` e `done`. O backend valida título, status, JSON e identificador; erros retornam JSON com a propriedade `error`.

## Decisões de arquitetura

- **Tabela única para o MVP:** tarefas são o único agregado do domínio; evitar relações prematuras mantém a solução legível.
- **Migração na inicialização:** torna o primeiro uso simples, sem comprometer a idempotência da execução local.
- **Atualização completa (`PUT`):** o client sempre envia o estado integral e elimina ambiguidades entre campo vazio e campo ausente.
- **Movimentação por botões:** mantém o fluxo claro, acessível por teclado e funcional em desktop e mobile sem depender de drag and drop.
- **Separação por componentes:** `KanbanBoard`, `Column` e `TaskCard` mantêm responsabilidades visuais pequenas; `App` concentra a sincronização com a API.

## Encerrando o ambiente

```bash
docker compose down
```

Os dados permanecem no volume `postgres_data`. Para removê-los intencionalmente, execute `docker compose down -v`.
