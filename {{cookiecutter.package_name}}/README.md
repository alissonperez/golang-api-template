# {{cookiecutter.project_name}}

## Dependencies

- **Go 1.19**: See how to setup and basic usage it [here](https://go.dev/doc/install).
- **Task**: More about [here](https://taskfile.dev/).

## Setup / Install

First, just install [**Task**](https://taskfile.dev/). Usually, just `go install github.com/go-task/task/v3/cmd/task@latest`, other options [here](https://taskfile.dev/#/installation)

Then, just run `setup` task to tidy and install dependencies:

```
$ task setup
```

## How to run

Using just start using task command:

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
