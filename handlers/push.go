package handlers

import (
	"encoding/json"
	"github.com/olivetree123/coco"
	"github.com/olivetree123/river/pocket"
)

func PushHandler(c *coco.Coco) coco.Result {
	param := make(map[string]interface{})
	decoder := json.NewDecoder(c.Request.Body)
	if err := decoder.Decode(&param); err != nil {
		return coco.ErrorResponse(100)
	}
	pocket.DataList.Push(param)
	return coco.APIResponse("Hello Coco !")
}