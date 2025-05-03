package main

import (
	"errtrack/cmd"
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(cmd.Execute())
}
