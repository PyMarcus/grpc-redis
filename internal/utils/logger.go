package utils

import (
	"log"
	"os"

	"github.com/natefinch/lumberjack"
)

// Init: inicializa o log com rotatividade
func Init() {
	err := os.MkdirAll("./logs", 0755)
	if err != nil {
		log.Fatalf("Erro ao criar diretório de logs: %v", err)
	}

	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/grpc-redis.log",
		MaxSize:    50, // MB
		MaxBackups: 3,
		MaxAge:     2, // dias
		Compress:   true,
	})

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("Logit has been initialized!")
}
