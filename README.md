# Fake Fintech - Monitoramento Financeiro

O **Fake Fintech OSS** é um projeto educacional open source voltado para o aprendizado de **arquitetura observável, mensageria e automação de eventos financeiros em tempo real**, utilizando uma stack moderna com **Go, RabbitMQ, Redis, PostgreSQL, Prometheus e Grafana**.

A ideia é simular o ecossistema técnico de uma fintech moderna — com foco em **desempenho, escalabilidade e visibilidade total do sistema**, aplicando conceitos de **observabilidade, mensageria assíncrona e automação de alertas**.

---

## 🚀 Objetivo do Projeto

Demonstrar como construir uma arquitetura **100% open source**, que:
- coleta e processa dados financeiros (reais ou simulados),
- gera métricas técnicas e de negócio,
- detecta eventos significativos (como variações bruscas de ativos),
- e reage automaticamente, registrando logs e disparando alertas.

---

## 🧩 Arquitetura Geral

O sistema é composto por **microsserviços Go** interligados por **RabbitMQ** e monitorados com **Prometheus + Grafana + Loki**.

```
┌──────────────────┐      ┌────────────────┐      ┌──────────────────┐
│ Asset Watcher    │ ---> │ RabbitMQ Queue │ ---> │ Analyzer Service  │
│ (coleta APIs)    │      │ (mensageria)   │      │ (gera métricas)  │
└──────────────────┘      └────────────────┘      └──────────────────┘
           │                                             │
           ▼                                             ▼
     PostgreSQL + Redis                          Prometheus + Loki
                                                      │
                                                      ▼
                                                  Grafana OSS
```

---

## ⚙️ Tecnologias Utilizadas

| Categoria | Tecnologia | Descrição |
|------------|-------------|------------|
| **Linguagem** | [Go](https://go.dev) | Backend e microsserviços concorrentes |
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

4. **Observabilidade**  
   - Logs coletados via Loki e exibidos no Grafana.  
   - Métricas e alertas exibidos em dashboards interativos.  

---

## 🧠 Conceitos Práticos Envolvidos

- **Arquitetura baseada em eventos (Event-Driven Architecture)**  
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
│   └── notifier/            # envia alertas ou mensagens
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
```

Após subir o ambiente:
- **RabbitMQ:** http://localhost:15672  
- **Prometheus:** http://localhost:9090  
- **Grafana:** http://localhost:3000  
- **Loki:** http://localhost:3100  

---

## 📊 Próximos Passos

- [ ] Criar dashboards de métricas no Grafana  
- [ ] Adicionar tracing distribuído com Jaeger  
- [ ] Implementar autenticação e controle de acesso  
- [ ] Publicar imagens no Docker Hub  
- [ ] Configurar pipeline CI/CD com Drone ou GitLab CE  
