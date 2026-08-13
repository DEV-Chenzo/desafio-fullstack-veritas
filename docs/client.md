# Client — Veritas Mini Kanban

Interface web do Mini Kanban, construída com React, Vite e Tailwind CSS. O client consome a API Go e mantém a tela sincronizada após cada operação de criação, edição, movimentação ou exclusão.

## Responsabilidades

- Exibir as colunas **A Fazer**, **Em Progresso** e **Concluído**.
- Criar uma tarefa com título, descrição opcional e status inicial.
- Editar e excluir cards.
- Mover cards para a coluna anterior ou seguinte.
- Exibir estado de carregamento e erros de comunicação com a API.

## Organização

```text
client/src/
├── components/
│   ├── KanbanBoard.jsx  # Formulário e composição das colunas
│   ├── Column.jsx       # Coluna, edição inline e movimentação
│   └── TaskCard.jsx     # Visualização e ações de uma tarefa
├── services/api.js      # Cliente Axios e operações HTTP
├── App.jsx              # Estado das tarefas e sincronização com a API
├── index.css            # Estilos globais e componentes Tailwind
└── main.jsx             # Ponto de entrada React
```

## Execução local

Pré-requisito: Node.js 18 ou superior. A API e o banco devem estar em execução conforme o [README da raiz](../README.md).

```bash
cd client
npm install
npm run dev
```

O Vite exibirá a URL local, normalmente `http://localhost:5173`.

## Variáveis de ambiente

A API padrão é `http://localhost:8080/api`. Para alterá-la, crie `client/.env.local`:

```env
VITE_API_URL=http://localhost:8080/api
```

Após mudar uma variável `VITE_*`, reinicie o servidor de desenvolvimento.

## Scripts

| Comando | Descrição |
| --- | --- |
| `npm run dev` | Inicia o ambiente de desenvolvimento. |
| `npm run build` | Gera a versão de produção em `dist/`. |
| `npm run preview` | Serve localmente o build de produção. |

## Comunicação com a API

O arquivo `src/services/api.js` expõe `taskService`, usado pelo `App.jsx`:

| Operação | Método | Rota |
| --- | --- | --- |
| Listar | `getAll` | `GET /tasks` |
| Criar | `create` | `POST /tasks` |
| Atualizar | `update` | `PUT /tasks/{id}` |
| Excluir | `delete` | `DELETE /tasks/{id}` |

O prefixo `/api` é definido pela variável `VITE_API_URL`.
