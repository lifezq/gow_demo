package v1

import "github.com/lifezq/gow"

type Demo gow.Controller

func (c *Demo) HelloAction() {
	var set = struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Data    []string `json:"data"`
	}{
		Code:    200,
		Message: "ok",
		Data:    []string{"Hello", "World!"},
	}

	defer c.Response.RenderJson(&set)
}
