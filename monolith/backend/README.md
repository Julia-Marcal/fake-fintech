# Laravel-reusable-api

## Overview
Laravel-reusable-api is a highly performant, reusable, and developer-friendly API built with Laravel. It offers a solid set of features with a focus on clean code, scalability, and security—ideal for building modern, maintainable web applications with ease.

### Features
- **Cache Support**: Utilizes Redis for caching.
- **Database**: Mysql (configurable via Laravel's database config).
- **Rate Limiting**: Built-in Laravel rate limiter to prevent spam and service crashes.
- **Docker Support**: Includes Dockerfile and docker-compose for containerization.
- **Authentication**: JWT/Bearer token and hashed passwords for security.

## Getting Started

### Prerequisites
- PHP (8.1+ recommended)
- Composer
- Docker
- Mysql
- Redis

### Running with Docker
Execute the following commands to build and run the Docker containers:

```sh
docker-compose up --build
```

This will start the Laravel application, Mysql, and Redis containers.

## Configuration

### Environment Variables
Copy `.env.example` to `.env` and update the following variables as needed:

```
DB_CONNECTION=pgsql
DB_HOST=your-db-host
DB_PORT=5432
DB_DATABASE=your-db-name
DB_USERNAME=your-db-user
DB_PASSWORD=your-db-password

REDIS_HOST=your-redis-host
REDIS_PORT=6379
```

### JWT Configuration
Set your JWT secret in `.env`:

```
JWT_SECRET=your_jwt_secret
```

## Performance
To ensure the highest level of performance, the API has:
- Cache optimization through Redis.
- Database indexes and optimized queries in Mysql.
- Clean and efficient codebase leveraging Laravel best practices.

## Security
- Passwords are hashed for secure storage.
- JWT/Bearer token is used for system authorization.

## API Requests

Examples for making API requests can be found in the `routes/` and `app/Http/Controllers/` folders. These should guide you on how to properly make requests to the API.

Remember to use the appropriate bearer token when making API calls.
