package logger

import "log"

func Info(format string, v ...interface{}) {
	if v == nil {
		log.Println(format)
	} else{
		log.Printf(format, v)
	}
}

func Error(format string, v ...interface{}) {
	if v == nil {
		log.Println(format)
	} else{
		log.Printf(format, v)
	}
}