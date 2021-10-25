FROM node:10.24.1

WORKDIR /app

COPY package*.json .

RUN npm install

COPY . .

RUN echo 

EXPOSE 8123

ENTRYPOINT [ "node", "server.js" ]