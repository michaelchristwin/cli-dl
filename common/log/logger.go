package log

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Logger struct {
	LogLevel     LogLevel
	IsWriteFile  bool
	LogFilePath  *string
	VarsRepRegex *regexp.Regexp
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

func GetCommandLine() string {
	return fmt.Sprintf("%s %s", os.Args[0], filepath.Join(os.Args[1:]...))
}
