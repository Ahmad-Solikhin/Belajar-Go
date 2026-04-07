package simple

import "fmt"

type Connection struct {
	*File
}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.CLose()
	}
}

func (c *Connection) CLose() {
	fmt.Println("Close connection", c.File.Name)
}
