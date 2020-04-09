# swagger-sample

## Let's run the server:

```
$ go run ./cmd/greeter/main.go --port 3000
```

## Let's now try to call our only defined operation, using the [httpie](https://httpie.org/) cli:

``` shell
$ http get :3000/hello
```

## Hurray, let's now greet Swagger:

``` shell
$ http get :3000/hello name==Swagger
```
