Собственный Golang (1.18+) JSONRPC-сервер по спецификации

https://www.jsonrpc.org/specification

Сервер должен уметь приимать пост запрос ->

curl -X POST -H 'Content-Type: application/json' -i http://localhost:8000/ --data '{"id": "385ad292-9830-4f14-bd4c-9118d85f0db8", "jsonrpc": "2.0", "method": "greeting", "params": { "name": "Vasy" }}'

И отдавать в ответ 

{"id": "385ad292-9830-4f14-bd4c-9118d85f0db8", "jsonrpc": "2.0", "result": "Hello, Vasy"}
