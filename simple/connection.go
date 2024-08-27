package simple

import "fmt"

type Connection struct {
	File *File
}

// Close closes the current connection.
//
// No parameters.
// No return values.
func (c *Connection) Close() {
	fmt.Println("Close Connection", c.File.Name)
}

// NewConnection creates a new connection with the given file.
//
// Parameter file is a pointer to the File struct.
// Returns a pointer to the Connection struct and a function to close the connection.
func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}
}
