package grifts

import (
	"github.com/gobuffalo/bloggy/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
