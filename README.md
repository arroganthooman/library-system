# Simple Library System

A simple library system in Go language. Purposed to complete PT DOT Backend Engineer recruitment test.

**The key features of this system are**
User:
- Login (POST /user/login)
- Register (POST /user/register)
- Edit User Info (PUT /user/edit)
- Get User Info (GET /user/info)

Books:
- Get All Book (GET /books)
- Get Book By ID (GET /books?id=1)
- Edit Book (PUT/PATCH /book)
- Delete Book (DELETE /book)
- Borrow Book (POST /book/borrow)
- Return Book (POST /book/return)

You can download postman collection of request in https://ristek.link/LibSystemPostmanColl. After that you can import it to postman.

**Tools Used**
- Gin
- GORM
- SQLite
- Redis

## Getting started

### Prerequisities
- Golang (>=1.18)
- Redis on localhost:6379 (optional)

### Installation and Running
`git clone https://github.com/arroganthooman/library-system.git`

`cd app`

`go run app.go`

Server will run on **localhost:8080**

### Pattern
3 layers of execution are implemented in this project. Layers can be found in internal folder. Following is the top-down explanation approach for the layer:

1. Delivery Layer

This layer handles the request and response from the http request. This layer mainly validates the request that user needs to fulfill before going into the next layer. The data is passed to the usecase layer.

2. Usecase Layer

This layer handle the business usecase of the application. This layer is connected to the repository layer.

3. Repository layer

This layer purpose is to serve data for usecase layer. Reading, writing, updating and deleting data from database happens in this layer.

<br/>
The reason for using this pattern is because it's the best practice and very clean. One layer focuses one job. Also, if you're familiar with unit testing, the implementation of this code can be unit-tested very easily. Usually we use go-mock for unit testing in go. However, in effect of interfacing method we use in injecting the repository in every usecase layer, we can define our own mocking struct of usecase/repository for every layer very easily for unit testing. See example in GetAllBook usecase for the unit test for context.
