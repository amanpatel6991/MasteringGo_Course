package hydraLogger

import (
	"log"
	"sync"
	"os"
)

type hydraLogger struct {
	*log.Logger               //Allows to output msggs to writer interface
	fileName string
}

var hlogger *hydraLogger
var once sync.Once

//Get Instance of Logger (Give a singleton instance)
func GetInstance() *hydraLogger{
	once.Do(func() {
		hlogger = createLogger("hydraLogs.log")
	})
	return hlogger
}

func createLogger(fname string) *hydraLogger{
	file , _ := os.OpenFile(fname , os.O_RDWR|os.O_CREATE|os.O_TRUNC , 0777) //*file implements writer interface

	return &hydraLogger{
		fileName: fname,
		Logger : log.New(file , "Hydra " ,log.Lshortfile),               //initialize the embedded logger type
	}
}