package log

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// NonAnsiWriter is a custom writer that removes ANSI escape sequences
type NonAnsiWriter struct {
	lastOut string
}

func (w *NonAnsiWriter) Write(p []byte) (n int, err error) {
	value := string(p)
	if w.lastOut == value {
		return len(p), nil
	}
	w.lastOut = value
	output := w.removeAnsiEscapeSequences(value)
	if strings.TrimSpace(output) == "" {
		return len(p), nil
	}
	return fmt.Print(output)
}

func (w *NonAnsiWriter) removeAnsiEscapeSequences(input string) string {
	output := regexp.MustCompile(`\x1B\[(\d+;?)+m`).ReplaceAllString(input, "")
	output = regexp.MustCompile(`\[\??\d+[AKlh]`).ReplaceAllString(output, "")
	output = regexp.MustCompile(`[\r\n] +`).ReplaceAllString(output, "")
	return output
}

// CustomAnsiConsole represents a console capable of writing ANSI escape sequences
type CustomAnsiConsole struct {
	writer io.Writer
}

func NewCustomAnsiConsole(forceAnsi, noAnsiColor bool) *CustomAnsiConsole {
	var writer io.Writer = os.Stdout
	if noAnsiColor {
		writer = &NonAnsiWriter{}
	}
	return &CustomAnsiConsole{writer: writer}
}

func (c *CustomAnsiConsole) Markup(value string) {
	fmt.Fprint(c.writer, value)
}

func (c *CustomAnsiConsole) MarkupLine(value string) {
	fmt.Fprintln(c.writer, value)
}
