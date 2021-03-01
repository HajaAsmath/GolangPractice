package Logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type hlogger struct {
	*log.Logger
	filename string
}

var once sync.Once
var hlog *hlogger

//return and instance if hydralogger
func GetLoggerInstance() *hlogger {
	fmt.Println("Creating Logger Instance")
	once.Do(func() {
		hlog = createLogger("hlogger.log")
	})
	return hlog
}

func createLogger(filename string) *hlogger {
	file, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	return &hlogger{
		Logger:   log.New(file, "Hydra: ", 8778),
		filename: filename,
	}
}
