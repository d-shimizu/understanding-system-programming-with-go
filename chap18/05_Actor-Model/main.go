package main

import (
	"fmt"
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
)

// メッセージの構造体
type hello struct{ Who string }

// アクターの構造体
// actor は Receiveメソッドを持つ構造体
type helloActor struct{}

// アクターのメールボックス受信時に呼ばれるメソッド
func (state *helloActor) Receive(context actor.Context) {
	// 型アサーション (Type Assertion)
	switch msg := context.Message().(type) {
	case *hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	//props := actor.FromInstance(&helloActor{})
	props := actor.PropsFromProducer(func() actor.Actor { return &helloActor{} })
	pid := actor.Spawn(props)
	pid.Tell(&hello{Who: "Roger"})
	console.ReadLine()
}
