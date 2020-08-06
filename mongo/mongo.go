package mongo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Client ...
type (
	Client interface {
		CreateEmployee(Employee) error
		GetEmployees() ([]Employee, error)
		GetEmployeeByID(id string) (Employee, error)
		Delete(id string) error
		Update(id string, employee Employee) error
	}
	client struct {
		ms *mgo.Session
		db string
	}
)

// COLLECTION name
const (
	COLLECTION = "goapi"
)

// Employee ...
type Employee struct {
	ID   string `bson:"_id", json:"id"`
	Name string `bson:"name", json :"name"`
	Age  int    `bson:"age", json :"age"`
}

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

// CreateEmployee ...
func (c *client) CreateEmployee(employee Employee) (err error) {

	sess := c.ms.Copy()
	defer sess.Close()

	coll := sess.DB(c.db).C(COLLECTION)
	err = coll.Insert(&Employee{
		ID:   employee.ID,
		Name: employee.Name,
		Age:  employee.Age,
	})

	if err != nil {
		panic(err)
	}

	return err
}

// GetEmployees ...
func (c *client) GetEmployees() ([]Employee, error) {

	sess := c.ms.Copy()
	defer sess.Close()

	var employees []Employee
	err := sess.DB(c.db).C(COLLECTION).Find(bson.M{}).All(&employees)
	return employees, err
}

// GetEmployeeByID ...
func (c *client) GetEmployeeByID(id string) (Employee, error) {
	sess := c.ms.Copy()
	defer sess.Close()
	var employee Employee
	err := sess.DB(c.db).C(COLLECTION).FindId(id).One(&employee)
	return employee, err
}

// Delete ...
func (c *client) Delete(id string) error {

	sess := c.ms.Copy()
	defer sess.Close()

	err := sess.DB(c.db).C(COLLECTION).RemoveId(id)
	return err

}

// Update ...
func (c *client) Update(id string, employee Employee) error {
	sess := c.ms.Copy()
	defer sess.Close()

	err := sess.DB(c.db).C(COLLECTION).UpdateId(id, &employee)
	return err
}
