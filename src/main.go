package main

import (
	"log"

	"github.com/Rgss/imp-boot/src/server/starter"
)

func main() {

	defer end()

	newApplication();

	newHttpServer()
}

func end() {
	log.Printf("quit.")
}

// new application
func newApplication() {
	starter.StartApplication()
}

// http server
// 主要用于web 配置
func newHttpServer() {
	starter.StartHttpServer()
}

