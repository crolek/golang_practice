package packet

import (
	//"fmt"
	"math/rand"
	//"strconv"
	"time"
)

type MessageSender struct {
	channelNumber    int
	randomNumber     int
	channelMaxNumber int
	randMaker        *rand.Rand
}

func (me *MessageSender) getRandomInt() int {
	return me.randMaker.Intn(me.channelMaxNumber)
}

func (me *MessageSender) sendMessage(channel int, seedNumber int64, max int) {
	me.channelMaxNumber = max
	me.randMaker = rand.New(rand.NewSource(seedNumber))
	for {

		time.Sleep(700 * time.Millisecond)
		mainChannel <- DataMessage{me.channelMaxNumber, channel, me.getRandomInt()}
	}
}

/*
func Hello() {
	fmt.Println("Hello")
}

func (t *Test) init() {
	randMaker = rand.New(rand.NewSource(0)) //to initalize the randMaker
}
*/
