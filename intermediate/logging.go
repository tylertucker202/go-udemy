package intermediate

import (
	"log"
	"os"
)

func main() {
	// log.Println("Hello, World!")

	// log.SetPrefix("INFO: ")
	// log.Println("This is an info message.")

	// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	// log.Println("This is a log message with date")

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	var (
		infoLogger    = log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		errorLogger   = log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
		warningLogger = log.New(logFile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	)

	infoLogger.Println("This is an info message using the infoLogger.")
	errorLogger.Println("This is an error message using the errorLogger.")
	warningLogger.Println("This is a warning message using the warningLogger.")

	debugLogger := log.New(logFile, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message written to the log file.")

}
