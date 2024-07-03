package pkg

import (
	"fmt"
	"net"
	"sync"
)

type User struct {
	Id   string
	Conn net.Conn
}

type Users struct {
	mutext sync.RWMutex
	Users  map[string]User
}

func InitUsersInstance() *Users {
	users := Users{}
	users.Users = make(map[string]User)
	return &users
}

func (u *Users) Create(id string, conn net.Conn) *User {
	u.mutext.Lock()
	defer u.mutext.Unlock()

	user := User{Id: fmt.Sprintf("#%s", id), Conn: conn}
	u.Users[conn.RemoteAddr().String()] = user
	return &user
}

func (u *Users) GetUser(conn net.Conn) *User {
	u.mutext.RLock()
	defer u.mutext.RUnlock()

	user := u.Users[conn.RemoteAddr().String()]
	return &user
}
