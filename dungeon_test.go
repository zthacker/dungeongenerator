package dungeoncreator

import (
	"fmt"
	"math/rand"
	"testing"
)

type Player struct {
	Name        string
	CurrentRoom *Room
}

func TestDungeonCreate(t *testing.T) {
	dung := NewDungeon(0)

	dung.CreateDungeon(dung, 10)
}

func TestDungeonMove(t *testing.T) {
	dung := NewDungeon(0)

	dung.CreateDungeon(dung, 10)
	p := &Player{
		Name:        "zach",
		CurrentRoom: dung,
	}

	for i := 0; i < 20; i++ {
		moveDir := rand.Intn(3)
		move(p, moveDir)
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
		} else {
			fmt.Printf("%s didn't move North because there isn't a door!\n", p.Name)
		}
	case 1:
		fmt.Printf("%s is checking to see if he can move to the East\n", p.Name)
		if p.CurrentRoom.doorEast != nil {
			p.CurrentRoom = p.CurrentRoom.doorEast
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		} else {
			fmt.Printf("%s didn't move East because there isn't a door!\n", p.Name)
		}
	case 2:
		fmt.Printf("%s is checking to see if he can move to the South\n", p.Name)
		if p.CurrentRoom.doorSouth != nil {
			p.CurrentRoom = p.CurrentRoom.doorSouth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		} else {
			fmt.Printf("%s didn't move South because there isn't a door!\n", p.Name)
		}
	case 3:
		fmt.Printf("%s is checking to see if he can move to the West\n", p.Name)
		if p.CurrentRoom.doorWest != nil {
			p.CurrentRoom = p.CurrentRoom.doorWest
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.name)
		} else {
			fmt.Printf("%s didn't move West because there isn't a door!\n", p.Name)
		}
	}
}
