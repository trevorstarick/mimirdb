package mimirdb

import (
	"net"
	"bufio"
	"net/textproto"
	"fmt"
	"strings"
	"io"
	"github.com/satori/uuid"
	"strconv"
)

type Connection struct {
	net.Conn

	Open bool
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
	return
}


type Command struct {
	Prefix string
	Args []string
}

func (c *Command) Parse(line string) Command {
	l := strings.SplitN(line, " ", -1)

	if len(l) > 1 {
		return Command{strings.ToUpper(l[0]), l[1:]}
	} else {
		return Command{strings.ToUpper(line), []string{}}
	}
}

func (c *Command) String() string {
	return fmt.Sprintf("%v { %v }", c.Prefix, strings.Join(c.Args, " "))
}


type Server struct {
	Addr string

	isDebug bool
	//ReadTimeout time.Duration
	//WriteTimeout time.Duration
	//IdleTimeout time.Duration
}

func (s *Server) handleConn(conn Connection) {
	fmt.Printf("[info] new connection from %v\n", conn.RemoteAddr())
	reader := textproto.NewReader(bufio.NewReader(conn))

	conn.Open = true

	for conn.Open {
		fmt.Fprint(conn, "\r> ")
		var cmd Command

		ln, err := reader.ReadLine()
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			conn.End()
		}

		cmd = cmd.Parse(ln)

		// https://redis.io/commands
		switch cmd.Prefix {
		case "BIND":
			if len(cmd.Args) < 2 {
				fmt.Fprintf(conn, "ERR: %v\r\n", "Not enough arguments")
			} else {
				fmt.Fprintf(conn, "Bound %v, and %v together\r\n",
					strings.Join(cmd.Args[:len(cmd.Args) - 1], ", "),
					cmd.Args[len(cmd.Args) - 1],
				)
			}
		case "EXISTS":
			var results = make([]bool, len(cmd.Args))

			fmt.Fprintf(conn, "%v\r\n", "[")
			for _, result := range results[:len(results) - 1] {
				fmt.Fprintf(conn, "  %v,\r\n", result)
			}

			fmt.Fprintf(conn, "  %v\r\n", results[len(results) - 1])
			fmt.Fprintf(conn, "%v\r\n", "]")
		case "GET":
			if len(cmd.Args) == 0 {
				fmt.Fprintf(conn, "ERR: %v\r\n", "Nothing to get")
			} else {
				fmt.Fprintf(conn, "%v\r\n", "GET")
			}
		case "INFO":
			fmt.Fprintf(conn, "%v\r\n", "INFO")
		case "SET":
			fmt.Fprintf(conn, "%v\r\n", cmd.Args)
		case "UUID":
			iter := 1
			if len(cmd.Args) > 0 {
				iter, err = strconv.Atoi(cmd.Args[0])
				if err != nil {
					iter = 1
				}
			}

			for i := 0; i < iter; i += 1{
				fmt.Fprintf(conn, "%v\r\n", uuid.NewV4().String())
			}
		case "QUIT":
			fmt.Fprintf(conn, "%v\r\n", "good bye!")
			conn.End()
		default:
			fmt.Printf("%v: %v\n", conn.RemoteAddr(), cmd.String())
		}
	}
}

func (s *Server) Serve (ln net.Listener) error {
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}

		c := Connection{
			Conn: conn,
		}

		go s.handleConn(c)

	}
}

func (s *Server) ListenAndServe (addr string) error {
	if addr == "" {
		addr = ":http"
	}

	ln, err := net.Listen("tcp", addr)

	if err != nil {

		return err

	}

	return s.Serve(ln)
}