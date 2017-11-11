package commands

import (
	"fmt"
	. "github.com/trevorstarick/mimirdb/connection"
)

func NotImplemented (conn Connection, _ []string) {
	fmt.Fprintf(conn, "%v\r\n", "Not implemented yet")
}
