# Products API

O projeto consiste em uma API para o gerenciamento de produtos.

### Como instalar

É possível inicializar a aplicação utilizando-se via **make** ou **docker-compose**

```shell
make up
```

OU

```shell
docker-compose up -d
```

### Como testar

#### 📦 Cadastro de produtos

```shell
curl --request POST \
  --url http://localhost:8000/products \
  --header 'Accept: application/json' \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "Product 3",
	"description": "Product 3",
	"price": 44.45
}'
```

#### 📦 Listagem de produtos

```shell
curl --request GET \
  --url http://localhost:8000/products \
  --header 'Accept: application/json'
```