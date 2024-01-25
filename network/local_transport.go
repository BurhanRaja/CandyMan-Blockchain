package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}

func NewLocalTransport(addr NetAddr) *LocalTransport {
	return &LocalTransport{
		addr:      addr,
		consumeCh: make(chan RPC),
		peers:     make(map[NetAddr]*LocalTransport),
	}
}

func (lt *LocalTransport) Consume() chan RPC {
	return lt.consumeCh
}

func (lt *LocalTransport) Connect(ltr *LocalTransport) error {
	lt.lock.Lock()         // To prevent a third party from adding a peer or modify while we're connecting.
	defer lt.lock.Unlock() // Unlocks after the function exits

	lt.peers[ltr.addr] = ltr // Adding peers to communicate with.
	return nil
}

func (lt *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
	lt.lock.RLock()
	defer lt.lock.RUnlock()

	peer, ok := lt.peers[to]

	if !ok {
		return fmt.Errorf("%s: could not send a message to %s", lt.addr, to)
	}

	peer.consumeCh <- RPC{
		From:    lt.addr,
		Payload: payload,
	}

	return nil
}

func (t *LocalTransport) Addr() NetAddr {
	return t.addr
}