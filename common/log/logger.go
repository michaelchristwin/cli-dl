package log

import (
	"fmt"
	"html"
	"os"
	"path/filepath"
	"regexp"
	"sync"
	"time"
)

type Logger struct {
	LogLevel     LogLevel
	IsWriteFile  bool
	LogFilePath  *string
	VarsRepRegex *regexp.Regexp
	LogWriteLock sync.RWMutex
}

func (l *Logger) InitLogFile() {
	if !l.IsWriteFile {
		return
	}
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Get the directory of the executable
	exeDir := filepath.Dir(exePath)
	logDir := filepath.Join(exeDir, "Logs")

	// Create the directory if it doesn't exist
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		fmt.Println("Error creating Logs directory:", err)
		return
	}

	now := time.Now() // Safe filename format
	logFilePath := filepath.Join(logDir, fmt.Sprintf("%s.log", now.Format("2006-01-02_15-04-05")))
	index := 1
	fileName := filepath.Base(logFilePath)
	commandline := GetCommandLine()
	init := fmt.Sprintf(
		"LOG %s\nSave Path: %s\nTask Start: %s\nTask Commandline: %s\n\n",
		now.Format("yyyy-MM-dd_HH-mm-ss-fff"),
		filepath.Dir(logFilePath),
		now.Format("yyyy/MM/dd HH:mm:ss"),
		commandline)
	for {
		if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
			break // File does not exist, exit the loop
		}
		// Generate new log file path with incremented index
		logFilePath = filepath.Join(filepath.Dir(logFilePath), fmt.Sprintf("%s-%d.log", fileName, index))
		index++
	}
	if err := os.WriteFile(logFilePath, []byte(init), 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func (l *Logger) GetCurrTime() string {
	return time.Now().Format("15:04:05.000")
}

func (l *Logger) HandleLog(write string, subWrite string) {
	console := NewCustomAnsiConsole(true, false)

	if write == "" {
		console.MarkupLine(write)
	} else {
		console.MarkupLine(write)
		fmt.Println(subWrite)
	}
	if l.IsWriteFile && fileExists(l.LogFilePath) {
		plain := fmt.Sprintf("%s%s", html.EscapeString(write), html.EscapeString(subWrite))
		l.LogWriteLock.Lock()
		defer l.LogWriteLock.Unlock()
		// Open the file in append mode. If it doesn't exist, create it with 0644 permissions.
		file, err := os.OpenFile(*l.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening or creating file:", err)
			return
		}
		defer file.Close() // Ensure the file is closed properly

		// Write the content to the file with a newline at the end
		if _, err := file.WriteString(plain + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

	}
}

func (l *Logger) ReplaceVars(data string, ps []interface{}) string {
	for _, p := range ps {

		data = l.VarsRepRegex.ReplaceAllString(data, fmt.Sprintf("%s", p))
	}
	return data
}

func GetCommandLine() string {
	return fmt.Sprintf("%s %s", os.Args[0], filepath.Join(os.Args[1:]...))
}

func fileExists(filePath *string) bool {
	if filePath == nil || *filePath == "" {
		fmt.Println("Invalid or nil file path.")
		return false
	}

	_, err := os.Stat(*filePath)
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return err == nil // File exists and no other error occurred
}
