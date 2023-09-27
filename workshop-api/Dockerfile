# O comando FROM serve para definir a partir de qual imagem base será criado a 
# imagem customizada
FROM golang:1.20.8-alpine AS builder

# O comando WORKDIR serve para criar um diretório e caso não exista e também define como o diretório
# principal após acessar o container
WORKDIR /app

# O comando COPY server parar copiar os arquivos, no exemplo esta copiando todos os diretorio da mesma raiz
# que o dockerfile e copiando para dentro do container
COPY . .

RUN go mod download \
    ; CGO_ENABLED=0 go build -o bin/api cmd/api/main.go

# O comando FROM serve para definir a partir de qual imagem base estamos criando a 
# nossa imagem customizada
# neste ponto estamos usando o conceito de multi-stage no qual é possível reutilizar dados de um
# stage anterior
FROM alpine:3.18

# O comando RUN serve para executar qualquer instrução, neste caso estamos ainstalando o pacote
# ca-certificates e ignorando caso haja algo no cache.
RUN apk --no-cache add ca-certificates

# O comando WORKDIR serve para criar um diretório e caso não exista e também define como o diretório
# principal após acessar o container
WORKDIR /app

# O comando COPY server parar copiar os arquivos, no exemplo
COPY --from=builder /app/bin/api .

# O comando ENTRYPOINT define um ponto de entrada do container (server para definir o processo principal)
ENTRYPOINT /app/api