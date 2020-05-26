package middelware

import (
	"log"
	"os"

	"github.com/labstack/echo/middleware"
)

//example the format
var format = "method=${method}, uri=${uri}, status=${status},latency:${latency}\n"

/*
Default Configuration:
Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
    `"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
    `"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
    `"bytes_out":${bytes_out}}` + "\n",
*/

//create file .log
func filelogger() *os.File {
	myLog, err := os.OpenFile(
		"logs.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666,
	)
	if err != nil {
		log.Fatalf("no cold open or create file logs")
	}
	//defer myLog.Close() //cerrar el archivo termine el proceso
	return myLog
}

//config the logger
var logConfig = middleware.LoggerConfig{
	Output: filelogger(),
	Format: format,
}
