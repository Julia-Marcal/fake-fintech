
# Fake Fintech

## Introduction

The **Fake Fintech** project is an initiative aimed at understanding how fintech applications work, with a focus on the underlying rules and concepts that drive them. This project explores integrating APIs connected to the economic fintech ecosystem while building a tracking system for users to keep up with their finances.

### Project Architecture

This project is structured with the following technologies:

- **Backend**: Built in **Golang** using the **Gin** framework.
- **Frontend**: The frontend is based on the **CoreUI** Angular 19 template, providing a responsive and modular UI for interacting with the backend.
- **Database**: **MongoDB** is used to store and manage financial data.
- **Cache**: Using **Redis**.

## Backend Setup

### Environment Configuration

To ensure the proper setup of your environment, the `backend/config/env/env.go` file contains the configuration for connecting to MongoDB. Make sure your environment variables are set correctly.

Example of `env.go` file:

```go
package env

import (
	"fmt"
	"os"
)

func setEnv() {
	os.Setenv("MONGO_HOST", "0.0.0.0")       
	os.Setenv("MONGO_PORT", "27017")           
	os.Setenv("MONGO_USER", "myname")             
	os.Setenv("MONGO_PASSWORD", "mypassword")      
	os.Setenv("MONGO_DATABASE", "mydatabase")
}

func GetMongoConnectionString() string {
	setEnv()
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")

	connectionStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/?directConnection=true&authSource=admin", user, password, host, port)
	return connectionStr
}

func GetDatabase() string {
	setEnv()
	return os.Getenv("MONGO_DATABASE")
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
   - Ensure that the `MONGO_*` environment variables are correctly set (as shown in the `env.go` file).
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
- **Banco de Dados**: **MongoDB** é utilizado para armazenar e gerenciar os dados financeiros.
- **Cache**: Usando **Redis**

## Configuração do Backend

### Configuração do Ambiente

Para garantir a configuração correta do seu ambiente, o arquivo `backend/config/env/env.go` contém a configuração para a conexão com o MongoDB. Certifique-se de que as variáveis de ambiente estejam corretamente configuradas.

Exemplo do arquivo `env.go`:

```go
package env

import (
	"fmt"
	"os"
)

func setEnv() {
	os.Setenv("MONGO_HOST", "0.0.0.0")       
	os.Setenv("MONGO_PORT", "27017")          
	os.Setenv("MONGO_USER", "meuusuario")     
	os.Setenv("MONGO_PASSWORD", "minhasenha") 
	os.Setenv("MONGO_DATABASE", "meubanco")   
}

func GetMongoConnectionString() string {
	setEnv()
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	user := os.Getenv("MONGO_USER")
	password := os.Getenv("MONGO_PASSWORD")

	connectionStr := fmt.Sprintf("mongodb://%s:%s@%s:%s/?directConnection=true&authSource=admin", user, password, host, port)
	return connectionStr
}

func GetDatabase() string {
	setEnv()
	return os.Getenv("MONGO_DATABASE")
}

func SetSalt() []byte {
	your_salt := "MinhaSalt"
	return []byte(your_salt)
}
