package main

import (
	"golib/log"
)

func main() {
	log.Info(" I'm info .")
	log.Debug(" I'm trace .")
	log.Warning(" I'm warning .")
	log.Error(" I'm Error .")
	log.Fatal(" I'm fatal .")
}
