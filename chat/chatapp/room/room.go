package room

import(
	"time"
)

type RoomManager interface {
	AddUser()
	RemoveUser()
}

type Broadcast struct {
	incoming chan string
	outgoing chan string
}

type ChatRoom struct {
	id        string
	name      string
	owner     string
	dispatch  Broadcast
	users     []*User
	migration chan string
}
			 
type Message struct {
	sender    string
	room      string
	time      time.Time
	content   string
}




