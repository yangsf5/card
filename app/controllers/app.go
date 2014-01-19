package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) SignIn(user string) revel.Result {
	c.Validation.Required(user)

	if c.Validation.HasErrors() {
		c.Flash.Error("Please choose a nick name.")
		return c.Redirect(App.Index)
	}

	return c.Redirect("/hall/hall?user=%s", user)
}
