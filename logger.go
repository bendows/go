package logger

import (
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type errorWriter struct {
	io.Writer
	timeFormat string
}

type infoWriter struct {
	io.Writer
	timeFormat string
}

var LogOn = false

func (w errorWriter) Write(b []byte) (n int, err error) {
	if LogOn == true {
		w.Writer.Write([]byte("\033[1;33m"))
		w.Writer.Write([]byte(time.Now().Format(w.timeFormat)))
		w.Writer.Write([]byte("\033[0;37m"))
		stra := strings.Fields(string(b))
		line, remainder := stra[0], stra[1:]
		w.Writer.Write([]byte(line))
		stra2 := strings.Join(remainder, " ")
		stra2 += "\n"
		w.Writer.Write([]byte(" "))
		w.Writer.Write([]byte("\033[0;31m"))
		w.Writer.Write([]byte(stra2))
		return w.Writer.Write([]byte("\033[0m"))
	}
	return 0, io.EOF
}

func (w infoWriter) Write(b []byte) (n int, err error) {
	if LogOn == true {
		w.Writer.Write([]byte("\033[1;33m"))
		w.Writer.Write([]byte(time.Now().Format(w.timeFormat)))
		w.Writer.Write([]byte("\033[0;37m"))
		stra := strings.Fields(string(b))
		line, remainder := stra[0], stra[1:]
		w.Writer.Write([]byte(line))
		stra2 := strings.Join(remainder, " ")
		stra2 += "\n"
		w.Writer.Write([]byte(" "))
		w.Writer.Write([]byte("\033[0m"))
		w.Writer.Write([]byte(stra2))
		return w.Writer.Write([]byte("\033[0m"))
	}
	return 0, io.EOF
}

var Loginfo = log.New(&infoWriter{os.Stdout, "2006-01-02 15:04:05 "}, "", log.Lshortfile)
var Logerror = log.New(&errorWriter{os.Stdout, "2006-01-02 15:04:05 "}, "", log.Lshortfile)
