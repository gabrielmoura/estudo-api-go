# Estudo de Go e API

[![wakatime](https://wakatime.com/badge/user/d38eb168-6d29-49d2-bed8-f1f729c66217/project/018b3f5e-43e7-49a0-a247-0f3df625c66b.svg)](https://wakatime.com/badge/user/d38eb168-6d29-49d2-bed8-f1f729c66217/project/018b3f5e-43e7-49a0-a247-0f3df625c66b)

## Objetivo

Desenvolver uma API eficiente para o cadastro de produtos e gerenciamento de carrinho de compras, incorporando
funcionalidades de login e autenticação, é o objetivo central deste projeto. Para alcançar esse propósito, será
elaborada uma API em Go, seguindo as boas práticas de desenvolvimento e utilizando variáveis de ambiente. As informações
serão armazenadas de forma segura em um banco de dados SQLite.

A estrutura do projeto é parcialmente baseada nas diretrizes
do [Golang Standards](https://github.com/golang-standards/project-layout), incorporando entidades, interfaces e
variáveis de ambiente sem depender de um ORM.

Além disso, o projeto incluirá tratamento de dados, uma camada de segurança adicional com a transformação da senha
utilizando o algoritmo Argon, e uma robusta suite de testes para assegurar a qualidade do código desenvolvido.

## Tecnologias Utilizadas

- Goland
- Go Gin
- viper
- Swagger
- Sqlite

### Como rodar o projeto ✅
```bash
cd cmd/server
go build -tags=jsoniter .
./server
```

###  Variável de Ambiente

```env
DB_DRIVER=sqlite
DB_NAME="./sqlite.db"

JWT_SECRET=secretAAAASSSSSSS
JWT_EXPIRESIN=3600

WEB_SERVER_PORT=8001

GIN_MODE=release
```

### Documentação

- [Swagger](http://localhost:8001/swagger/index.html)