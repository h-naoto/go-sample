// log_train.go
package main

import (
	"fmt"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	loadConfig()
	log.Trace("TraceLog")
	log.Debug("DebugLog")
	log.Info("InfoLog")
	log.Warn("WarnLog")
	log.Error("ErrorLog")

}
func loadConfig() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		fmt.Println(err)
		panic("fail to load config")
	}
	log.ReplaceLogger(logger)
}
