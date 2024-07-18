Rodar o arquivo docker compose: \
`docker-compose up -d`

Acessar o container do mysql: \
`docker-compose exec mysql bash`

Dentro do container do mysql, acessar seu root: \
`mysql -uroot -p goexpert`

Dentro do root do seu mysql, criar a tabela e fazer os comandos necessários: \
`create table wsqlproducts (id varchar(255), name varchar(80), price decimal(10,2), primary key(id))`

*A criação da tabela é apenas para o exemplo writtensql, no exemplo do gorm, as tabelas são criadas automaticamente :D*


