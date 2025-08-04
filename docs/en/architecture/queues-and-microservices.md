# 🧩 Service Communication with Message Queues

This document explains the service architecture and why we use **RabbitMQ** to connect Laravel and Go microservices asynchronously.

---

## 🧱 Architecture Overview

```
[User / Frontend]
        ↓
 ┌──────────────────────┐
 │   Laravel (PHP) API  │  ← Main Service
 │ - Auth, DB, Business │
 │ - UI / Admin Panel   │
 └────────┬─────────────┘
          ↓
 ┌─────────────────────┐
 │ Microservice: Go    │  ← External APIs, scraping, etc.
 └─────────────────────┘
          ↓
 ┌─────────────────────┐
 │ Microservice: Go    │  ← Batch jobs, analytics
 └─────────────────────┘
```

---

## 🔁 Message Queue Communication

This project uses **asynchronous communication** via message queues to connect Laravel and Go services — improving performance, fault tolerance, and scalability.

- 📨 Laravel **dispatches jobs** to the queue  
- ⚙️ Go microservices **consume and process** them  
- 💾 Results are **stored in the DB** or optionally sent back via a response queue

### Supported broker:

- 🐇 [RabbitMQ](https://www.rabbitmq.com/)

---

## ✅ Why PHP + Go + Queues?

| Benefit                   | Why It Matters                                              |
|---------------------------|-------------------------------------------------------------|
| **Developer Speed**       | Laravel is fast to build, ideal for core app + admin UI     |
| **High Performance**      | Go handles intensive, parallel, or low-latency tasks        |
| **Resilience**            | Queues help recover from temporary failures automatically   |
| **Scalability**           | Each service scales independently (e.g., more Go workers)   |

> 🏦 Ideal for **fintech systems** where real-time data, third-party APIs, and event-driven logic are critical.