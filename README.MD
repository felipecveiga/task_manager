# 🧩 Task Manager API

API de gerenciamento de tarefas desenvolvida em **Golang** com **GORM** e **MySQL**, containerizada com **Docker**.

> ⚠️ Este projeto implementa apenas o **backend**, com foco no CRUD de usuários e tarefas.  
> A autenticação JWT, frontend e testes unitários não foram implementados.

---

## 🚀 Tecnologias Utilizadas

- **Golang**
- **Echo Framework**
- **GORM** (ORM para Go)
- **MySQL**
- **Docker** e **Docker Compose**

---

## 📁 Estrutura do Projeto

.
├── db/ # Conexão com o banco de dados
├── handler/ # Handlers HTTP (controladores)
├── repository/ # Camada de acesso ao banco
├── service/ # Regras de negócio
├── model/ # Models GORM
├── .env.docker # Variáveis de ambiente para Docker
├── Dockerfile
├── docker-compose.yml
└── main.go


---

## 🧠 Funcionalidades Implementadas

- **Usuário**
  - Criar usuário (`POST /user`)

- **Tarefas**
  - Criar tarefa (`POST /user/:id/tasks`)
  - Buscar tarefas de um usuário (`GET /user/:id/tasks`)
  - Atualizar tarefa (`PATCH /user/:id/tasks/:task_id`)
  - Deletar tarefa (deleção lógica - `DELETE /user/:id/tasks/:task_id`)

> Cada tarefa é associada a um usuário por meio do campo `user_id`.  
> A deleção é lógica, ou seja, o registro é apenas marcado como excluído, e não removido fisicamente do banco.

---

## ⚙️ Configuração de Ambiente

Antes de rodar o projeto, verifique se o **Docker Desktop** está aberto.

O projeto utiliza o arquivo `.env` para configurar o banco de dados:


## 🐳 Rodando o Projeto com Docker

1 - Clone o repositório:

``` git clone https://github.com/felipecveiga/task_manager.git ```

2 - Construa e suba os containers:

```docker-compose up --build```

3 - Aguarde o MySQL inicializar.

O container da aplicação (app) só inicia após o banco estar pronto.

``` Conexao com o banco de dados realizada com sucesso```

4 - Acesse a API:

Base URL: ```http://localhost:8080 ```

## Observações

O projeto utiliza GORM AutoMigrate, então não é necessário ter arquivos de migração SQL obrigatoriamente.

Para deleção de tarefas, é feita uma deleção lógica utilizando DeletedAt.