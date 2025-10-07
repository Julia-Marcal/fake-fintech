# Fake Fintech OSS — Monitoramento Financeiro

O **Fake Fintech OSS** é um projeto educacional open source voltado para o aprendizado de **arquitetura observável, mensageria e automação de eventos financeiros em tempo real**, utilizando uma stack moderna com **Go, GraphQL, RabbitMQ, Redis, PostgreSQL, Prometheus e Grafana**.

A ideia é simular o ecossistema técnico de uma fintech moderna — com foco em **desempenho, escalabilidade e visibilidade total do sistema**, aplicando conceitos de **observabilidade, mensageria assíncrona e automação de alertas**.

---

## 🚀 Objetivo do Projeto

Demonstrar como construir uma arquitetura **100% open source**, que:
- coleta e processa dados financeiros (reais ou simulados),
- disponibiliza informações via **GraphQL API** unificada,
- gera métricas técnicas e de negócio,
- detecta eventos significativos (como variações bruscas de ativos),
- e reage automaticamente, registrando logs e disparando alertas.

---

## 🧩 Arquitetura Geral

O sistema é composto por **microsserviços Go** interligados por **RabbitMQ**, **Redis** e monitorados com **Prometheus + Grafana + Loki**.  
A camada **GraphQL Gateway** atua como ponto central de acesso a dados financeiros, métricas e logs.

```
┌──────────────────────────────┐
│       GraphQL Gateway        │
│   (Go + gqlgen API layer)    │
└───────────────┬──────────────┘
                │
┌───────────────┼───────────────────────────────┐
│               │                               │
▼               ▼                               ▼
Asset Watcher   Analyzer Service                Notifier
(coleta APIs)   (gera métricas)                 (envia alertas)
│               │                               │
▼               ▼                               ▼
PostgreSQL + Redis       Prometheus + Loki       RabbitMQ
                                 │
                                 ▼
                           Grafana OSS
```

---

## ⚙️ Tecnologias Utilizadas

| Categoria | Tecnologia | Descrição |
|------------|-------------|------------|
| **API Gateway** | [gqlgen](https://github.com/99designs/gqlgen) | Exposição de dados via GraphQL |
| **Linguagem** | [Go](https://go.dev) | Microsserviços e API GraphQL |
| **Mensageria** | [RabbitMQ](https://www.rabbitmq.com) | Comunicação assíncrona entre serviços |
| **Cache e Fila leve** | [Redis OSS](https://redis.io) | Cache de resultados e mensagens temporárias |
| **Banco de Dados** | [PostgreSQL](https://www.postgresql.org) | Armazenamento de dados estruturados |
| **Métricas** | [Prometheus](https://prometheus.io) | Coleta e exposição de métricas |
| **Dashboards** | [Grafana OSS](https://grafana.com/oss/) | Visualização de métricas e alertas |
| **Logs** | [Loki + Promtail](https://grafana.com/oss/loki/) | Coleta e consulta de logs distribuídos |
| **Alertas** | [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/) | Envio de alertas baseados em métricas |
| **Testes Automatizados** | [Testify](https://github.com/stretchr/testify) | Testes unitários e de integração em Go |
| **Containerização** | [Docker Compose](https://docs.docker.com/compose/) | Orquestração dos serviços localmente |

---

## 💼 Regras de Negócio Simuladas

O Fake Fintech OSS implementa **regras automatizadas baseadas em eventos financeiros** simulados.  
Cada serviço participa de um ciclo de monitoramento:

1. **Asset Watcher**  
   - Coleta cotações de ativos em APIs públicas (ex: CoinCap, AlphaVantage).  
   - Publica os resultados em uma fila RabbitMQ (`assets-tasks`).

2. **Analyzer Service**  
   - Consome mensagens do RabbitMQ.  
   - Aplica regras de negócio, como:  
     - “Alertar se o preço subir mais de 10% em 5 minutos.”  
     - “Registrar evento se o ativo cair abaixo da média histórica.”  
   - Expõe métricas via endpoint Prometheus (`/metrics`).

3. **Notifier (opcional)**  
   - Consome alertas e envia notificações (ex: email, Telegram, Slack).  

4. **GraphQL Gateway**  
   - Atua como camada única de consulta, agregando dados de múltiplas fontes:  
     - PostgreSQL (transações e portfólios)  
     - Redis (cache de ativos)  
     - Prometheus (métricas em tempo real)  
   - Permite que dashboards e clientes consultem dados com flexibilidade, sem sobrecarregar os microsserviços.

5. **Observabilidade**  
   - Logs coletados via Loki e exibidos no Grafana.  
   - Métricas e alertas exibidos em dashboards interativos.  

---

## 🧠 Conceitos Práticos Envolvidos

- **Arquitetura baseada em eventos (Event-Driven Architecture)**  
- **API unificada via GraphQL (Go + gqlgen)**  
- **Mensageria e processamento assíncrono com RabbitMQ**  
- **Observabilidade com Prometheus, Grafana e Loki**  
- **Concorrência e paralelismo com goroutines e channels**  
- **Testes automatizados e cobertura de código com Go**  
- **Infraestrutura containerizada com Docker Compose**

---

## 🐳 Estrutura do Projeto

```
fake-fintech/
│
├── services/
│   ├── asset-watcher/       # coleta dados de APIs financeiras
│   ├── analyzer/            # consome eventos e gera métricas
│   ├── notifier/            # envia alertas ou mensagens
│   └── graphql-api/         # gateway GraphQL (Go + gqlgen)
│
├── configs/
│   ├── prometheus.yml       # configuração do Prometheus
│   ├── alertmanager.yml     # regras de alerta
│   └── grafana/             # dashboards e datasources
│
├── docker-compose.yml       # orquestra todos os serviços
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Exemplo de Configuração GraphQL

Dentro de `services/graphql-api/`:

```bash
go mod init github.com/seuusuario/fake-fintech-oss/graphql-api
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

Isso gera:
- `graph/schema.graphqls` — definição dos tipos e queries/mutations.  
- `graph/resolver.go` — implementação das resolvers (consultando PostgreSQL, Redis e Prometheus).  
- `server.go` — servidor principal expondo `/graphql`.

### Exemplo de Query
```graphql
query {
  asset(id: "BTC") {
    name
    price
    variation24h
  }
}
```

---

## 🧰 Como Executar Localmente

### Pré-requisitos
- Docker e Docker Compose instalados  
- Go 1.22+ instalado  

### Passos

```bash
# Clone o repositório
git clone https://github.com/seuusuario/fake-fintech-oss.git
cd fake-fintech-oss

# Suba os containers principais
docker compose up -d

# Rode os serviços Go
cd services/asset-watcher && go run main.go
cd ../analyzer && go run main.go
cd ../graphql-api && go run server.go
```

Após subir o ambiente:
- **GraphQL Playground:** http://localhost:8080/graphql  
- **RabbitMQ:** http://localhost:15672  
- **Prometheus:** http://localhost:9090  
- **Grafana:** http://localhost:3000  
- **Loki:** http://localhost:3100  

---

## ✅ Testes

Execute todos os testes com cobertura:
```bash
go test ./... -v -cover
```

---

## 📊 Próximos Passos

- [ ] Criar dashboards de métricas no Grafana  
- [ ] Adicionar tracing distribuído com Jaeger  
- [ ] Implementar autenticação e controle de acesso no GraphQL  
- [ ] Publicar imagens no Docker Hub  
- [ ] Configurar pipeline CI/CD com Drone ou GitLab CE  

---

## 🧑‍💻 Autora

Desenvolvido por **Julia Schmerz** 💡  
> Projeto educacional open source para aprendizado em Go, GraphQL, mensageria e observabilidade de sistemas financeiros.

---

## 📄 Licença

Este projeto é licenciado sob a [MIT License](LICENSE).
