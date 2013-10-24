package packet

import (
	"fmt"
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

func (me MessageSender) init(max int) {
	me.channelMaxNumber = max
	fmt.Println("test")
}

func (me MessageSender) GetRandomInt() int {
	//	me.channelMaxNumber = 20
	//	fmt.Println("Max: " + strconv.Itoa(me.channelMaxNumber))
	return me.randMaker.Intn(me.channelMaxNumber)
}

func (me MessageSender) SendMessage(channel int, seedNumber int64, max int) {
	me.channelMaxNumber = max
	me.randMaker = rand.New(rand.NewSource(seedNumber))
	for {

		time.Sleep(700 * time.Millisecond)
		mainChannel <- DataMessage{me.channelMaxNumber, channel, me.GetRandomInt()}
	}
}
