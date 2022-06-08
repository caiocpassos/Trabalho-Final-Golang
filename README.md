# Estante virtual em Go

## Introdução

- Esse é um trabalho para a disciplina de "Programação Avançada para Web", ministrada pelo professor Ricardo Mendes em nome da Universidade Vila Velha (UVV)


- **Componentes do grupo**: Aline Amy Kato, Caio Costa Passos, Diana Poltronieri Rosalem


## Conceito Geral do Projeto

- O projeto consiste em uma  estante virtual, que funciona como uma lista.


- Nela, os usuários poderão inserir livros em uma tabela, e manipulá-los como quiser.


- O projeto conta com um CRUD básico, feito utilizando os princípios de Go para Web

## Rotas Criadas

### Listagem de livros - ("/")

- A rota de listagem de livros é a própria home.


- Dentro dela, haverá uma tabela que irá listar todos os livros já inseridos no banco de dados.


- Ao lado da tabela, haverá dois botões, que irão redirecionar o usuário para a rota de alteração dos dados do livro e melhor visualização deles.


- A cada reinicialização da página, ela irá rodar uma query para atualizar a listagem

### Cadastro de livros - ("/new")

- No header da página home, haverá um botão "Novo cadastro".


- Ao clicar nesse botão, o usuário será redirecionado para a tela de cadastro.


- Nessa tela, a página pedirá os dados de um novo livro, para que ele possa ser inserido no banco de dados.


- Ao clicar no botão "Salvar livro", uma query do banco de dados irá rodar, e o livro será inserido no banco de dados.


### Atualização de livros - ("/edit")

- Na última coluna da lista de livros, é possível encontrar um botão "Editar".


- Ao clicar nesse botão, o usuário é levado para uma página com todos os dados do livro alocados em campos de texto.


- O usuário poderá alterar os campos de texto para que o livro possa ser atualizado.


- Ao clicar em "Salvar livro", uma query do banco de dados irá rodar, e o livro será alterado.


### Visualização de dados de livros - ("/show")

- Na última coluna da lista de livros, é possível encontrar um botão "Visualizar".


- Ao clicar nesse botão, o usuário é levado para uma página com todos os dados do livro disponíveis para visualização.


- Dentro da página de visualização, é possível encontrar um caminho alternativo para as funcionalidades de "Editar" e "Deletar"

### Deleção de livros - ("Sem rota")

- Ao entrar no modo de visualização de um livro, é possível encontrar um botão vermelho escrito "Deletar".


- Ao clicar nesse botão, uma query do banco de dados será acionada, e o livro selecionado será excluido do banco de dados.


## Banco de dados

### MySQL

- Para esse trabalho, optamos por utilizar o banco MySQL.


- Apesar de utilizarmos esse banco, é possível utilizar sem problemas um banco de dados não relacional, visto que existe apenas uma tabela no banco

### Estrutura do banco

- Criamos apenas uma tabela que recebeu o nome de "livros"


- Logicamente, essa tabela armazenará todos os livros existentes no banco de dados.


- Cada registro de livro contará com os seguintes atributos
    - ID
    - Titulo
    - Nome do Autor
    - Número de Páginas
    - Gênero


### Criação da tabela

- Para criar a tabela, será apenas necessário rodar os seguintes comandos.

```
CREATE DATABASE estanteLivros;

CREATE TABLE livros(
    id             int(6)   unsigned   auto_increment   primary key,
    nome           varchar(60) not null,
    autor          varchar(30) not null,
    num_paginas    int         not null,
    genero         varchar(30) not null
) charset = latin1;
```

## Stack

#### Linguagens (Programação, Marcação, Estilização e Modelagem)

- HTML
- CSS
- GOLANG
- MYSQL

#### Versionamento

- GIT
- GITHUB

#### IDEs

- GOLAND
- DATAGRIP

#### Outros

- XAMPP