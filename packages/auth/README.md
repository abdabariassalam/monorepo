# nodejs-api-message

## Install dependencies.

`$ npm install`

## jalankan project

`$ npm start`

## Testing 

script berada di `tes/messages.test.js`

Run : `$ npm test`

## test function

### Post
untuk mencoba api sending message bisa menggunkan **postman**

atau menggunakan syntax curl dibawah ini

```
curl --location --request POST 'http://localhost:8123/messages' \
--header 'Content-Type: text/plain' \
--data-raw **message**
```

syntax diatas untuk parameter body yang dikirimkan text atau string

jika ingin menggunakan parameter JSON

```
curl --location --request POST 'http://localhost:8123/messages' \
--header 'Content-Type: application/json' \
--data-raw '{
	"message": **message**	
}'
```

tinggal ganti **message** dengan pesan yang diinginkan 

harus menggunakan `double quotation mark " "` untuk content-type nya json

bisa menggunakan `double quotation mark ataupun single quotation mark ' ' ` 

### Get

untuk melihat semua pesan yang telah di kumpul kan tinggal panggil

`localhost:8123/messages`

ke browser atau menggunakan **postman** dengan menggunakan get request

jika ingin menggunakan curl

```
curl --location --request GET 'http://localhost:8123/messages' \
--header 'Content-Type: application/json'
```

### websocket

gunakan console pada browser google chrome

salin script do bawah ini

```
let c = new WebSocket('ws://localhost:8123/ws');
c.onmessage = (msg) => {
	let data = JSON.parse(msg.data);
	console.log(data);
};
```

# Semoga bermanfaat
