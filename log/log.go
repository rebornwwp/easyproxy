package log

import (
	"log"
	"os"
	"path/filepath"
)

const logDir = "logs"

// Init init logs directory
func Init(name string) {
	filename := filepath.Join(logDir, name)
	_ := os.Mkdir(logDir, os.ModePerm)
	logFile, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Println("can not create file:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("success create log file,", logFile.Name())
}
