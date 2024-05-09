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

-  DB_HOST=host_do_seu_banco_de_dados
-  DB_PORT=porta_do_banco_de_dados
-  DB_USER=usuario_do_banco_de_dados
-  DB_PASSWORD=senha_do_usuario_do_banco_de_dados
-  DB_NAME=nome_do_banco_de_dados


Certifique-se de substituir os valores acima pelos detalhes do seu banco de dados PostgreSQL.

## Executando a Aplicação

Para executar a aplicação, siga estas etapas:

1. Instale as dependências do projeto:

  ## go mod tidy

3. Compile o código:

 ## go build


4. Execute a aplicação:

 ## ./nome_do_executável


A aplicação estará disponível em `http://localhost:8000`.

## Contribuição

Se você encontrar problemas ou tiver sugestões para melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).


## Capítulo 1

1. Criamos o nosso projeto no workspace correto, dentro do GOPATH (dentro da pasta src, github.com, seguido do nome de usuário do Github);

2. Aprendemos como subir um servidor através da função http.ListenAndServe(), exibindo uma tabela com nossos produtos;

3. Em seguida criamos uma struct de Produto, onde instanciamos alguns produtos e exibimos de forma dinâmica em nossa index.html.


## Capítulo 2 

1. Instalamos o Postgres para armazenar nossos produtos de forma segura;

2. Criamos uma função chamada conectaComBancoDeDados() para abrir a conexão com o banco de dados;

3. Alteramos nosso código para exibir os produtos que estão cadastrados lá no banco de dados.


## Capítulo 3

1. Modularizamos o código para tornar a manutenção e/ou atualização mais clara, criando as pastas controllers, db, models, routes e templates;

2. Criamos uma página para criar novos produtos e uma rota capaz de atender essa requisição http.HandleFunc("/new", controllers.New);

3. Buscamos os dados da página new com o código r.FormValue() para cada input (nome, descrição, preço e quantidade) no controller de produtos;

4. Salvamos o produto através do modelo de produto criando a função CriaNovoProduto().


## Capítulo 4

1. Criamos um botão na linha de cada produto que assim que clicado, deletava o produto do banco de dados;

2. Para melhorar a remoção dos produtos, criamos uma função em Javascript perguntando se queremos de fato, deletar o produto;

<<<<<<< HEAD
3. Removemos o código HTML duplicado da [index] e do arquivo new, criando as seguintes partials: [_head.html] e [_menu.html].
=======
3. Removemos o código HTML duplicado da [index] e do arquivo new, criando as seguintes partials: [_head.html] e [_menu.html].


## Capítulo 5

1. Criamos um botão na linha de cada produto que nos direciona para a tela de edição;

2. Após criar a tela de edição, preenchemos o formulário com as informações do produto exibindo os dados já cadastrados;

3. Atualizamos o produto, buscando os dados alterados na tela e executando o [update] (atualização) no banco de dados.
>>>>>>> ffa32ee2ce53d0fc2f26cc244496a714d6257920
