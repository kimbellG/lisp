package command

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneParse(t *testing.T) {
	tt := []struct {
		reader    CmdReaderImpl
		runes     []rune
		want      string
		wantCount int
	}{
		{
			CmdReaderImpl{
				countOfOpenOperation: 0,
				result:               strings.Builder{},
			},

			[]rune{'z', 'b', '(', '(', '(', 'z', 'd', 'v'},
			"zb(((zdv",
			3,
		},
		{
			CmdReaderImpl{
				countOfOpenOperation: 0,
				result:               strings.Builder{},
			},

			[]rune{'z', 'b', '(', '(', '(', ')', ')', 'z', 'd', 'v'},
			"zb((())zdv",
			1,
		},
	}

	for _, tc := range tt {
		for _, r := range tc.runes {
			if err := tc.reader.parseRune(r); err != nil {
				t.Errorf("failed to parse rune: %v", err)
				continue
			}
		}

		assert.Equal(t, tc.want, tc.reader.result.String())
		assert.Equal(t, tc.wantCount, tc.reader.countOfOpenOperation)
	}
}

func TestStringParse(t *testing.T) {
	tt := []struct {
		msg           string
		reader        CmdReaderImpl
		input         string
		want          string
		isFinishedCmd bool
	}{
		{
			"for right input",
			CmdReaderImpl{},
			"(+ 4 5)",
			"(+ 4 5)",
			true,
		},
		{
			"for a lot of space",
			CmdReaderImpl{},
			"(+           4 5)",
			"(+ 4 5)",
			true,
		},
		{
			"for not finished command",
			CmdReaderImpl{},
			"(+ (+ 2 5) (* 2 3)",
			"(+ (+ 2 5) (* 2 3)",
			false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.msg, func(t *testing.T) {

			assert.NoErrorf(t, tc.reader.parseInputString(tc.input), "failed to parse string %s", tc.input)

			assert.Equal(t, tc.want, tc.reader.result.String())
			assert.Equal(t, tc.isFinishedCmd, tc.reader.isEndOfCommand())
		})
	}
}

func TestGetCommand(t *testing.T) {
	tt := []struct {
		msg    string
		reader CmdReaderImpl
		buf    *bytes.Buffer
		want   string
	}{
		{
			"simple command",
			CmdReaderImpl{},
			bytes.NewBufferString("(+ 3 5)"),
			"(+ 3 5)",
		},
		{
			"huge command",
			CmdReaderImpl{},
			bytes.NewBufferString("(+ (* 3\n\t(+ (* 2 4)\n\t(+ 3 5)))\n\t(+ (- 10 7)\n\t6))"),
			"(+ (* 3 (+ (* 2 4) (+ 3 5))) (+ (- 10 7) 6))",
		},
	}

	for _, tc := range tt {
		t.Run(tc.msg, func(t *testing.T) {
			cmd, err := tc.reader.GetCommand(tc.buf)
			assert.NoError(t, err, "get command is failed %s", tc.buf.String())
			assert.Equal(t, tc.want, cmd)
		})
	}
}
