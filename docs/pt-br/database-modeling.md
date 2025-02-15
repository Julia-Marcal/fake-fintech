# Arquitetura do Banco de Dados  


Este documento descreve a modelagem atual do banco de dados para o backend. O esquema consiste em seis entidades principais: `User`, `Wallet`, `WatchWallet`, `WalletAcoes`, `WatchWalletAcoes`, and `Acoes`.

## Diagrama de Relacionamento de Entidades (ERD)  

```
+-----------------+          +-----------------+          +-----------------+
|     User        |          |     Wallet      |          |  WatchWallet    |
+-----------------+          +-----------------+          +-----------------+
| Id (PK)         |<-------->| UserId (FK)     |          | UserId (FK)     |
| Name            |          | Id (PK)         |          | Id (PK)         |
| LastName        |          | CreatedAt       |          | CreatedAt       |
| Age             |          | UpdatedAt       |          | UpdatedAt       |
| Email (Unique)  |          +-----------------+          +-----------------+
| Password        |               |      |                     |      |
| CreatedAt       |               |      |                     |      |
| UpdatedAt       |               |      |                     |      |
+-----------------+               |      |                     |      |
                                  |      |                     |      |
                                  v      v                     v      v
                          +-----------------+          +-----------------+
                          |  WalletAcoes    |          | WatchWalletAcoes|
                          +-----------------+          +-----------------+
                          | WalletId (FK)   |          | WatchWalletId(FK)|
                          | AcoesId (FK)    |          | AcoesId (FK)     |
                          | Quantity        |          | Quantity         |
                          +-----------------+          +-----------------+
                                  |      |                     |      |
                                  |      |                     |      |
                                  v      v                     v      v
                          +-----------------+
                          |     Acoes       |
                          +-----------------+
                          | Id (PK)         |
                          | Name            |
                          | Type            |
                          | Price           |
                          | CreatedAt       |
                          | UpdatedAt       |
                          +-----------------+
```


## Notas  
- O `UserId` em `Wallet` e `WatchWallet` estabelece um relacionamento entre os usuários e suas respectivas entidades.  
- As tabelas `WalletAcoes` e `WatchWalletAcoes` criam relações muitos-para-muitos entre carteiras/watch wallets e ações.  
- Os timestamps (`CreatedAt`, `UpdatedAt`) estão incluídos em todas as tabelas para fins de rastreamento.  