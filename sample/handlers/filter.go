package handlers

import "github.com/simbory/mego"

func testFilter(ctx *mego.HttpCtx) interface{} {
	data := ctx.GetItem("user")
	if data != nil {
		return data.(string)
	}
	return nil
}