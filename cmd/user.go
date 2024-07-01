package cmd

import (
	"fmt"
	"net"
	"sync"
)

type user struct {
	Id   string
	Conn net.Conn
}

type Users struct {
	Mutext sync.RWMutex
	Users  map[string]user
}

func UsersInit() *Users {
	users := Users{}
	users.Users = make(map[string]user)
	return &users
}

func (u *Users) Create(conn net.Conn, id string) *user {
	u.Mutext.Lock()
	defer u.Mutext.Unlock()

	user := user{Id: fmt.Sprintf("#%s", id), Conn: conn}
	u.Users[conn.RemoteAddr().String()] = user
	return &user
}

func (u *Users) Get(id string) user {
	u.Mutext.RLock()
	defer u.Mutext.RUnlock()
	return u.Users[id]
}

func (u *Users) Delete(id string) {
	u.Mutext.Lock()
	defer u.Mutext.Unlock()
	delete(u.Users, id)
}

func (u *Users) AvailableUserNames() []string {
	u.Mutext.Lock()
	defer u.Mutext.Unlock()
	names := []string{}
	for _, user := range u.Users {
		names = append(names, user.Id)
	}
	return names
}

func (u *Users) GetUserByConn(conn net.Conn) user {
	u.Mutext.RLock()
	defer u.Mutext.RUnlock()
	return u.Users[conn.RemoteAddr().String()]
}

func (u *Users) BroadcastMessage(conn net.Conn, message string) {
	for _, user := range u.Users {
		if conn != user.Conn {
			go u.SendMessage(user.Conn, message)
		}
	}
}

func (u *Users) SendMessage(conn net.Conn, message string) {
	conn.Write([]byte(message))
}
