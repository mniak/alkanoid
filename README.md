Alkanoid
================


## Compilar/Executar

### Usando Docker

É necessário ter o docker instalado e rodando em sua máquina.

#### Gerando a imagem Docker
```
docker build . -f cmd/http/Dockerfile -t alkanoid:http
```

#### Rodando um container Docker
```
docker run --rm -p 8080:8080 alkanoid:http
```

### Usando o Go

Você deve ter o Go instalado no seu computador. 
Para mais instruções, acesse https://golang.org/doc/install.

Em seguida, execute o seguinte comando:
```
go run ./cmd/http
```

Pronto, a aplicação já está rodando e poderá receber chamadas em http://localhost:8080!
