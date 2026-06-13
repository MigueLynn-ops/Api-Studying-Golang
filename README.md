Esta aplicação foi atualizada para fazer a integração com o Banco de Dados - Rota de requisições consumindo API do Google Books, segue abaixo algumas instruções:

- Para rodar a aplicação, é necessário criar um arquivo .env dentro da pasta config seguindo o modelo .env-example, onde você pode configurar:
1. Qual API key será utilizada para fazer as requisições (Atualmente as API keys são obrigatórias para fazer as requisições com o Google Books)
2. O endereço da porta principal utilizada para rodar a aplicação (Porta 8080 ou outra)
3. O endereço do banco de dados que será utilizado (juntamente com usuário, root ou admin, e senha)
OBS: Por enquanto a secretkey não será necessária alterar

- As requisições de livros serão feitas no Postman utilizando o método GET no seguinte endereço: http://localhost:8080/books/search
Modelo a ser digitado no body:
intitle:nome_do_livro_aqui (Retorna resultados em que o texto após essa palavra-chave é encontrado apenas no título do livro.)
inauthor:nome_do_autor_aqui (Retorna resultados em que o texto após essa palavra-chave é encontrado apenas no nome do autor.)

- Para testar as requisições que simulam a criação de usuário no Postman, utilize o método POST no seguinte endereço: http://localhost:8080/users
Modelo a ser digitado no body:
{
    "username": "Seu Nome",
    "email": "seuemail@email.com",
    "password": "123456",
    "cpf": "xxx.xxx.xxx-xx"
}