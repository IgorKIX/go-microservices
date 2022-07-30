# go-microservices

Study of microservices with the help of the golang language

### Branches:
`vanilla-http` - basic REST server build with vanila golang methods

### Testing:
- `curl localhost:9090/` - GET: returns all products from server
- `curl localhost:9090/ -d '{"name": "tea", "description": "A nice cup of tea"}'` - POST: Add product to var in backend
- `curl localhost:9090/1 -XPUT -d '{"name": "tea", "description": "A nice cup of tea"}'` PUT: Update the product with given id ( id = 1 in this example)
