# Shop application 
### with REST-webapi as public interface and gRPC as internals services

![alt text](https://github.com/a-korkin/shop/blob/main/img/schema.jpg?raw=true)

```console
curl localhost:8080/items?page=1&limit=10
curl localhost:8080/items/1
curl -X POST localhost/8080/items \
    -d '{"title": "item 1", "price": 33.52, "category": "toys"}'
```
