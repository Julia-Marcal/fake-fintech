# Fake Fintech

## Introduction

The **Fake Fintech** project is an initiative aimed at understanding how fintech applications work, with a focus on the underlying rules and concepts that drive them. This project explores integrating APIs connected to the economic fintech ecosystem while building a tracking system for users to keep up with their finances.

### Project Architecture

This project is structured with the following technologies:

- **Backend**: Built in **PHP** using the **Laravel** framework.
- **Frontend**: The frontend is based on the **CoreUI** Angular 19 template, providing a responsive and modular UI for interacting with the backend.
- **Database**: **PostgreSQL** is used to store and manage financial data.
- **Cache**: Using **Redis**.

## Backend Setup

### Environment Configuration

To ensure the proper setup of your environment, the `backend/env` file contains the configuration for connecting to PostgreSQL. Make sure your environment variables are set correctly.

Example of `.env` file:

```php
DB_CONNECTION=mysql
DB_HOST=127.0.0.1
DB_PORT=3333
DB_DATABASE=fake-fintech
DB_USERNAME=root
DB_PASSWORD=secret

CACHE_DRIVER=redis

REDIS_CLIENT=predis
REDIS_HOST=127.0.0.1
REDIS_PASSWORD=null
REDIS_PORT=6333
REDIS_DB=0
REDIS_CACHE_DB=1
REDIS_CACHE_CONNECTION=cache

SESSION_DRIVER=redis
SESSION_LIFETIME=120

APP_KEY=base64:your-app-key
JWT_SECRET=your-jwt-secret

```

### Setting Up the Backend

1. **Install dependencies**:
   - Install PHP 8.2+
   - Install composer.
   - Run the Docker Compose file to set up MySQL, Redis, and the Laravel environment.

2. **Run the backend**:
  - Copy the example environment file and configure it:

```bash
cp .env.example .env
php artisan key:generate
```

- Run the migrations:
```bash
php artisan migrate
```
- Start the Laravel development server:

``` bash
php artisan serve
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
