# Fake Fintech

## Introduction

The **Fake Fintech** project is an initiative aimed at understanding how fintech applications work, with a focus on the underlying rules and concepts that drive them. This project explores integrating APIs connected to the economic fintech ecosystem while building a tracking system for users to keep up with their finances.

### Project Architecture

This project is structured with the following technologies:

- **Backend**: Built in **Golang** using the **Gin** framework.
- **Frontend**: The frontend is based on the **CoreUI** Angular 19 template, providing a responsive and modular UI for interacting with the backend.
- **Database**: **PostgreSQL** is used to store and manage financial data.
- **Cache**: Using **Redis**.

## Backend Setup

### Environment Configuration

To ensure the proper setup of your environment, the `backend/config/env/env.go` file contains the configuration for connecting to PostgreSQL. Make sure your environment variables are set correctly.

Example of `env.go` file:

```go
package env

import (
	"fmt"
	"os"
)

func setEnv() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	if host == "" || port == "" || user == "" || password == "" || database == "" {
		panic("Missing required environment variables for PostgreSQL connection")
	}

	os.Setenv("POSTGRES_HOST", host)
	os.Setenv("POSTGRES_PORT", port)
	os.Setenv("POSTGRES_USER", user)
	os.Setenv("POSTGRES_PASSWORD", password)
	os.Setenv("POSTGRES_DATABASE", database)
}

func GetPostgresConnectionString() string {
	setEnv()
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	connectionStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	return connectionStr
}

func GetDatabase() string {
	setEnv()
	return os.Getenv("POSTGRES_DATABASE")
}

func SetSalt() []byte {
	your_salt := "MySalt"
	return []byte(your_salt)
}
```

### Setting Up the Backend

1. **Install dependencies**:
   - Install [Go](https://golang.org/dl/) (Golang 1.18+ recommended).
   - Run the Docker Compose file to set up the backend environment.

2. **Run the backend**:
   - Ensure that the `POSTGRES_*` environment variables are correctly set (as shown in the `env.go` file).
   - Start the backend service:
     ```bash
     go run main.go
     ```

## Frontend Setup

The frontend is built using **Angular 19** and based on the **CoreUI** template, offering a user-friendly interface to interact with the backend.

### Setting Up the Frontend

1. **Clone the repository**:
   ```bash
   git clone https://github.com/your-username/fake-fintech-frontend.git
   cd fake-fintech-frontend
   ```

2. **Install dependencies**:
   - Install [Node.js](https://nodejs.org/en/download/) (v16 or later recommended).
   - Install frontend dependencies:
     ```bash
     npm install
     ```

3. **Run the frontend**:
   - Start the Angular application:
     ```bash
     ng serve
     ```

   - The app will be accessible at `http://localhost:4200`.

## Future Development

In the future, this project will integrate with third-party fintech APIs to retrieve real-time financial data and expand its functionality with advanced features like:

- **Transactions**: Manage and calculate financial transactions.
- **Interest Rates**: Retrieve and apply interest calculations.
- **User Accounts**: Create and manage user data.
- **Loan Calculations**: Implement loan calculation functionalities.
- **Transaction Analytics**: Provide detailed insights on transactions.
- **Account and Portfolio Management**: Manage user accounts and their portfolios.

---

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

No futuro, este projeto integrará APIs financeiras para buscar dados em tempo real e expandir suas funcionalidades com recursos avançados como análise de transações, cálculos de empréstimos e gerenciamento de portfólio.

