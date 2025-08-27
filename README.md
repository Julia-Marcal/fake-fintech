# Fake Fintech

## Overview

**Fake Fintech** is a learning project designed to explore the architecture and rules behind fintech applications. It demonstrates how to integrate third-party APIs, manage asynchronous tasks, and build a robust financial tracking system using modern technologies.

## Architecture

The project is split into two main parts:

### 1. Monolithic Application

- **Backend:** PHP (Laravel)
- **Frontend:** Angular 19 (CoreUI template)
- **Database:** PostgreSQL
- **Cache:** Redis

### 2. Microservices

- **Messaging:** RabbitMQ
- **Asset Watcher Service:** Golang (fetches assets prices from multiple API)

#### Service Flow

1. **User Action:** Triggers a task (e.g., fetch crypto price).
2. **Task Queuing:** Task is sent to RabbitMQ (`assets-tasks` queue).
3. **Asset Watcher:** Golang service consumes messages, fetches prices from CoinCap processes results and return the result.

## Goals right now:
- **Dashboard**: Create a clean dashboard so user can easily access their data
- **User Financial data**: how much they invest, and use on a daily basis

## Future Development
In the future, this project will expand its functionality with advanced features like:

- **Transactions**: Manage and calculate financial transactions.
- **Interest Rates**: Retrieve and apply interest calculations.
- **User Accounts**: Create and manage user data.
- **Loan Calculations**: Implement loan calculation functionalities.
- **Transaction Analytics**: Provide detailed insights on transactions.
- **Account and Portfolio Management**: Manage user accounts and their portfolios.

---

### Fake Fintech - PT-br

[README em Português](./docs/pt-br/Readme.md)
