# Fake Fintech

## Visão Geral

**Fake Fintech** é um projeto de aprendizado criado para explorar a arquitetura e as regras por trás de aplicações fintech. Ele demonstra como integrar APIs de terceiros, gerenciar tarefas assíncronas e construir um sistema robusto de acompanhamento financeiro utilizando tecnologias modernas.

## Arquitetura

O projeto está dividido em duas partes principais:

### 1. Aplicação Monolítica

- **Backend:** PHP (Laravel)
- **Frontend:** Angular 19 (template CoreUI)
- **Banco de Dados:** PostgreSQL
- **Cache:** Redis

### 2. Microsserviços

- **Mensageria:** RabbitMQ
- **Serviço Asset Watcher:** Golang (busca preços de ativos em múltiplas APIs)

#### Fluxo do Serviço

1. **Ação do Usuário:** Dispara uma tarefa (ex: buscar preço de criptomoeda).
2. **Fila de Tarefas:** A tarefa é enviada para o RabbitMQ (fila `assets-tasks`).
3. **Asset Watcher:** Serviço em Golang consome mensagens, busca preços no CoinCap, processa resultados e retorna o resultado.

## Objetivos atuais:
- **Dashboard**: Criar um dashboard limpo para que o usuário acesse seus dados facilmente
- **Dados financeiros do usuário**: quanto investe e utiliza no dia a dia

## Desenvolvimento Futuro
No futuro, este projeto irá expandir suas funcionalidades com recursos avançados como:

- **Transações**: Gerenciar e calcular transações financeiras.
- **Taxas de Juros**: Buscar e aplicar cálculos de juros.
- **Contas de Usuário**: Criar e gerenciar dados de usuários.
- **Cálculo de Empréstimos**: Implementar funcionalidades de cálculo de empréstimos.
- **Análise de Transações**: Fornecer insights detalhados sobre transações.
- **Gestão de Contas e Portfólios**: Gerenciar contas de usuários e seus portfólios.

---

### Fake Fintech - PT-br

[README em Português](./docs/pt-br/Readme.md)