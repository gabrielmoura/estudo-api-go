# Estudo de Go e API

[![wakatime](https://wakatime.com/badge/user/d38eb168-6d29-49d2-bed8-f1f729c66217/project/018b3f5e-43e7-49a0-a247-0f3df625c66b.svg)](https://wakatime.com/badge/user/d38eb168-6d29-49d2-bed8-f1f729c66217/project/018b3f5e-43e7-49a0-a247-0f3df625c66b)

- Dificuldades superadas
- Criar Docker

Documentação:

- Funcionalidades
- Objetivos

## Objetivo

Este projeto consiste na elaboração de uma API em Go fazendo uso de boas praticas e variaveis de ambiente, que salva as informações em um banco sqlite.

Com a infraestrutura parcialmente baseada em [Golang Standards](https://github.com/golang-standards/project-layout) com entidades, interfaces e variaveis de ambiente sem o uso de ORM.

Há tratamento de dados, transformação da senha com argon e testes para garantir a qualidade.

## Ferramenta usadas

- Goland

## Artigos e Cursos

- [Formação Go - Alura](https://cursos.alura.com.br/formacao-go)
- [Go (Golang): Explorando a Linguagem do Google - COD3R](https://www.udemy.com/course/curso-go)
- [how-to-hash-and-verify-passwords-with-argon2-in-go](https://www.alexedwards.net/blog/how-to-hash-and-verify-passwords-with-argon2-in-go)

### Rotas

- Usuários

```http request
GET /user
```

- Produtos

```http request
GET /product
```
