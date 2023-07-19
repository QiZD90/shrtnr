# shrtnr
Link shortener in Go + redis, built as a rite of passage.

# How to build and run
* Configure your `URL_PREFIX` env for `shrtnr` container
* Run `docker compose up`

# API
### `POST /shorten` with POST parameter `url`
Response:
```
{"short_link": "${URL_PREFIX}/${shortened link}"}
```

### `GET /unshorten/{link}`

Response:
```
{"link": "${full link}"}
```
or
```
{"error": "No such link"}
```

### `GET /{link}`
Response:

`Redirect to the URL`

or

```
{"error": "No such link"}
```