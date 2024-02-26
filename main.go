package main

import (
    "fmt"

    console "github.com/asynkron/goconsole"
    "github.com/asynkron/protoactor-go/actor"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
    switch msg := context.Message().(type) {
    case Hello:
        fmt.Printf("Hello %v\n", msg.Who)
    }
}

func main() {
    system := actor.NewActorSystem()
    props := actor.PropsFromProducer(func() actor.Actor { return &HelloActor{} })
    pid := system.Root.Spawn(props)
    system.Root.Send(pid, Hello{Who: "Roger"})
    _ , _ = console.ReadLine()
}