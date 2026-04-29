Esta aplicação é referente a Entrega 2 - Rota consumindo API do Google Books, segue abaixo algumas instruções:

- A aplicação estará rodando na porta 8080, mas caso essa porta já esteja sendo utilizada na sua máquina, você poderá (no arquivo main.go) trocar para outra porta que esteja disponível.
- As requisições serão feitas no Postman utilizando o seguinte endereço: http://localhost:8080/books/search
<br/><br/>IMPORTANTE
- Será necessário utilizar sua API key no código para que as requisições sejam bem sucedidas:
1. Crie um arquivo .env na raiz do projeto
2. Adicione sua API key do Google Books seguindo o exemplo --> GOOGLE_API_KEY=your_api_key_here
3. Rode a aplicação normalmente e teste as requisições no Postman
