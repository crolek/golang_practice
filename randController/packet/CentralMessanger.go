package packet

import (
	//"com.chuckrolek/randController"
	"fmt"
	"math/rand"
	//"ChannelMaker"
	//"com.chuckrolek/randController/app/"
	"strconv"
	//"time"
)

type DataMessage struct {
	channelMax    int
	channelNumber int
	randomNumber  int
}

type Totals struct {
	grandTotal    int
	channel1Total int
	channel2Total int
}

var (
	r2          = rand.New(rand.NewSource(10))
	mainChannel = make(chan DataMessage)
	seedMaker   = rand.New(rand.NewSource(99))
	totals      = new(Totals)
)

func Run() {
	totals.initTotals()
	var test1 = new(MessageSender)
	var test2 = new(MessageSender)
	go test1.sendMessage(1, getRandomSeedStarter(), 5)
	go test2.sendMessage(2, getRandomSeedStarter(), 100)

	fmt.Println("output stuff:")
	processChannelInfo()
}

func processChannelInfo() {
	var currentChannelTotal = 0
	for info := range mainChannel {
		if info.channelNumber == 1 {
			totals.channel1Total += info.randomNumber
			currentChannelTotal = totals.channel1Total
		} else if info.channelNumber == 2 {
			totals.channel2Total += info.randomNumber
			currentChannelTotal = totals.channel2Total
		}
		totals.grandTotal += info.randomNumber

		fmt.Printf("---- Channel %v Size ---- ---- Random ---- ---- Total ---- ---- Grand Total ----\n", strconv.Itoa(info.channelNumber))
		fmt.Printf("          %v                     %v               %v                 %v \n", strconv.Itoa(info.channelMax), strconv.Itoa(info.randomNumber), strconv.Itoa(currentChannelTotal), strconv.Itoa(totals.grandTotal))
	}
}

func getRandomSeedStarter() int64 {
	return seedMaker.Int63n(100)
}

func (t *Totals) initTotals() {
	totals.grandTotal = 0
	totals.channel1Total = 0
	totals.channel2Total = 0
}

/*

	//go sendMessage(1, getRandomSeedStarter(), 20)
	//go sendMessage(2, getRandomSeedStarter(), 5)
	//go sendChannel1()
	//go sendChannel2()


func sendChannel1() {
	for {
		time.Sleep(200 * time.Millisecond)
		mainChannel <- DataMessage{1, getRandomInt()}
	}
}

func sendChannel2() {
	for {
		time.Sleep(700 * time.Millisecond)
		mainChannel <- DataMessage{2, getRandomInt()}
	}
}
*/
