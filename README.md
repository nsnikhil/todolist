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
to use helm for deployment 

```
$ make helm-serve
```
##### The application runs on $(minikube ip):nodeport when started using kubernetes


## DB Migration

```
$ make migrate
```

## DB Rollback

```
$ make rollback
```

## Know Issues
- The config command only outputs the address the app will run on.
- Id returned by add method from store is not being used in test
- Count returned by remove and update is not being used in test