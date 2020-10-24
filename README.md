# Avocado-Backend
Backend service for Avocado Timer

## Prerequisite
1. GO version 1.15
2. GNU Make v4.2.1
3. docker v19.x.x
4. docker-compose v1.24.0

## Development

You can build this project for dev using the following command:<br>
```
make dev
```

There is a hot reload feature that allows the project to be rebuilt and run when you save any of the go source files.<br>

You can subsequently clean up the project using the following command:<br>
```
make clean
```

## Production
You can build this project for production using the following command:<br>
```
make prod
```
