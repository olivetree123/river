package handlers

import (
	"github.com/olivetree123/coco"
	"github.com/olivetree123/river/pocket"
)

func PopHandler(c *coco.Coco) coco.Result {
	node := pocket.DataList.Pop()
	return coco.APIResponse(node)
}