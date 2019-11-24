# TODOLIST

## Description
A task management application.

## Setup
This service runs on go. Follow instructions on https://golang.org/doc/install for setting up Go environment.
Checkout the code, install the dependencies and build the project and do testing:

```
git clone github.com/nsnikhil/todolist
cd ${todolist}
make setup
```

## Application

```
$ make serve
```

## Docker Start
#### Starts the application and postgres in a container
```
$ make docker-serve
```

## Kubernetes 
#### Starts the application and postgres in a minikube
```
$ make k8-serve
```

## DB Migration

```
$ make migrate
```

## DB Rollback

```
$ make rollback
```

## Disclaimer
- Helm chart is still in development.
- The `constant` and `apperror` package needs cleanup.
- The config command only outputs the address the app will run on.