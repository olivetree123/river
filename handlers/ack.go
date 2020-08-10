package handlers

import (
	"github.com/olivetree123/coco"
	"github.com/olivetree123/river/pocket"
)

func AckHandler(c *coco.Coco) coco.Result {
	msgID := c.Params.ByName("msgID")
	node := pocket.ConsumingPool.Load(msgID)
	pocket.GarbageList.Push(node)
	return coco.APIResponse(nil)
}