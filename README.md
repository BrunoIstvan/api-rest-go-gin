## Executar API

Fazer download do código-fonte

    git clone https://github.com/BrunoIstvan/api-rest-go-gin.git

Acessar a pasta da aplicação

    cd api-rest-go-gin

Instalar dependências

    go install

Iniciar container do banco de dados 

    docker compose up

Iniciar a aplicação

    go run main.go

Página de listagem de todos os alunos

    http://localhost:8080/index

Página não encontrada

    http://localhost:8080/home