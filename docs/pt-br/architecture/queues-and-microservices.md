# 🧩 Comunicação entre Serviços com Filas de Mensagens

Este documento explica a arquitetura de serviços e por que usamos **RabbitMQ** para conectar microserviços Laravel e Go de forma assíncrona.

---

## 🧱 Visão Geral da Arquitetura

```
[Usuário / Frontend]
        ↓
 ┌──────────────────────┐
 │   Laravel (PHP) API  │  ← Serviço Principal
 │ - Autenticação, BD,  │
 │   Regras de Negócio  │
 │ - UI / Painel Admin  │
 └────────┬─────────────┘
          ↓
 ┌─────────────────────┐
 │ Microserviço: Go    │  ← APIs externas, scraping, etc.
 └─────────────────────┘
          ↓
 ┌─────────────────────┐
 │ Microserviço: Go    │  ← Processos em lote, analytics
 └─────────────────────┘
```

---

## 🔁 Comunicação via Filas de Mensagens

Este projeto utiliza **comunicação assíncrona** por meio de filas de mensagens para conectar os serviços Laravel e Go — melhorando desempenho, tolerância a falhas e escalabilidade.

- 📨 Laravel **envia tarefas** para a fila  
- ⚙️ Microserviços Go **consomem e processam** as tarefas  
- 💾 Resultados são **armazenados no BD** ou opcionalmente enviados de volta por uma fila de resposta

### Broker suportado:

- 🐇 [RabbitMQ](https://www.rabbitmq.com/)

---

## ✅ Por que PHP + Go + Filas?

| Benefício                | Por que é importante                                         |
|--------------------------|-------------------------------------------------------------|
| **Agilidade no Dev**     | Laravel é rápido para construir, ideal para app principal e UI admin |
| **Alta Performance**     | Go lida com tarefas intensivas, paralelas ou de baixa latência |
| **Resiliência**          | Filas ajudam a recuperar de falhas temporárias automaticamente |
| **Escalabilidade**       | Cada serviço escala de forma independente (ex: mais workers Go) |

> 🏦 Ideal para **sistemas fintech** onde dados em tempo real, APIs de terceiros e lógica orientada a eventos são críticos.