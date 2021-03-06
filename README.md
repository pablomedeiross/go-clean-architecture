# Clean architecture in Golang
This project is a sample of an implementation of Clean Architecture in GoLang.

<br/>

## Summary
The clean architecture is a model proposed by [Robert C. Martin (Uncle Bob)](http://cleancoder.com/).
Briefly, this model seeks to strengthen the following points in our software:

- Increase expressivity of problem that the software is solutioning.
- Independence of external structures, example: Frameworks, DB, UI, WEB.
- Testability

<br/>

## Use case
The use case chosen for this sample is a service of user's management. It was selected for keep the simplicity. With it we'll can: 

- Create new user
- Search users
- Remove user
- Set user's password

<br/>

## More about
- [My implementation of clean architecture](my-implementation.md)
- [Tests](test-implementation-architecture.md)

<br/>

## Technologies

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/)
- [MongoDB](https://www.mongodb.com/)

<br/>

## Running

```bash

# move to your workspace
cd workspace

# clone into your $GOPATH/src
git clone https://github.com/pablomedeiross/go-clean-architecture.git

# move to project 
cd go-clean-architecture

# run application in local profile 
go run main.go --profile=local

```
