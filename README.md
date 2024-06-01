# simple-api-rbac

# How To run
```docker-compose up --wait```
```docker-compose up ```

# Available End Point

for post, put , delete need auth

```[GET] /books```
response 
```
{
    "count": 3,
    "result": [
        {
            "ID": 1,
            "CreatedAt": "2024-06-01T02:39:09Z",
            "UpdatedAt": "2024-06-01T02:39:09Z",
            "DeletedAt": null,
            "title": "the land",
            "author": "yasir",
            "isbn": "1dfsafd-fdsafdsa",
            "published_date": ""
        },
        {
            "ID": 3,
            "CreatedAt": "2024-06-01T04:36:24Z",
            "UpdatedAt": "2024-06-01T04:36:24Z",
            "DeletedAt": null,
            "title": "the king",
            "author": "yasir",
            "isbn": "1dfsafd-fdsafdsa",
            "published_date": ""
        },
        {
            "ID": 4,
            "CreatedAt": "2024-06-01T04:38:01Z",
            "UpdatedAt": "2024-06-01T04:38:01Z",
            "DeletedAt": null,
            "title": "the king",
            "author": "yasir",
            "isbn": "1dfsafd-fdsafdsa",
            "published_date": ""
        }
    ]
}
```
    
```[GET] /books/{id} ```
response
```
{
    "count": 1,
    "result": {
        "ID": 1,
        "CreatedAt": "2024-06-01T02:39:09Z",
        "UpdatedAt": "2024-06-01T02:39:09Z",
        "DeletedAt": null,
        "title": "the land",
        "author": "yasir",
        "isbn": "1dfsafd-fdsafdsa",
        "published_date": ""
    }
}
```

```[POST] /books```
request 
```
{
    "title": "the king",
    "author": "yasir",
    "isbn": "1dfsafd-fdsafdsa",
    "publish_date": "2024-03-04"

}
response 
```
{
    "result": {
        "ID": 3,
        "CreatedAt": "2024-06-01T04:36:24.404405091Z",
        "UpdatedAt": "2024-06-01T04:36:24.404405091Z",
        "DeletedAt": null,
        "title": "the king",
        "author": "yasir",
        "isbn": "1dfsafd-fdsafdsa",
        "published_date": ""
    }
}
```

```[PUT] /books/{id}```
request 
```
{
    "title": "the king",
    "author": "yasir",
    "isbn": "1dfsafd-fdsafdsa",
    "publish_date": "2024-03-04"

}
```
response 
```
{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "title": "the king",
        "author": "yasir",
        "isbn": "1dfsafd-fdsafdsa",
        "published_date": ""
    },
    "result": "successfully updated data"
}
```

```[DELETE] /books/{id}```
response
```
{
    "result": "Data deleted successfully"
}
```
note : 
1. when we do docker compose-up there is flow to run unit test
