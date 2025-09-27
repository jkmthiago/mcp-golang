# Example of a MCP Server

## Visao Geral

Este repositorio fornece um servidor MCP escrito em Go. O servidor utiliza a biblioteca [`github.com/mark3labs/mcp-go`](https://github.com/mark3labs/mcp-go) para registrar ferramentas MCP e expor uma API HTTP em `/api/v1/mcp`. O objetivo atual e disponibilizar um exemplo simples implementando uma ferramenta de calculadora que executa operacoes matematicas basicas.

## Estrutura do Projeto

- `docker-compose.yaml`: orquestra containers para desenvolvimento (servidor Go).
- `src/main.go`: inicializa o servidor MCP HTTP na porta 8080.
- `src/mcp/mcp.go`: configura o servidor MCP e registra as ferramentas disponiveis.
- `src/tools/tools.go`: implementa os handlers das ferramentas (hoje apenas a calculadora).

## Requisitos

- Docker e Docker Compose.
- Go 1.25+ (opcional, caso queira rodar fora do Docker).

## Como Executar

### Usando Docker Compose

1. Na raiz do projeto, execute `docker compose up -d` para iniciar o container `mcp`.
2. O servidor MCP ficara acessivel em `http://localhost:8080/api/v1/mcp`.
3. Para acompanhar logs, use `docker compose logs -f mcp`.
4. Para encerrar, rode `docker compose down`.

> O volume `./src:/go/src/mcp` permite editar o codigo localmente e refletir as mudancas no container imediatamente.

### Execucao Local sem Docker

1. Entre na pasta `src`.
2. Instale as dependencias com `go mod tidy` (somente na primeira vez).
3. Rode o servidor com `go run .`.
4. A API MCP continua disponivel na porta 8080.

## Ferramentas MCP Disponiveis

- `calculator`
  - `operation`: `add`, `subtract`, `multiply`, `divide`.
  - `x`: numero real (primeiro operando).
  - `y`: numero real (segundo operando).
  - Retorna o resultado formatado com 7 casas decimais ou uma mensagem de erro (`invalid operation` ou `cannot divide by zero`).

## Endpoints de Referencia

- `POST /api/v1/mcp` (protocolo MCP streamable HTTP).
- Exemplo simples usando `curl`:
  ```bash
  curl -sS -X POST \
    -H 'Content-Type: application/json' \
    -d '{"method":"tools/call","params":{"tool":"calculator","arguments":{"operation":"add","x":2,"y":3}}}' \
    http://localhost:8080/api/v1/mcp
  ```
