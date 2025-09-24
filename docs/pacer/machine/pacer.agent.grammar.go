package main

import "fmt"

type CommandType string

const (
	CommandStart    CommandType = "start"
	CommandReview   CommandType = "review"
	CommandDone     CommandType = "done"
	CommandAssign   CommandType = "assign"
	CommandBlock    CommandType = "block"
	CommandUnblock  CommandType = "unblock"
	CommandNote     CommandType = "note"
	CommandSetdod   CommandType = "setdod"
	CommandSetfield CommandType = "setfield"
	CommandRollback CommandType = "rollback"
)

type Command struct {
	Type     CommandType
	Start    bool
	Review   bool
	Done     bool
	Assign   bool
	Block    bool
	Unblock  bool
	Note     bool
	Setdod   bool
	Setfield bool
	Rollback bool
}

func main() {
	command := Command{
		Type:   CommandStart,
		Start:  true,
		Review: false,
		Done:   false,
		Assign: false,
		Block:  false,
	}
	fmt.Println(command)
}
