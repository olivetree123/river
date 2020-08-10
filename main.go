package main

import (
	"github.com/olivetree123/coco"
	"github.com/olivetree123/river/handlers"
	"github.com/olivetree123/river/pocket"
)

func Persistence() {
	for {
		// TODO: 如何持久化数据
		pocket.DataList.Persistence()
	}
}

func ClearGarbage() {
	for {
		// TODO: 需要清理持久化的数据
		pocket.GarbageList.Pop()
	}
}

func main() {
	c := coco.NewCoco()
	go Persistence()
	go ClearGarbage()
	c.AddRouter("post", "/push", handlers.PushHandler)
	c.AddRouter("post", "/pop", handlers.PopHandler)
	c.AddRouter("post", "/ack", handlers.AckHandler)
	if err := c.Run("0.0.0.0", 5000); err != nil {
		panic(err)
	}
}
