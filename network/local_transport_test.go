package network

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestConnection(t *testing.T) {
// 	lt := NewLocalTransport("localhost:8080")
// 	lt2 := NewLocalTransport("localhost:8081")
// 	lt3 := NewLocalTransport("localhost:8082")

// 	lt.Connect(lt2)
// 	lt2.Connect(lt3)
// 	lt3.Connect(lt)

// 	// fmt.Println(lt.peers[lt2.addr])
// 	// fmt.Println(lt2.peers)
// 	// fmt.Println(lt3.peers)

// 	assert.Equal(t, lt.peers[lt2.addr], lt2)
// 	assert.Equal(t, lt2.peers[lt3.addr], lt3)
// 	assert.Equal(t, lt3.peers[lt.addr], lt)
// }

func TestSendMessage(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)

	msg := []byte("hello world")
	tra.SendMessage(trb.addr, msg)

	rpc := <-trb.Consume()

	assert.Equal(t, rpc.Payload, msg)
	assert.Equal(t, rpc.From, tra.addr)
}
