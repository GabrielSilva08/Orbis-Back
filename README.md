# Orbis

Aplicação mobile criada para ser apresentada durante a disciplina de Engenharia de Software 2025.1.

## Requisitos

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Passos para executar a aplicação

1. **Configurar variáveis de ambiente**  
   Copie o arquivo `.env.sample` para `.env`:
   ```bash
   cp .env.sample .env

2.**Subir o banco de dados com Docker Compose**
   No terminal, execute:
   ```bash
   docker compose up db

3.**Executar a aplicação**
   ```bash
   cd cmd
   go run main.go
