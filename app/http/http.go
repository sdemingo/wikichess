package http

import (
	"errors"
	"net/http"

	"app/core"

	"appengine/srv"
)

func init() {
	http.HandleFunc("/logout", logout)
}

func logout(w http.ResponseWriter, r *http.Request) {
	srv.RedirectUserLogin(w, r)
}

func RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	srv.RedirectUserLogin(w, r)
}

var baseTmpl = "static/html/start.html"
var helpTmpl = "static/html/help.html"
var adminTmpl = "static/html/admin.html"

func Help(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {
	return helpTmpl, nil
}

func Welcome(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {
	return baseTmpl, nil
}

func Admin(wr srv.WrapperRequest, tc map[string]interface{}) (string, error) {
	if wr.NU.GetRole() < core.ROLE_ADMIN {
		return "", errors.New(core.ERR_NOTOPERATIONALLOWED)
	}

	return adminTmpl, nil
}
