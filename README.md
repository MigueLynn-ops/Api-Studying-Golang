Esta aplicação foi atualizada para integrar o CRUD e o Validador de CPF - Rota consumindo API do Google Books, segue abaixo algumas instruções:

- A aplicação estará rodando na porta 8080, mas caso essa porta já esteja sendo utilizada na sua máquina, você poderá (no arquivo main.go) trocar para outra porta que esteja disponível.

- As requisições serão feitas no Postman utilizando o método GET no seguinte endereço: http://localhost:8080/books/search
<br/><br/>IMPORTANTE
- Atualmente será necessário utilizar sua API key no código para que as requisições sejam bem sucedidas:
1. Crie um arquivo .env na raiz do projeto
2. Adicione sua API key do Google Books seguindo o exemplo --> GOOGLE_API_KEY=your_api_key_here
3. Rode a aplicação normalmente e teste as requisições no Postman

- Para testar as requisições que simulam a criação de usuário no Postman, utilize o método POST no seguinte endereço: http://localhost:8080/users
Modelo a ser digitado no body:
{
    "username": "Seu Nome",
    "email": "seuemail@email.com",
    "password": "123456",
    "cpf": "xxx.xxx.xxx-xx"
}