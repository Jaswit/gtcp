package client

import (
	"fmt"
	"github.com/Jaswit/gtcp/header"
	"log"
	"testing"
)


type S5Client struct {
	Client
}

func (c *S5Client) OnReceive(header header.IHeader, body []byte) error {
	log.Printf("my OnReceive")

	return fmt.Errorf("err")
}

func (c *S5Client) SendMsg(MsgID int32, data []byte) (int, error) {
	head := &header.HeadInfo{
		MsgLen: int32(len(data)),
		MsgID:  MsgID,
	}
	buff, err := head.Pack()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	buff = append(buff, data...)

	if n, err := c.Send(buff); err != nil {
		log.Println("Send Buff Data error:, ", err, " Conn Writer exit")
		return 0, err
	} else {
		return n, nil
	}
}

func TestClient(t *testing.T) {
	cli := &S5Client{}
	err := cli.Open("tcp", "127.0.0.1:9995")
	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < 5; i++ {
		_, err := cli.SendMsg(int32(i), []byte(fmt.Sprintf("helloword_:%d", i)))
		if err != nil {
			log.Println(err)
			return
		}
	}
}
