# Plain Text Formatter for Logrus Log
A plain text formatter for logrus logger which is suitable for screen print

```
package main

import (
	logformatter "github.com/jeffwubj/plaintextformatter"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&logformatter.PlainTextFormatter{
		ShowLevel: true,
		ShowTime:  true,
	})
	log.Infoln("Info")
	log.Warnln("Warn")
	log.Errorln("Error")
	log.Debugln("Debug")
}
```

will display
```
INFO 2019-01-22T10:51:02.945215+08:00 Info
WARNING 2019-01-22T10:51:02.945372+08:00 Warn
ERROR 2019-01-22T10:51:02.945376+08:00 Error
DEBUG 2019-01-22T10:51:02.94538+08:00 Debug
```
