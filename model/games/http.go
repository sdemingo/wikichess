package games

import (
	"fmt"

	"appengine/srv"
)

// Templates
var newTmpl = "static/html/newGame.html"
var viewTmpl = "static/html/viewGame.html"
var errorTmpl = "static/html/error.html"

func NewGameHandler(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {

	wr.Parse()

	wr.R.ParseForm()

	tc["UserName"] = wr.NU.GetEmail()
	return newTmpl, nil
}

func SaveGameHandler(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {
	// if wr.NU.GetRole() < core.ROLE_ADMIN {
	// 	return infoTmpl, fmt.Errorf("points: addpoint: %s", core.ERR_NOTOPERATIONALLOWED)
	// }
	//wr.R.ParseForm()

	pgn := wr.R.PostFormValue("pgntext")
	title := wr.R.PostFormValue("title")

	srv.Log(wr, fmt.Sprintf("%s  %s", title, pgn))

	gm := New(title, pgn)
	gm.UserName = wr.NU.GetEmail()

	err := Save(wr, gm)
	if err != nil {
		return errorTmpl, fmt.Errorf("games: savegamehandler: %v", err)
	}

	tc["Content"] = gm

	return viewTmpl, nil
}
