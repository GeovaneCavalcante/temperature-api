
# Temperature API - README.md

## Introdução

Bem-vindo ao repositório da Temperature API, uma solução simples e eficiente para gerenciar e acessar dados de temperatura. Este README serve como um guia para configurar e utilizar a aplicação.

## Pré-requisitos

Antes de começar, certifique-se de ter as seguintes ferramentas instaladas:

- Docker
- Make
- Git (para clonar o repositório)

## Configuração

Siga os passos abaixo para configurar o ambiente do projeto.

### Clonando o Repositório

Primeiro, clone o repositório da Temperature API para a sua máquina local:

```bash
git clone https://github.com/GeovaneCavalcante/temperature-api
cd temperature-api
```

### Construindo a Imagem Docker

Para construir a imagem Docker do projeto, execute o seguinte comando:

```bash
make docker-build-image
```

Este comando irá construir uma imagem Docker com todas as dependências necessárias para executar a aplicação.

### Subindo o Projeto

Para subir o projeto, você tem duas opções:

1. **Usando o Docker Compose:**

   Execute o comando abaixo para subir o projeto com o Docker Compose. Isso iniciará todos os serviços necessários definidos no `docker-compose.yml`.

   ```bash
   make docker-up
   ```

2. **Rodando Diretamente o Container:**

   Se preferir, você pode executar diretamente o container Docker:

   ```bash
   make docker-run
   ```

### Modo Desenvolvimento

Para trabalhar no modo de desenvolvimento:

1. **Construa a imagem e suba o projeto:**

   Siga os comandos abaixo para construir a imagem e subir o projeto:

   ```bash
   make docker-build-image
   make docker-up
   ```

2. **Acessando o Bash do Container:**

   Após subir o projeto, você pode acessar o bash do container para executar a aplicação ou realizar outras atividades de desenvolvimento:

   ```bash
   make open-bash
   ```

   Dentro do bash, você terá um ambiente configurado onde pode executar a aplicação como desejar.

## Executando Testes

Para executar os testes da aplicação, utilize o comando:

```bash
make test
```

Para gerar a cobertura de código e um relatório em `coverage.html`, execute:

```bash
make test-coverage
```

Após a execução, o arquivo `coverage.html` será gerado na raiz do projeto, permitindo que você visualize detalhadamente a cobertura de testes.

## Uso da API

A Temperature API está disponível no seguinte link: [Temperature API](https://temperature-api-mszsnkdbwq-uc.a.run.app/temperature?cep=65909001)

Para utilizar, passe o CEP desejado via query params, por exemplo: `https://temperature-api-mszsnkdbwq-uc.a.run.app/temperature?cep=SEU_CEP_AQUI`.
