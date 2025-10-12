## Uso do Makefile (publicar mensagens no RabbitMQ)

Este repositório contém um `Makefile` que facilita a publicação de uma mensagem `assets-tasks` no RabbitMQ, usando o CLI publisher localizado em `assets-publisher/cmd`.

Alvos (targets)
- `make publish` — executa o publisher com `go run` (não gera binário).
- `make publish-binary` — compila o binário e o executa.
- `make build-publish` — compila apenas o binário do publisher.
- `make clean` — remove o binário compilado.
- `make help` — mostra uma ajuda curta.

Variáveis
- `ACTION` — ação da tarefa (padrão: `FETCH_PRICE`).
- `MARKET` — mercado (padrão: `crypto`).
- `SYMBOL` — símbolo do ativo (padrão: `BTC`).
- `URL` — URL de conexão do RabbitMQ. Se não for fornecida, o Makefile usa a variável de ambiente `RABBITMQ_URL`.

Exemplos (PowerShell)

Defina `RABBITMQ_URL` na sessão e execute `make publish`:

```powershell
$env:RABBITMQ_URL='amqp://guest:guest@localhost:5672/'
make publish ACTION=FETCH_PRICE MARKET=crypto SYMBOL=BTC
```

Ou passe a URL diretamente ao executar `make`:

```powershell
make publish ACTION=FETCH_PRICE MARKET=crypto SYMBOL=BTC URL=amqp://guest:guest@localhost:5672/
```

Compilar e executar o binário compilado:

```powershell
make publish-binary ACTION=FETCH_PRICE MARKET=crypto SYMBOL=BTC URL=amqp://guest:guest@localhost:5672/
```

Notas e resolução de problemas
- O Makefile foi escrito para ser portátil e usa verificações internas do `make` para garantir que a variável `URL` (ou `RABBITMQ_URL`) esteja definida antes de executar o publisher. Se `URL` estiver ausente, o `make` abortará com uma mensagem de erro clara.
- Em alguns ambientes Windows, `make` pode invocar comandos via PowerShell/`cmd.exe`. Se você tiver erros relacionados a `sh` ou comandos POSIX ausentes, instale o Git for Windows e execute `make` pelo Git Bash, ou ajuste seu PATH para ter um `sh` compatível.
- Se preferir uma experiência totalmente nativa do PowerShell, eu posso adicionar um script wrapper `scripts/publish.ps1` que faz as verificações e chama o `go run` ou o binário de forma adequada ao PowerShell.

Formato da mensagem
- O publisher envia um JSON compatível com `internal/domain/entity/AssetTask`:

```json
{
  "action": "FETCH_PRICE",
  "market": "crypto",
  "symbol": "BTC"
}
```

Próximos passos (opções)
- Posso adicionar o wrapper PowerShell `scripts/publish.ps1` (recomendado para usuários Windows).
- Posso definir um `URL` padrão no Makefile (`amqp://guest:guest@localhost:5672/`) para não precisar passar a URL sempre.

Diga qual opção prefere e eu implemento.
# Using the Makefile (publish messages to RabbitMQ)

This repository includes a Makefile that simplifies publishing an `assets-tasks` message to RabbitMQ using the small publisher CLI in `cmd/publish-message`.

Targets
- `make publish` — runs the publisher using `go run` (no binary output).
- `make publish-binary` — builds the binary and runs it.
- `make build-publish` — builds the publisher binary only.
- `make clean` — removes the built binary.
- `make help` — shows help text.

Variables
- `ACTION` — task action (default: `FETCH_PRICE`).
- `MARKET` — market (default: `crypto`).
- `SYMBOL` — asset symbol (default: `BTC`).
- `URL` — RabbitMQ connection URL. If not provided, the Makefile will use the `RABBITMQ_URL` environment variable.

PowerShell examples

Set `RABBITMQ_URL` in your session and run `make publish`:

```powershell
make publish ACTION=FETCH_PRICE MARKET=crypto SYMBOL=BTC URL=amqp://guest:guest@localhost:5672/
```

Notes and troubleshooting
- On Windows, `make` typically runs shell commands using a Unix-like shell (MSYS, Git Bash) provided by Git for Windows. The Makefile uses `[` style tests for checking variables, which require a Unix-like shell. If your `make` invocation runs commands through PowerShell instead, those `[` checks will fail.
- If you run into issues like `sh: 1: [: not found`, install Git for Windows and run `make` from a shell that uses MSYS (Git Bash), or update your PATH so `sh` is available.
- If you prefer PowerShell-only invocation, I can add a small `scripts/publish.ps1` wrapper that sets defaults and calls the compiled binary or `go run` directly.

Message format
- The publisher sends a JSON object compatible with `internal/domain/entity/AssetTask`:

```json
{
  "action": "FETCH_PRICE",
  "market": "crypto",
  "symbol": "BTC"
}
```

If you want, I can also add the PowerShell wrapper or adjust the Makefile to be PowerShell-compatible. Which would you prefer?
