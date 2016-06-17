package games

import (
	"appengine/data"
	"appengine/srv"

	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"
)

type Game struct {
	Id        int64     `json:",string" datastore:"-"`
	Url       string    `json:"`
	UserName  string    `json:"`
	TimeStamp time.Time `json:"`
	Title     string
	PGNText   string
}

func randString() string {
	size := 32
	rb := make([]byte, size)
	rand.Read(rb)

	rs := base64.URLEncoding.EncodeToString(rb)
	return rs
}

func New(title string, pgn string) *Game {
	g := new(Game)

	g.Title = title
	g.PGNText = pgn
	g.TimeStamp = time.Now()
	g.UserName = "Unkown"
	g.Url = randString()

	return g
}

func (n *Game) ID() int64 {
	return n.Id
}

func (n *Game) SetID(id int64) {
	n.Id = id
}

type GameBuffer []*Game

func NewGameBuffer() GameBuffer {
	return make([]*Game, 0)
}

func (v GameBuffer) At(i int) data.DataItem {
	return data.DataItem(v[i])
}

func (v GameBuffer) Set(i int, t data.DataItem) {
	v[i] = t.(*Game)
}

func (v GameBuffer) Len() int {
	return len(v)
}

func Save(wr srv.WrapperRequest, g *Game) error {
	q := data.NewConn(wr, "games")

	err := q.Put(g)
	if err != nil {
		return fmt.Errorf("save: %v", err)
	}

	return nil
}

func GetByUrl(wr srv.WrapperRequest, url string) (*Game, error) {
	gms := NewGameBuffer()
	gm := new(Game)

	q := data.NewConn(wr, "games")
	q.AddFilter("Url =", url)
	q.GetMany(&gms)
	if len(gms) == 0 {
		return gm, fmt.Errorf("getbyurl: id game not found")
	}
	gm = gms[0]

	return gm, nil
}
