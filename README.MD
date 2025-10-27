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

```bash
.
├── db/
│   └── connection.go       # Configuração do cliente MySQL                 
├── handler/
│   └── user.go             # Handlers HTTP (controladores)
│   └── task.go
├── model/
│   └── user.go             # Define os dados e structs (User, Task)
├── repository/
│   └── user.go             # Interações com o banco de dados
│   └── task.go
├── service/
│   └── user.go             # Regras de negócio (validações e lógica)
│   └── task.go
├── main.go                 # Ponto de entrada da aplicação
├── Dockerfile
├── docker-compose.yml 
├── .env                    # Variáveis de ambiente
├── go.mod                  # Dependências do projeto
├── go.sum                  # Sums de verificação
├── Makefile                # Comando make run para rodar o projeto
└── README.md               # Documentação
```


---

## 🧠 Funcionalidades Implementadas

- **Usuário**
> Não é possivel criar mais de um usuário com o mesmo E-mail.
  - Criar usuário (`POST /user`)
  ```
    {
	    "nome": "nome1",
	    "email": "nome1@gmail.com",
	    "password": "123456"
    }

  ```

- **Tarefas**
  - Criar tarefa (`POST /user/:id/tasks`)
	```
		{
   			 "title": "Software Developer",
   			 "description": "Desenvolvendo um software de gerenciamento de tarefas", 
   			 "status": "done"
		}
	```

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
