package command

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

//CmdReaderImpl is implementation of cmdReader interface
type CmdReaderImpl struct {
	countOfOpenOperation int
	result               strings.Builder
}

//GetCommand constructs finished command from the reader lines.
//It clears away unnecessary spaces and character of new line.
func (cri *CmdReaderImpl) GetCommand(reader io.Reader) (string, error) {
	defer cri.result.Reset()
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("scan string: %w", err)
		}

		str := strings.TrimSpace(scanner.Text())

		if err := cri.parseInputString(str); err != nil {
			return "", fmt.Errorf("invalid string %s: %w", str, err)
		}

		if cri.isEndOfCommand() {
			break
		}

		cri.result.WriteRune(' ')
	}

	return cri.result.String(), nil
}

func (cri *CmdReaderImpl) parseInputString(input string) error {
	var isPreviousSpace bool
	for _, r := range input {
		if unicode.IsSpace(r) {
			if isPreviousSpace {
				continue
			}
			isPreviousSpace = true
		} else {
			isPreviousSpace = false
		}

		if err := cri.parseRune(r); err != nil {
			return fmt.Errorf("parsing rune: %w", err)
		}
	}

	return nil
}

func (cri *CmdReaderImpl) parseRune(r rune) error {
	switch {
	case cri.isStartOperation(r):
		cri.countOfOpenOperation++
	case cri.isEndOperation(r):
		cri.countOfOpenOperation--
	}

	_, err := cri.result.WriteRune(r)
	if err != nil {
		return fmt.Errorf("don't write rune in result builder: %w", err)
	}

	return nil
}

func (cri *CmdReaderImpl) isStartOperation(r rune) bool {
	return r == '('
}

func (cri *CmdReaderImpl) isEndOperation(r rune) bool {
	return r == ')'
}

func (cri *CmdReaderImpl) isEndOfCommand() bool {
	return cri.countOfOpenOperation == 0
}
