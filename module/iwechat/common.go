package iwechat

import (
	slogger "github.com/myafeier/logger"
	"log"
	"os"
)

var logger slogger.ILogger

func init() {
	logger = slogger.NewSimpleLogger2(os.Stdout, "[Module_wechat]", log.Lshortfile|log.Ldate|log.Lmicroseconds)

}
