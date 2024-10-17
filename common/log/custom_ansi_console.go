package log

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type MessageType int

const (
	Success MessageType = iota
	Error
	Debug
	Info
	Warn
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
		color.NoColor = true // Disable fatih/color output globally.
	} else {
		color.NoColor = false
	}
	return &CustomAnsiConsole{writer: writer}
}

func (c *CustomAnsiConsole) Markup(value string) {
	fmt.Fprint(c.writer, value)
}

func (c *CustomAnsiConsole) MarkupLine(value string) {
	fmt.Fprintln(c.writer, value)
}

func (c *CustomAnsiConsole) PrintMessage(messageType MessageType, message string) {
	var col *color.Color

	switch messageType {
	case Success:
		col = color.New(color.FgGreen)
	case Error:
		col = color.New(color.FgRed)
	case Debug:
		col = color.New(color.FgHiBlack)
	case Info:
		col = color.New(color.FgHiGreen)
	case Warn:
		col = color.New(color.FgYellow)
	}

	col.Add(color.Underline)
	col.Fprintln(c.writer, message)
}

func (c *CustomAnsiConsole) SuccessMessage(message string) {
	c.PrintMessage(Success, message)
}

func (c *CustomAnsiConsole) ErrorMessage(message string) {
	c.PrintMessage(Error, message)
}

func (c *CustomAnsiConsole) DebugMessage(message string) {
	c.PrintMessage(Debug, message)
}

func (c *CustomAnsiConsole) InfoMessage(message string) {
	c.PrintMessage(Info, message)
}

func (c *CustomAnsiConsole) WarnMessage(message string) {
	c.PrintMessage(Warn, message)
}
