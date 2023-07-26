# shrtnr
Link shortener in Go + Redis, built as a rite of passage.

# How to build and run
* Configure in `.env`
* Run `docker compose up`

# API
### `POST /shorten`
Request:
```
{"link": "${full link"}}
```

Response:
```
200 {"short_link": "${URL_PREFIX}/${shortened link}"}
```
or
```
400 {"error": "${error message}"}
```
or
```
500 {"error": "${error message}"}
```

### `GET /unshorten/{link}`
Response:
```
200 {"full_link": "${full link}"}
```
or
```
404 {"error": "No such link"}
```
or
```
500 {"error": "${error message}"}
```


### `GET /{link}`
Response:

`Redirect to the URL`

or
```
404 {"error": "No such link"}
```
or
```
500 {"error": "${error message}"}
```