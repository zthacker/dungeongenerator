package dungeongenerator

import (
	"image"
	"math/rand"
	"sync"
	"time"
)

type Dungeon struct {
	*Room
	roomCounter int
	lock        sync.Mutex
	Rooms       []*Room
}

type Room struct {
	roomNumber int
	size       image.Rectangle
	doorNorth  *Room
	doorSouth  *Room
	doorEast   *Room
	doorWest   *Room
	//monster   *Monster
}

func NewDungeon() *Dungeon {
	return &Dungeon{}
}

func (r *Room) GetRoom() int {
	return r.roomNumber
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

func (r *Room) GetNorthDoor() *Room {
	if r.doorNorth != nil {
		return r.doorNorth
	}
	return nil
}

func (r *Room) GetSouthDoor() *Room {
	if r.doorSouth != nil {
		return r.doorSouth
	}
	return nil
}

func (r *Room) GetEastDoor() *Room {
	if r.doorEast != nil {
		return r.doorEast
	}
	return nil
}

func (r *Room) GetWestDoor() *Room {
	if r.doorWest != nil {
		return r.doorWest
	}
	return nil
}

func (d *Dungeon) CreateDungeon(rootRoom *Room, depth int) {
	r := rootRoom
	for i := 0; i < depth; i++ {
		r = d.createRoom(r)
	}
}

func (d *Dungeon) createRoom(room *Room) *Room {

	tempCreated := []*Room{}
	doors := []int{0, 1, 2, 3}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(doors), func(i, j int) { doors[i], doors[j] = doors[j], doors[i] })
	amountOfDoors := rand.Intn(4)
	if amountOfDoors == 0 {
		amountOfDoors = 2
	}
	for i := 0; i <= amountOfDoors; i++ {
		switch doors[i] {
		case 0:
			if room.doorNorth == nil {
				d.roomCounter++
				room.SetNorthDoor(&Room{roomNumber: d.roomCounter, doorSouth: room, size: image.Rect(room.size.Min.X+2, room.size.Min.Y+2, room.size.Max.X+2, room.size.Max.Y+2)})
				d.lock.Lock()
				tempCreated = append(tempCreated, room.doorNorth)
				d.Rooms = append(d.Rooms, room.doorNorth)
				d.lock.Unlock()
			} else {
				break
			}
		case 1:
			if room.doorEast == nil {
				d.roomCounter++
				room.SetEastDoor(&Room{roomNumber: d.roomCounter, doorWest: room})
				d.lock.Lock()
				tempCreated = append(tempCreated, room.doorEast)
				d.Rooms = append(d.Rooms, room.doorEast)
				d.lock.Unlock()
			} else {
				break
			}
		case 2:
			if room.doorSouth == nil {
				d.roomCounter++
				room.SetSouthDoor(&Room{roomNumber: d.roomCounter, doorNorth: room})
				d.lock.Lock()
				tempCreated = append(tempCreated, room.doorSouth)
				d.Rooms = append(d.Rooms, room.doorSouth)
				d.lock.Unlock()
			} else {
				break
			}
		case 3:
			if room.doorWest == nil {
				d.roomCounter++
				room.SetWestDoor(&Room{roomNumber: d.roomCounter, doorEast: room})
				d.lock.Lock()
				tempCreated = append(tempCreated, room.doorWest)
				d.Rooms = append(d.Rooms, room.doorWest)
				d.lock.Unlock()
			} else {
				break
			}
		}
	}

	randRoom := rand.Intn(len(tempCreated))
	return tempCreated[randRoom]
}
