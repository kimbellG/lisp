package command

import "io"

//CmdReader is interface of command provider module, which reader line from standard reader interface
type CmdReader interface {
	GetCommand(reader io.Reader) (string, error)
}
