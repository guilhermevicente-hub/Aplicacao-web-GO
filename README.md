# Aplicação Web em Go com Conexão ao PostgreSQL

Este repositório contém uma aplicação web desenvolvida em Go que se conecta a um banco de dados PostgreSQL. Este projeto faz parte da formação em Linguagem Go oferecida pela Alura.

## Pré-requisitos

Antes de executar esta aplicação, certifique-se de ter o seguinte instalado em sua máquina:

- [Go](https://golang.org/) (versão 1.16 ou superior)
- [PostgreSQL](https://www.postgresql.org/)

## Configuração do Banco de Dados

Antes de iniciar a aplicação, é necessário configurar o banco de dados PostgreSQL. Você pode usar as seguintes etapas como referência:

1. Instale o PostgreSQL em sua máquina, se ainda não estiver instalado.
2. Crie um banco de dados PostgreSQL.
3. Crie uma tabela no banco de dados para armazenar os dados da aplicação.

## Configuração da Aplicação

Antes de executar a aplicação, você precisará configurar as variáveis de ambiente necessárias. Para isso, crie um arquivo `.env` na raiz do projeto e defina as seguintes variáveis:

DB_HOST=host_do_seu_banco_de_dados
DB_PORT=porta_do_banco_de_dados
DB_USER=usuario_do_banco_de_dados
DB_PASSWORD=senha_do_usuario_do_banco_de_dados
DB_NAME=nome_do_banco_de_dados


Certifique-se de substituir os valores acima pelos detalhes do seu banco de dados PostgreSQL.

## Executando a Aplicação

Para executar a aplicação, siga estas etapas:

1. Instale as dependências do projeto:

  go mod tidy

3. Compile o código:

  go build


4. Execute a aplicação:

  ./nome_do_executável


A aplicação estará disponível em `http://localhost:8000`.

## Contribuição

Se você encontrar problemas ou tiver sugestões para melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).




