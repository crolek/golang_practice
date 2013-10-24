package packet

import (
	"fmt"
	//"math/rand"
	//"strconv"
	"flag"
)

var (
	numberOfMesangers int
)

func parseCommandLine() {
	flag.IntVar(&numberOfMesangers, "messangers", 4, "a help message")
	flag.Parse()
	fmt.Println(numberOfMesangers)
}
