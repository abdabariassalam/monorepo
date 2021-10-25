# Multi-Language Monorepo

Features:

- mempunyai 2 apps dengan berbeda bahasa: nodeJs dan Golang
- menggunakan docker-compose untuk setting up dan organize multiple app
- auth-app untuk melakukan register user, login dan verify token
- fetch-app untuk melakukan aggregate data, currency converter dan verify token

## Setup explained

### Tooling

-   docker-compose:

    -   `docker-compose up --build` - running dengan perintah build atau rebuild applikasi.
    -   `docker-compose up` - running dengan settingan terakhir kali build.
    -   `docker-compose down` - stop dan remove resource.


### Included sample packages

-   **auth-app**
    -   [Express](https://github.com/expressjs/express) aplikasi.
    -   Listens on http://localhost:8123 .

    -   **fetch-app**
    -   [Gin](https://github.com/gin-gonic/gin) untuk route aplikasi.
    -   Listens on http://localhost:8080 .

### Dokumentasi Endpoint

untuk dokumentasi Api dapat di import file berikut ke postman

 [![postman.json](github.com/bariasabda/monorepo/postman.json)]

## dev

kita bisa menjalan langsung semua app sekaligus dengan menggunakan syntax berikut:
```
docker-compose up --build
```


### auth-app

untuk melakukan developmen pada auth app harus disiap kan node versi 10.24.1 untuk memastikan tidak ada perbedaan depedensi.

untuk menjalankan app nya dapat menggunakan syntax berikut ini:
    apabila kita berada di root monorepo:

    ```
    cd packages/auth/
    npm install
    npm start
    atau
    node server.js
    ```

kalau menggunakan docker-compose:

    ```
    cd packages/fetch/
    go mod tidy
    docker-compose up --build
    ```
tapi kita harus menyiapkan waktu yang sedikit lebih lama ketimbang langsung menjalan kan app karena kita harus menunggu container nya di build

### fetch-app

untuk melakukan developmen pada fetch app harus disiap kan golang versi 1.16 untuk memastikan tidak ada perbedaan depedensi.

untuk menjalankan app nya dapat menggunakan syntax berikut ini:
    apabila kita berada di root monorepo:

    ```
    cd packages/fetch/
    go mod tidy
    go run main.go
    ```

kalau menggunakan docker-compose:

    ```
    cd packages/fetch/
    go mod tidy
    docker-compose up --build
    ```
tapi kita harus menyiapkan waktu yang sedikit lebih lama ketimbang langsung menjalan kan app karena kita harus menunggu container nya di build.

## unittest

untuk unittest biasa nya dilakukan pada masing2 app

### auth-app
pada auth-app menggunakan bantuan mocha dan chai untuk melakukan unittest
```
cd packages/auth
npm test
```
### fetch-app
untuk fetch-app menggunakan bantuan mockgen segabai penyedia mocking data.
```
mockgen -source=domain/repository/repository.go -destination=domain/repository/mock/repository_mock.go -package=mock
go test -timeout 30s -coverprofile=coverprofile github.com/bariasabda/monorepo/packages/fetch/domain/service
```
untuk mockgen cukup 1 kali saja atau ketika ada perubahan atau penambahan func pada repository