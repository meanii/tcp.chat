package pkg

import (
	"fmt"
	"net"
	"sync"
)

type Room struct {
	Id      string
	Group   bool
	Private bool
	Users   []*User
}

type Rooms struct {
	mutext     sync.RWMutex
	Rooms      map[string]*Room
	Identities map[string]string
}

func InitRoomsInstance() *Rooms {
	rooms := Rooms{}
	rooms.Rooms = make(map[string]*Room)
	rooms.Identities = make(map[string]string)
	return &rooms
}

func (r *Rooms) CreateGroup(id string) *Room {
	r.mutext.Lock()
	defer r.mutext.Unlock()

	roomId := fmt.Sprintf("#%s", id)
	room := &Room{Id: roomId, Group: true, Private: false}
	r.Rooms[roomId] = room
	return room
}

func (r *Rooms) JoinGroup(id string, user User) *Room {
	r.mutext.Lock()
	defer r.mutext.Unlock()

	room, ok := r.Rooms[id]
	if !ok {
		return &Room{}
	}

	// Add user to the room
	room.Users = append(room.Users, &user)
	r.Identities[user.Conn.RemoteAddr().String()] = id
	return room
}

func (r *Rooms) GetRoomByUser(user *User) *Room {
	r.mutext.RLock()
	defer r.mutext.RUnlock()

	roomId, ok := r.Identities[user.Conn.RemoteAddr().String()]
	if !ok {
		return &Room{}
	}
	room := r.Rooms[roomId]
	return room
}

func (r *Room) RoomBroadcastNotification(conn net.Conn, message string, messageFun func(conn net.Conn, userId string, message string)) {
	for _, user := range r.Users {
		if user.Conn != conn {
			messageFun(user.Conn, r.Id, message)
		}
	}
}

func (r *Room) RoomBroadcast(conn net.Conn, groupId string, userId string, message string, messageFun func(conn net.Conn, groupId string, userId string, message string)) {
	for _, user := range r.Users {
		if user.Conn != conn {
			messageFun(user.Conn, groupId, userId, message)
		}
	}
}
