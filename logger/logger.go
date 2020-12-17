package logger

import(
	"log"
	"lemon_ssh/util"
	"os"
)

var logger *log.Logger

func init(){
	logger = log.New(os.Stdout,"INFO",log.Ldate|log.Ltime)
}

func Error(format string,args ...interface{}){
	logger.SetPrefix(util.RedString("ERROR "))
	logger.Printf(format,args...)
}

func Debug(format string,args ...interface{}){
	logger.SetPrefix("DEBUG ")
	logger.Printf(format,args...)
}

func Warn(format string,args ...interface{}){
	logger.SetPrefix("WARN ")
	logger.Printf(format,args...)
}

func INFO(format string,args ...interface{}){
	logger.SetPrefix(util.GreenString("INFO "))
	logger.Printf(format,args...)
}
