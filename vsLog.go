/*vsLog : Very Simple Log package*/
/*Started: 18/03/2015*/
/*By: GregDuhamel <gregory.duhamel@outlook.com>*/

package vsLog

import (
	"io"
	"log"
	"os"
	"time"
)

type Log struct {
	Info           *log.Logger
	Warning        *log.Logger
	Error          *log.Logger
	FileDescriptor []byte
	LogDir         string
	LogFileName    string
	Output         string
	Flags          string
	LogRotate      time.Duration
	FlushOutput    time.Duration
}

func NewLog(outPut string, flags string) *Log {
	var infoHandle io.Writer = os.Stdout
	var warningHandle io.Writer = os.Stdout
	var errorHandle io.Writer = os.Stderr
	l := new(Log)

	l.Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)

	l.Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return (l)
}
