# Go-HTTP

Simple Golang HTTP library for CRUD operations

## How to install

```
go get github.com/cetric32/GoHTTP
```

## Import Package

```
import "github.com/cetric32/GoHTTP"
```

## Create Instance

```
goHttp := GoHTTP.NewGoHTTP()
```

## Make a GET request

```
result, statusCode, err := goHttp.Get("https://jsonplaceholder.typicode.com/posts/1")
```

## Make a POST request

```
result, statusCode, err := goHttp.Post("https://jsonplaceholder.typicode.com/posts", nil)
```

## POST FormData 

```
result, statusCode, err := goHttp.PostForm("https://jsonplaceholder.typicode.com/posts", map[string][]string{"title": {"foo"}, "body": {"bar"}, "userId": {"1"}})
```

## Make a PUT request

```
result, statusCode, err := goHttp.Put("https://jsonplaceholder.typicode.com/posts/1", nil)
```

## Make a PATCH request

```
result, statusCode, err := goHttp.Patch("https://jsonplaceholder.typicode.com/posts/1", nil)
```

## Make a DELETE request 

```
result, statusCode, err := goHttp.Delete("https://jsonplaceholder.typicode.com/posts/1")
```

## Add Request headers

```
goHttp.AddHeaders(map[string]string{"Content-Type": "application/json", "Authorization": "Bearer <token>"})
```
