# Clean architecture in Golang
This project is a sample of a implementation Clean Architecture in GoLang.

<br/>

### Summary
The clean architecture is a model proposed by Robert C. Martin (Uncle Bob). <br/>
Briefly, this model seeks to strengthen the following points in our software:


- <span style="color:grey"> Increase expressivity of the problem that the software is solutioning.</span>
- Independence of external structures, example: Frameworks, DB, UI, WEB).
- Testability

<br/>

### Use case
The use case chosen for this sample is a service of user's management. It was selected for keep the simplicity. With it we will can: 

- [create new user]()
- [search users]()
- [remove user]()
- [set user's password]()

<br/>

### More about
- [Architecture](architecture.md)
- [Tests](tests.md)
- [Api](api.md)

<br/>

### Technologies

- [Golang](https://golang.org/)
- [Docker](https://www.docker.com/)
- [MongoDB](https://www.mongodb.com/)

<br/>

### Runing

```bash

# move to your workspace
cd workspace

# clone into your $GOPATH/src
git clone https://github.com/pablomedeiross/go-clean-architecture.git

# move to project 
cd go-clean-architecture

# run application
go run main.go 

```