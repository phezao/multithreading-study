# Multithreading study with Go

This repository was created to study go routines. The project sends two http requests at the same time (using go routines) to fetch a Brazilian address hitting two APIs (Brasil API and ViaCEP). We get on stdout the response of the quickest API, with a time limit (timeout) of 1 second, therefore if any of the APIs takes more than 1 second to respond, neither of them will show on stdout.

## How to execute the project

```
go run main.go {cep}
```
