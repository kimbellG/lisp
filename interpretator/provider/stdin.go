package provider

import (
	"fmt"
	"os"

	"github.com/kimbellG/lisp/interpretator/provider/command"
)

//TerminalProvider is implementaion of CmdReader. The lines are received through the terminal.
type TerminalProvider struct {
	reader command.CmdReader
}

func (tp *TerminalProvider) initReader() {
	if tp.reader != nil {
		tp.reader = &command.CmdReaderImpl{}
	}
}

//ReadCommand reads command from stdin and return finished command
func (tp *TerminalProvider) ReadCommand() (string, error) {
	tp.initReader()

	cmd, err := tp.reader.GetCommand(os.Stdin)
	if err != nil {
		return "", fmt.Errorf("failed to get command from stdin: %w", err)
	}

	return cmd, nil
}
