package packet

import (
	//"fmt"
	//"strconv"
	"testing"
)

var (
	messageSender1       = new(MessageSender)
	testSeedNumber int64 = 10
	testMax        int   = 10
)

func Test_GetRandomInt_1(t *testing.T) {
	messageSender1.init(10)
	var test = messageSender1.GetRandomInt()
	//fmt.Println(strconv.Itoa(messageSender1.GetRandomInt()))

	if test != 0 {
		t.Error("something")
	}

	t.Log("passed")
}
