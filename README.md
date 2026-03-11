# GitHub Gist API

A simple HTTP API written in **Go** that retrieves the list of **public GitHub gists** for a given user.

The service queries the GitHub API and returns a simplified JSON response.

---

## Endpoint

```
GET /<USER>
```

Example:

```
curl http://localhost:8080/octocat
```

Health check:

```
GET /health
```

---

## Run Locally

Requires **Go 1.26.1**

```
go run main.go
```

Server starts on:

```
http://localhost:8080
```

---

## Run Tests

```
go test ./...
```

---

## Build Docker Image

```
docker build -t gist-api:v1 .
```

---

## Run Container

```
docker run \
--name gist-api \
-p 8080:8080 \
--read-only \
--cap-drop=ALL \
--security-opt=no-new-privileges \
gist-api:v1
```

Test:

```
curl http://localhost:8080/octocat
```
