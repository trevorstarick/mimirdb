package connection

import (
	"net"
	"fmt"
)

type Connection struct {
	net.Conn

	Open bool

	IsTransaction bool
	TransactionInitial string
	TransactionOperand string
	TransactionStorage []string
}

func (c *Connection) New (conn net.Conn) Connection {
	var connection Connection
	connection.Conn = conn

	return connection
}

func (c *Connection) End () {
	fmt.Printf("[info] connection from %v closed\n", c.RemoteAddr())
	c.Close()
	c.Open = false

}