package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"math/rand"
)

type Room struct {
	name      int
	doorNorth *Room
	doorSouth *Room
	doorEast  *Room
	doorWest  *Room
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

func (r *Room) CreateRooms(theRoom *Room, depth int) {

	depth--

	if depth == 0 {
		return
	}
	if theRoom == nil {
		return
	}

	chooseDoor := rand.Intn(3)
	switch chooseDoor {
	case 0:
		if theRoom.doorNorth == nil {
			theRoom.SetNorthDoor(&Room{name: chooseDoor + rand.Intn(100), doorSouth: theRoom})
			theRoom.CreateRooms(theRoom.doorNorth, depth)
			return
		}
	case 1:
		if theRoom.doorEast == nil {
			theRoom.SetEastDoor(&Room{name: chooseDoor + rand.Intn(100), doorWest: theRoom})
			theRoom.CreateRooms(theRoom.doorEast, depth)
			return

		}
	case 2:
		if theRoom.doorSouth == nil {
			theRoom.SetSouthDoor(&Room{name: chooseDoor + rand.Intn(100), doorNorth: theRoom})
			theRoom.CreateRooms(theRoom.doorSouth, depth)
			return
		}
	case 3:
		if theRoom.doorWest == nil {
			theRoom.SetWestDoor(&Room{name: chooseDoor + rand.Intn(100), doorEast: theRoom})
			theRoom.CreateRooms(theRoom.doorWest, depth)
			return
		}
	}

}

func move(p *Player, direction int) {

	fmt.Printf("%s is in Room: %d\n", p.Name, p.CurrentRoom.name)
	switch direction {
	case 0:
		fmt.Printf("%s is checking to see if he can move to the North\n", p.Name)
		if p.CurrentRoom.doorNorth != nil {
			p.CurrentRoom = p.CurrentRoom.doorNorth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		}
		fmt.Printf("%s didn't move North because there isn't a door!\n", p.Name)
		return
	case 1:
		fmt.Printf("%s is checking to see if he can move to the East\n", p.Name)
		if p.CurrentRoom.doorEast != nil {
			p.CurrentRoom = p.CurrentRoom.doorEast
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		}
		fmt.Printf("%s didn't move East because there isn't a door!\n", p.Name)

		return
	case 2:
		fmt.Printf("%s is checking to see if he can move to the South\n", p.Name)
		if p.CurrentRoom.doorSouth != nil {
			p.CurrentRoom = p.CurrentRoom.doorSouth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		}
		fmt.Printf("%s didn't move South because there isn't a door!\n", p.Name)
		return
	case 3:
		fmt.Printf("%s is checking to see if he can move to the West\n", p.Name)
		if p.CurrentRoom.doorWest != nil {
			p.CurrentRoom = p.CurrentRoom.doorWest
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		}
		fmt.Printf("%s didn't move West because there isn't a door!\n", p.Name)
		return
	}

	fmt.Printf("no room for %s to move to\n", p.Name)
}

type Player struct {
	Name        string
	CurrentRoom *Room
}

func main() {

	dungeon := &Room{name: 0}

	dungeon.CreateRooms(dungeon, 10)
	spew.Dump(dungeon)

	p := &Player{
		Name:        "Zach",
		CurrentRoom: dungeon,
	}

	for i := 0; i <= 20; i++ {
		move(p, rand.Intn(3))
	}
}
