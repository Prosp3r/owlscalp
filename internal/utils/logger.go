package utils

import "log"

func InitLogger() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Logger initialized - Owlscalp modular backtester ready")
}
