package mimirdb

import (
	. "github.com/trevorstarick/mimirdb/connection"
	. "github.com/trevorstarick/mimirdb/commands"
)

type CommandFn = func (Connection, []string)

var Commands map[string]CommandFn

func init() {
	var transactions Transactions
	Commands = make(map[string]func (Connection, []string))

	// Transactions
	Commands["AND"] = transactions.AND
	Commands["END"] = transactions.END
	Commands["OR"] = transactions.OR
	Commands["NAND"] = transactions.NAND
	Commands["START"] = transactions.START

	// Connection
	Commands["AUTH"] = NotImplemented
	Commands["LIST"] = NotImplemented
	Commands["PING"] = NotImplemented
	Commands["QUIT"] = NotImplemented
	Commands["SELECT"] = NotImplemented

	// Geo
	Commands["BOUNDS"] = NotImplemented
	Commands["POINT"] = NotImplemented

	// Time
	Commands["TRANGE"] = NotImplemented

	// Keys
	Commands["BIND"] = NotImplemented
	Commands["DEL"] = NotImplemented
	Commands["EXISTS"] = NotImplemented
	Commands["GET"] = NotImplemented
	Commands["NUKE"] = NotImplemented
	Commands["RANDOM"] = NotImplemented
	Commands["SET"] = NotImplemented
	Commands["UNBIND"] = NotImplemented
	Commands["WATCH"] = NotImplemented

	// Server Admin
	Commands["CLIENTS"] = NotImplemented
	Commands["FLUSHALL"] = NotImplemented
	Commands["FLUSHDB"] = NotImplemented
	Commands["INFO"] = NotImplemented
	Commands["SAVE"] = NotImplemented
	Commands["SHUTDOWN"] = NotImplemented
	Commands["TIME"] = NotImplemented

	// Utils
	Commands["UUID"] = NotImplemented
}