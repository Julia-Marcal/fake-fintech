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

### Setting Up SonarQube

1. **Run SonarQube On Docker**  
   Follow the official SonarQube guide to set it up on the code locally.

2. **Install Required Extensions**  
   - Navigate to **Administration > Marketplace** within the SonarQube dashboard.  
   - Search for and install the **Dependency-Check** extension.

3. **Access SonarQube**  
   - Once running, the application will be available at [`http://localhost:9000`](http://localhost:9000).

## Future Development

In the future, this project will integrate with third-party fintech APIs to retrieve real-time financial data and expand its functionality with advanced features like:

- **Transactions**: Manage and calculate financial transactions.
- **Interest Rates**: Retrieve and apply interest calculations.
- **User Accounts**: Create and manage user data.
- **Loan Calculations**: Implement loan calculation functionalities.
- **Transaction Analytics**: Provide detailed insights on transactions.
- **Account and Portfolio Management**: Manage user accounts and their portfolios.

---

### Fake Fintech - PT-br

[README em Português](./docs/pt-br/Readme.md)