package commands

import (
	"fmt"
	. "github.com/trevorstarick/mimirdb/connection"
)

type Transactions struct {}
func (_ *Transactions) AND (conn Connection, _ []string) {
	conn.TransactionOperand = "AND"
}

func (_ *Transactions) END (conn Connection, args []string) {
	conn.IsTransaction = false
	conn.TransactionOperand = ""
	conn.TransactionStorage = []string{}

	fmt.Fprintf(conn, "%v\r\n", conn.TransactionStorage)
}

func (_ *Transactions) OR (conn Connection, _ []string) {
	conn.TransactionOperand = "OR"
}

func (_ *Transactions) NAND (conn Connection, _ []string) {
	conn.TransactionOperand = "NAND"
}

func (_ *Transactions) START (conn Connection, args []string) {
	var initial string

	if len(args) > 0 {
		for _, arg := range args {
			if arg == "ALL" || arg == "NONE" {
				initial = arg
			}
		}
	}

	conn.IsTransaction = true
	conn.TransactionInitial = initial
	conn.TransactionOperand = "OR"
	conn.TransactionStorage = []string{}
}
