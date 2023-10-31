package dungeongenerator

import (
	"fmt"
	"math/rand"
)

type Room struct {
	name      int
	doorNorth *Room
	doorSouth *Room
	doorEast  *Room
	doorWest  *Room
	//monster   *Monster
}

func NewDungeon(rootRoomName int) *Room {
	return &Room{name: rootRoomName}
}

func (r *Room) GetRoom() int {
	return r.name
}

func (r *Room) SetNorthDoor(end *Room) {
	r.doorNorth = end
}

func (r *Room) SetSouthDoor(end *Room) {
	r.doorSouth = end
}

func (r *Room) SetEastDoor(end *Room) {
	r.doorEast = end
}

func (r *Room) SetWestDoor(end *Room) {
	r.doorWest = end
}

func (r *Room) CreateDungeon(room *Room, depth int) {
	depth--

	if depth <= 0 {
		fmt.Println("DEPTH IS 0")
		return
	}

	chooseDoor := rand.Intn(3)
	switch chooseDoor {
	case 0:
		if room.doorNorth == nil {
			room.SetNorthDoor(&Room{name: chooseDoor + rand.Intn(100), doorSouth: room})
			room.CreateDungeon(room.doorNorth, depth)
		} else {
			depth++
			room.CreateDungeon(room, depth)
		}
	case 1:
		if room.doorEast == nil {
			room.SetEastDoor(&Room{name: chooseDoor + rand.Intn(100), doorWest: room})
			room.CreateDungeon(room.doorEast, depth)
		} else {
			depth++
			room.CreateDungeon(room, depth)
		}
	case 2:
		if room.doorSouth == nil {
			room.SetSouthDoor(&Room{name: chooseDoor + rand.Intn(100), doorNorth: room})
			room.CreateDungeon(room.doorSouth, depth)
		} else {
			depth++
			room.CreateDungeon(room, depth)
		}
	case 3:
		if room.doorWest == nil {
			room.SetWestDoor(&Room{name: chooseDoor + rand.Intn(100), doorEast: room})
			room.CreateDungeon(room.doorWest, depth)
		} else {
			depth++
			room.CreateDungeon(room, depth)
		}
	}
}
