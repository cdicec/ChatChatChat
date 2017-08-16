package room

import(
)

// type UserManager interface {
// 	ChangeRoom()
// 	ChangeNick()
// 	CreateRoom()
// 	DeleteRoom()
// }



type User struct {
	id	   string
	name   string
	role   string  // RoomOwner
	room   string  // RoomId
	active time.Time
	chan   chan string
}


