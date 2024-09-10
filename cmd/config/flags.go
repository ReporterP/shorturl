package config

import (
    "flag"
)

var FlagRunAddrAndPort string
var FlagRunBaseAddr string

// parseFlags обрабатывает аргументы командной строки 
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
    flag.StringVar(&FlagRunAddrAndPort, "a", ":8080", "address and port to run server")
    flag.StringVar(&FlagRunBaseAddr, "b", "http://localhost:8080", "base address to run server")
    // парсим переданные серверу аргументы в зарегистрированные переменные
    flag.Parse()
}