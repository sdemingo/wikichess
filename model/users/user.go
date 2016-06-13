package users

import (
	"fmt"
	"time"

	"app/core"

	"appengine/data"
	"appengine/srv"
)

const (
	MAXSZUSERNAME = 100

	ERR_NOTVALIDUSER   = "Usuario no valido"
	ERR_DUPLICATEDUSER = "Usuario duplicado"
	ERR_USERNOTFOUND   = "Usuario no encontrado"
)

type NUser struct {
	Id        int64 `json:",string" datastore:"-"`
	Mail      string
	Name      string
	Role      int8      `json:",string"`
	TimeStamp time.Time `json:"`
}

func (n *NUser) GetRole() int8 {
	return n.Role
}

func (n *NUser) GetEmail() string {
	return n.Mail
}

func (n *NUser) GetInfo() map[string]string {
	info := make(map[string]string)

	info["Username"] = n.Name
	if int(n.Role) < len(core.RoleNames) {
		info["RoleName"] = core.RoleNames[n.Role]
	} else {
		info["RoleName"] = ""
	}
	info["TimeStamp"] = n.TimeStamp.Format("02/01/2006")

	return info
}

func (n *NUser) ID() int64 {
	return n.Id
}

func (n *NUser) SetID(id int64) {
	n.Id = id
}

type NUserBuffer []*NUser

func NewNUserBuffer() NUserBuffer {
	return make([]*NUser, 0)
}

func (v NUserBuffer) At(i int) data.DataItem {
	return data.DataItem(v[i])
}

func (v NUserBuffer) Set(i int, t data.DataItem) {
	v[i] = t.(*NUser)
}

func (v NUserBuffer) Len() int {
	return len(v)
}

func putUser(wr srv.WrapperRequest, nu *NUser) error {

	nu.TimeStamp = time.Now()

	_, err := getUserByMail(wr, nu.Mail)
	if err == nil {
		return fmt.Errorf("putuser: %s", ERR_DUPLICATEDUSER)
	}

	q := data.NewConn(wr, "users")
	err = q.Put(nu)

	if err != nil {
		return fmt.Errorf("putuser: %v", err)
	}
	return nil
}

func updateUser(wr srv.WrapperRequest, nu *NUser) error {

	old, err := getUserById(wr, nu.Id)
	if err != nil {
		return fmt.Errorf("updateuser: %v", err)
	}

	// invariant fields
	nu.Mail = old.Mail
	nu.Id = old.Id
	nu.TimeStamp = old.TimeStamp

	q := data.NewConn(wr, "users")
	err = q.Put(nu)
	if err != nil {
		return fmt.Errorf("updateuser: %v", err)
	}
	return nil
}

func deleteUser(wr srv.WrapperRequest, nu *NUser) error {

	q := data.NewConn(wr, "users")
	err := q.Delete(nu)
	if err != nil {
		return fmt.Errorf("deleteuser: %v", err)
	}
	return nil
}

func getAllUsers(wr srv.WrapperRequest) ([]*NUser, error) {
	nus := NewNUserBuffer()

	q := data.NewConn(wr, "users")
	q.GetMany(&nus)

	return nus, nil
}

func getUserByMail(wr srv.WrapperRequest, email string) (*NUser, error) {
	nus := NewNUserBuffer()
	nu := new(NUser)

	q := data.NewConn(wr, "users")
	q.AddFilter("Mail =", email)
	q.GetMany(&nus)
	if len(nus) == 0 {
		return nu, fmt.Errorf("%s", ERR_USERNOTFOUND)
	}
	nu = nus[0]

	return nu, nil
}

func getUserById(wr srv.WrapperRequest, id int64) (*NUser, error) {
	nu := new(NUser)
	nu.Id = id
	q := data.NewConn(wr, "users")
	err := q.Get(nu)
	if err != nil {
		return nu, fmt.Errorf("getuserbyid: %v: %s", err, ERR_USERNOTFOUND)
	}

	return nu, nil
}
