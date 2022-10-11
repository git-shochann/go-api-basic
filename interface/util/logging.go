// interface層

package util

import (
	"io"
	"log"
	"os"
)

type LoggingLogic interface {
	LoggingSetting()
}

type loggingLogic struct{}

func NewLoggingLogic() LoggingLogic {
	return &loggingLogic{}
}

// main()
// 1. logging := logic.NewLogging()
// 2. logging.LoggingSetting()

func (ll *loggingLogic) LoggingSetting() {
	file, err := os.OpenFile("logging.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile) // 出力のフォーマットを設定
	multiLogFile := io.MultiWriter(os.Stdout, file)      // 出力先を2つ設定
	log.SetOutput(multiLogFile)                          // 実際に設定を反映
}