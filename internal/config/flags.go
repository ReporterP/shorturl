package config

import (
    "flag"
)

var (
    FlagRunAddrAndPort string
    FlagRunBaseAddr    string
    FlagLogLevel       string
)

// parseFlags обрабатывает аргументы командной строки 
// и сохраняет их значения в соответствующих переменных
func ParseFlags() {
    flag.StringVar(&FlagRunAddrAndPort, "a", ":8080", "address and port to run server")
    flag.StringVar(&FlagRunBaseAddr, "b", "http://localhost:8080", "base address to run server")
    flag.StringVar(&FlagLogLevel, "l", "info", "log level")

    flag.Parse()
}