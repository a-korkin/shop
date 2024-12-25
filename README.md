# Shop application 
### with REST-webapi as public interface and gRPC as internals services

![alt text](https://github.com/a-korkin/shop/blob/main/img/schema.jpg?raw=true)

Examples of requests:
```console
curl localhost:8080/items?page=1&limit=10

curl localhost:8080/items/1

curl -X POST localhost:8080/items \
    -d '{"title": "item 1", "price": 33.52, "category": "toys"}'

curl localhost:8080/users/1

curl -X POST localhost:8080/buy \
    -d '{"user_id": 1, "item_id": 1, "count_items": 32}'
```
