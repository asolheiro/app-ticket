FROM node:22-slim

WORKDIR /home/node/app

RUN apt update && apt install -y openssl procps

COPY ./apps/partner1 ./apps/partner1
COPY ./prisma ./prisma
COPY ./libs/core ./libs/core
COPY ./.env.partner1 ./
COPY ./*.json ./

RUN npm install

CMD ["npm", "run", "migrate:partner1"]

ENTRYPOINT [ "npm", "run", "start:dev" ]
