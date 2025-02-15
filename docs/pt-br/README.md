# Fake Fintech

## Introdução

O projeto **Fake Fintech** é uma iniciativa com o objetivo de entender como as aplicações de fintech funcionam, com foco nas regras e conceitos subjacentes que as impulsionam. Este projeto explora a integração com APIs conectadas ao ecossistema financeiro de fintechs, além de construir um sistema de rastreamento para que os usuários acompanhem suas finanças.

### Arquitetura do Projeto

Este projeto está estruturado com as seguintes tecnologias:

- **Backend**: Construído em **Golang** usando o framework **Gin**.
- **Frontend**: O frontend é baseado no template **CoreUI** Angular 19, proporcionando uma interface responsiva e modular para interação com o backend.
- **Banco de Dados**: **PostgreSQL** é utilizado para armazenar e gerenciar os dados financeiros.
- **Cache**: Usando **Redis**.

## Configuração do Backend

### Configuração do Ambiente

Para garantir a configuração correta do seu ambiente, o arquivo `backend/config/env/env.go` contém a configuração para a conexão com o PostgreSQL. Certifique-se de que as variáveis de ambiente estejam corretamente configuradas.

### Configuração do Backend

1. **Instale as dependências**:
   - Instale o [Go](https://golang.org/dl/) (Golang 1.18+ recomendado).
   - Execute o Docker Compose para configurar o ambiente do backend.

2. **Execute o backend**:
   - Certifique-se de que as variáveis de ambiente `POSTGRES_*` estejam corretamente configuradas.
   - Inicie o backend:
     ```bash
     go run main.go
     ```

## Configuração do Frontend

O frontend é construído em **Angular 19** e baseado no template **CoreUI**, oferecendo uma interface amigável para interagir com o backend.

### Configuração do Frontend

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/your-username/fake-fintech-frontend.git
   cd fake-fintech-frontend
   ```

2. **Instale as dependências**:
   - Instale o [Node.js](https://nodejs.org/en/download/) (recomendado v16 ou superior).
   - Instale as dependências do frontend:
     ```bash
     npm install
     ```

3. **Execute o frontend**:
   - Inicie a aplicação Angular:
     ```bash
     ng serve
     ```
   - O app estará acessível em `http://localhost:4200`.

## Desenvolvimento Futuro

No futuro, este projeto integrará com APIs de fintech de terceiros para obter dados financeiros em tempo real e expandir sua funcionalidade com recursos avançados, como:

- **Transações**: Gerenciar e calcular transações financeiras.
- **Taxas de Juros**: Recuperar e aplicar cálculos de juros.
- **Contas de Usuários**: Criar e gerenciar dados de usuários.
- **Cálculos de Empréstimos**: Implementar funcionalidades de cálculo de empréstimos.
- **Análise de Transações**: Fornecer insights detalhados sobre transações.
- **Gerenciamento de Contas e Portfólios**: Gerenciar contas de usuários e seus portfólios.