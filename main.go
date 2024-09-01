package main

import (
	"daily/cmd"
	"log"
)

func main() {
	cmd.Execute()
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
}
