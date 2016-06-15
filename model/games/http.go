package games

import "appengine/srv"

// Templates
var newTmpl = "static/html/newGame.html"

func NewGame(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {

	wr.Parse()

	wr.R.ParseForm()

	//tc["Content"] = points
	return newTmpl, nil
}
