FROM node:14-alpine

RUN yarn global add serve

WORKDIR /app

COPY ./frontend/package*.json ./

RUN yarn install

COPY ./frontend/ .
COPY ./frontend/.env .

RUN yarn build

CMD ["serve", "-s", "build", "-l", "tcp://0.0.0.0:10005" ]
