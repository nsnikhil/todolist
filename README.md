# TODOLIST

## Description
A task management application.

## Setup
This service runs on go. Follow instructions on https://golang.org/doc/install for setting up Go environment.
Checkout the code, install the dependencies and build the project and do testing:

```
git clone github.com/nsnikhil/todolist
cd ${todolist}
make deps
make test
make serve
```

## DB Migration

```
$ make migrate
```

## DB Rollback

```
$ make rollback
```