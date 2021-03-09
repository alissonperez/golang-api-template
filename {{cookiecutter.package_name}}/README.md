# {{cookiecutter.project_name}}

# Dependencies

- Go 1.15.*

## Install

```
$ go install
```

## How to run

Install [Task](https://taskfile.dev/#/installation), then run:

```
$ task
```

### Testing endpoint /hello

```
$ curl --location --request GET 'http://localhost:8000/v1/hello/123' \
--header 'Authorization: Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJjbGllbnRJZCI6MTIzLCJqdGkiOiJhNGJmYzAxZi00NmQ0LTQyNzgtOTM5ZS0yOWRkNjI2N2M5OTIiLCJpYXQiOjE2MTQ3MTIzNjMsImV4cCI6MTYxNDcxNTk2M30.4mPtjKqu8C3i0v4TEpnthp1_FyYspVwXFfa2S0EuApo'
```

## How to test

```
$ task test
```
