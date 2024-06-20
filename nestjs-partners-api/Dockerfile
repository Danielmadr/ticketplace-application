FROM node:21-slim

# Instala as dependências e limpa o cache do APT em um único comando RUN
RUN apt update && apt install -y openssl procps 

RUN npm install -g @nestjs/cli@10.3.2

WORKDIR /home/node/app

# Define o usuário 'node' para executar o container
USER node

# O comando para iniciar a aplicação (ajuste conforme necessário)
CMD tail -f /dev/null
