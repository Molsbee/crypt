package gui

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type Hello struct {
	app.Compo
}

func (h *Hello) Render() app.UI {
	return app.Div().Body(app.H1().Body(app.Text("Hello")))
}
