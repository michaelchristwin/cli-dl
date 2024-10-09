package log

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync"
)

// NonAnsiWriter removes ANSI escape sequences from output
type NonAnsiWriter struct {
	out     io.Writer
	lastOut string
	mu      sync.Mutex
}

func NewNonAnsiWriter(w io.Writer) *NonAnsiWriter {
	return &NonAnsiWriter{
		out: w,
	}
}

func (w *NonAnsiWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	input := string(p)
	if w.lastOut == input {
		return len(p), nil
	}
	w.lastOut = input

	output := w.removeAnsiEscapeSequences(input)
	if strings.TrimSpace(output) == "" {
		return len(p), nil
	}

	return w.out.Write([]byte(output))
}

func (w *NonAnsiWriter) removeAnsiEscapeSequences(input string) string {
	patterns := []string{
		`\x1B\[(\d+;?)+m`,
		`\[\??\d+[AKlh]`,
	}

	output := input
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)
		output = re.ReplaceAllString(output, "")
	}

	// Remove line breaks followed by spaces
	re := regexp.MustCompile(`[\r\n] +`)
	output = re.ReplaceAllString(output, "")

	return output
}

// CustomAnsiConsole provides ANSI-aware console output
type CustomAnsiConsole struct {
	out         io.Writer
	forceAnsi   bool
	noAnsiColor bool
}

var defaultConsole = &CustomAnsiConsole{
	out: os.Stdout,
}

// InitConsole initializes the console with specified settings
func InitConsole(forceAnsi, noAnsiColor bool) {
	var writer io.Writer = os.Stdout
	if noAnsiColor {
		writer = NewNonAnsiWriter(os.Stdout)
	}

	defaultConsole = &CustomAnsiConsole{
		out:         writer,
		forceAnsi:   forceAnsi,
		noAnsiColor: noAnsiColor,
	}
}

// Markup writes the specified markup to the console
func Markup(format string, a ...interface{}) {
	defaultConsole.markup(false, format, a...)
}

// MarkupLine writes the specified markup followed by a newline to the console
func MarkupLine(format string, a ...interface{}) {
	defaultConsole.markup(true, format, a...)
}

func (c *CustomAnsiConsole) markup(newline bool, format string, a ...interface{}) {
	// This is a simplified markup implementation
	// You may want to add more sophisticated markup parsing here
	text := fmt.Sprintf(format, a...)

	if newline {
		fmt.Fprintln(c.out, text)
	} else {
		fmt.Fprint(c.out, text)
	}
}

// Additional utility functions can be added here as needed

// Example color constants
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
)
