package GoMicroColor

import (
	"fmt"
	"io"
	"os"
)

type ColorWriter struct {
	io.Writer // embedded interface (its not neccesarry but its good for reading)
	output    *os.File
	fgColor   TerminalColorType
	bgColor   TerminalColorType
}

func NewColorWriter(f *os.File, fgColor TerminalColorType, bgColor TerminalColorType) *ColorWriter {
	return &ColorWriter{output: f, fgColor: fgColor, bgColor: bgColor}
}

func (cw ColorWriter) Write(p []byte) (n int, err error) {
	lastIndex := len(p) - 1
	if p[lastIndex] == '\n' {
		if _, err = fmt.Fprintf(cw.output, "\033[%d;5;%dm%s\033[0m\n", cw.fgColor, cw.bgColor, string(p[:lastIndex])); err != nil {
			return
		}

	} else {
		if _, err = fmt.Fprintf(cw.output, "\033[%d;5;%dm%s\033[0m", cw.fgColor, cw.bgColor, string(p)); err != nil {
			return
		}
	}
	return len(p), nil

}
