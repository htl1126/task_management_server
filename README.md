# Task Management Server

A task management server implemented with Go

## Introduction

The task management server (TMS) allows users to manage tasks by creation, reading, updating and deletion. TMS is a RESTful API server and I implemented it with Go 1.22.4.

## Implmentation

I used `gin-gonic` to build the server. Besides the `main` package, other packages have their own responsibilities. Please see **Table 1**.

**Table 1.**
| name      |   functionality                        |
|-----------|----------------------------------------|
| `api`     | API implementation                     |
| `model`   | data, request and response definitions |
| `storage` | task pool logic                        |
| `unittest`| unit tests                             |

Most logic is in package `storage`, so tests in `unittest` are mainly
for testing the code of package `storage`.

TMS stores tasks in the memory. We can set the task pool size with `TASKPOOLSIZE` attribute in `Dockerfile`.

## Build and run

We can directly run TMS by going to the project root directory.

```
go run main.go
```

Or we can run TMS inside a docker container. We can firstly build the docker image with,

```
sudo docker build -t api_server .
```

And run the server with,

```
sudo docker run --rm -d -p 8080:8080 api_server
```

## Unit test

We can use the following command to run all the tests with detailed results

```
go test ./... -v
```

## Usage

We can send requests with `curl`.

1. Create a task
```
curl -X POST http://localhost:8080/tasks -d '{"name":"task one"}'
```

2. List tasks
```
curl -X GET http://localhost:8080/tasks
```

3. Update a task
```
curl -X PUT http://localhost:8080/tasks/91 -d '{"name":"abc"}'
```
We might consider to use `PATCH` instead of `PUT`.

4. Delete a task
```
curl -X DELETE http://localhost:8080/tasks/10
```