# User Flow — Veritas Mini Kanban

Este documento descreve o caminho principal de uso da aplicação. A interface foi pensada para que cada ação tenha uma consequência visual imediata e persista no PostgreSQL por meio da API.

```mermaid
flowchart TD
  A([Usuário acessa o Kanban]) --> B[Client solicita tarefas à API]
  B --> C[Quadro exibe A Fazer, Em Progresso e Concluído]
  C --> D{Qual ação deseja realizar?}

  D -->|Criar| E[Informa título, descrição opcional e coluna]
  E --> F[Envia POST /api/tasks]
  F --> G{Dados válidos?}
  G -->|Sim| H[Tarefa é salva no PostgreSQL]
  H --> I[Card aparece na coluna selecionada]
  G -->|Não| J[Interface exibe mensagem de erro]

  D -->|Editar| K[Abre edição inline no card]
  K --> L[Altera título ou descrição e salva]
  L --> M[Envia PUT /api/tasks/id]
  M --> N[Card é atualizado]

  D -->|Mover| O[Seleciona seta anterior ou próxima]
  O --> P[Envia PUT com novo status]
  P --> Q[Card muda de coluna]

  D -->|Excluir| R[Seleciona excluir]
  R --> S[Envia DELETE /api/tasks/id]
  S --> T[Card deixa o quadro]

  I --> C
  J --> C
  N --> C
  Q --> C
  T --> C
```

## Regras do fluxo

- O título é obrigatório e aceita até 120 caracteres.
- A descrição é opcional.
- Os únicos status aceitos são `todo`, `doing` e `done`.
- A movimentação segue a ordem **A Fazer → Em Progresso → Concluído**.
- Os botões de movimentação foram escolhidos em vez de drag and drop para preservar previsibilidade em telas pequenas e uso por teclado.
- Após recarregar a página, o client busca novamente as tarefas persistidas no banco.
