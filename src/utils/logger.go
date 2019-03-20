package infrastructure

import (
	"fmt"
	"log"
	"time"

	"github.com/vjeantet/jodaTime"
)

type LogLevel string

var (
	INFO    LogLevel = "INFO"
	WARNING LogLevel = "WARN"
	ERROR   LogLevel = "ERROR"
	FATAL   LogLevel = "FATAL"
)

func LogIt(level LogLevel, caller string, message string) {
	timestamp := jodaTime.Format("YYYY/MM/dd HH:mm:ss", time.Now())
	msg := fmt.Sprintf("[CRUDDY] - %s - %s - %s - %s", timestamp, level, caller, message)
	log.Println(msg)
}
