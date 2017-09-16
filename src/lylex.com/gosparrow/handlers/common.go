package handlers

import (
	"github.com/unrolled/render"
)

var gRender *render.Render

func init() {
	gRender = render.New(
		render.Options{
			IndentJSON: true,
		},
	)
}
