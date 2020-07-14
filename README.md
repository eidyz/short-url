# Short URL API
A little API that shortens URL's written in Go

# Documentation

## Creating new short URL
* **URL**

  `/new`

* **Method:**

  `POST`

* **Sample Call:**

```bash
curl --request POST \
  --url http://localhost:8080/new \
  --header 'content-type: application/json' \
  --data '{"fullURL": "https://google.com"}'
```

## Getting the full URL by shortened URL
* **URL**

  `/{short_url}`

* **Method:**

  `GET`

* **Sample Call:**

```bash
curl --request GET \
  --url http://localhost:8080/{short_url}
```