package games

import (
	"fmt"
	"strconv"

	"appengine/srv"
)

// Templates
var newTmpl = "static/html/newGame.html"
var viewTmpl = "static/html/viewGame.html"
var redirectTmpl = "static/html/redirect.html"
var errorTmpl = "static/html/error.html"

func NewGameHandler(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {

	wr.Parse()

	wr.R.ParseForm()

	tc["UserName"] = wr.NU.GetEmail()
	return newTmpl, nil
}

func SaveGameHandler(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {

	pgn := wr.R.PostFormValue("pgntext")
	title := wr.R.PostFormValue("title")

	if pgn == "" || title == "" {
		return errorTmpl, fmt.Errorf("games: savegamehandler: some fields are empty")
	}

	gm := New(title, pgn)
	gm.UserName = wr.NU.GetEmail()

	err := Save(wr, gm)
	if err != nil {
		return errorTmpl, fmt.Errorf("games: savegamehandler: %v", err)
	}

	tc["Content"] = gm

	return redirectTmpl, nil
}

func GetGameHandler(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {

	wr.Parse()
	sid := wr.Values.Get("id")
	id, err := strconv.ParseInt(sid, 10, 64)
	if sid == "" || err != nil {
		return errorTmpl, fmt.Errorf("games: getgamehandler: bad id")
	}

	game, err := GetById(wr, id)
	if err != nil {
		return errorTmpl, fmt.Errorf("games: getgamehandler: %v", err)
	}

	tc["Content"] = game

	return viewTmpl, nil
}
