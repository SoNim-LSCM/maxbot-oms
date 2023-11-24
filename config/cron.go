package config

import (
	"errors"
	"os"
	"time"

	errorHandler "github.com/SoNim-LSCM/maxbot_oms/errors"
	"github.com/robfig/cron/v3"
)

func SetupLogCron(f *os.File) {
	logPath := os.Getenv("LOG_PATH")
	now := time.Now()
	err := errors.New("")
	f, err = os.OpenFile(logPath+"/TKOH-OMS-LOGS-"+now.Format("2006-01-02"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorHandler.CheckFatalError(err)
	c1 := cron.New()
	c1.AddFunc("0 0 * * *", func() {
		// fmt.Println("test")
		f, err = os.OpenFile(logPath+"/TKOH-OMS-LOGS-"+now.Format("2006-01-02"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		errorHandler.CheckFatalError(err)
	})
	c1.Start()

	for {
		time.Sleep(time.Second)
	}
}
