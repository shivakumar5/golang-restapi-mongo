# golang-restapi-mongo

This repo contains code for Golang Rest api with mongo db connection

Used Gorilla Mux package to implement a request router and dispatcher for matching incoming requests to their respective handler.


This API Contains 5 different endpoints:

   - CreateEmployee
   - GetAllEmployees
   - GetAllEmployees/{id}
   - DeleteEmployee/{id}
   - UpdateEmployee/{id}

Note that, added {id} to endpoint path. This will represent id variable that we will be able to use when we wish to return only the employee record that features that exact key.

### Mongo DB

The separate client file has been included for mongodb connection database stuff. added the directory called 'mongo' which has db connection code inside mong.go file in this directory. 

**Mongo DB Client Code:**

```go
// NewClient ...
func NewClient(addrs []string, db, user, pass string, timeout int) (Client, error) {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    addrs,
		Database: db,
		Username: user,
		Password: pass,
		Timeout:  time.Duration(timeout) * time.Second,
	})

	if err != nil {
		return nil, err
	}

	return &client{ms: session, db: db}, nil

}
```
Using this client connection, we can do CRUD Operations in mongo
