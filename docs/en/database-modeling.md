# Database Architecture

This document describes the current database modeling for the backend. The schema consists of five main entities: `User`, `Wallet`, `WatchWallet`, `WalletAcoes`, `WatchWalletAcoes`, and `Acoes`.

## Entity Relationship Diagram (ERD)

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

## Notes
- The `UserId` in `Wallet` and `WatchWallet` establishes a relationship between users and their respective entities.
- The `WalletAcoes` and `WatchWalletAcoes` tables create many-to-many relationships between wallets/watch wallets and stocks.
- Timestamps (`CreatedAt`, `UpdatedAt`) are included in all tables for tracking purposes.