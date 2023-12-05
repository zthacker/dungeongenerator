package dungeongenerator

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"image"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"testing"
)

type Player struct {
	Name        string
	CurrentRoom *Room
}

var rootRoom = &Room{
	roomNumber: 0,
	size:       image.Rect(0, 0, 4, 4),
	doorNorth:  nil,
	doorSouth:  nil,
	doorEast:   nil,
	doorWest:   nil,
}

func TestRoomCreate(t *testing.T) {
	d := NewDungeon()

	r := &Room{}

	d.CreateDungeon(r, 10)

	for _, r := range d.Rooms {
		//logrus.Infof("main room: %d", r.roomNumber)
		//if r.doorNorth != nil {
		//	logrus.Infof("north: %d", r.doorNorth.roomNumber)
		//}
		//if r.doorSouth != nil {
		//	logrus.Infof("south: %d", r.doorSouth.roomNumber)
		//
		//}
		//if r.doorEast != nil {
		//	logrus.Infof("east: %d", r.doorEast.roomNumber)
		//
		//}
		//if r.doorWest != nil {
		//	logrus.Infof("west: %d", r.doorWest.roomNumber)
		//}
		logrus.Infof("room: %d has a size of: %+v", r.roomNumber, r.size)
	}
}

func TestMapCreate(t *testing.T) {
	dungeonImage := image.Rect(0, 0, 500, 500)
	rootImage := image.NewRGBA(dungeonImage)

	d := NewDungeon()
	d.CreateDungeon(rootRoom, 2)

	draw.Draw(rootImage, rootRoom.size, image.White, rootRoom.size.Min, draw.Src)

	for _, r := range d.Rooms {
		logrus.Infof("main room: %d", r.roomNumber)
		if r.doorNorth != nil {
			logrus.Infof("north: %d", r.doorNorth.roomNumber)
			draw.Draw(rootImage, r.doorNorth.size, image.White, r.doorNorth.size.Min, draw.Src)
		}
		if r.doorSouth != nil {
			logrus.Infof("south: %d", r.doorSouth.roomNumber)
			draw.Draw(rootImage, r.doorSouth.size, image.White, r.doorSouth.size.Min, draw.Src)
		}
		if r.doorEast != nil {
			logrus.Infof("east: %d", r.doorEast.roomNumber)
			draw.Draw(rootImage, r.doorEast.size, image.White, r.doorEast.size.Min, draw.Src)
		}
		if r.doorWest != nil {
			logrus.Infof("west: %d", r.doorWest.roomNumber)
			draw.Draw(rootImage, r.doorWest.size, image.White, r.doorWest.size.Min, draw.Src)
		}
	}

	f, err := os.Create("dungeon.png")
	if err != nil {
		logrus.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, rootImage)
}

func TestDungeonMove(t *testing.T) {

	dung := NewDungeon()

	dung.CreateDungeon(rootRoom, 10)
	p := &Player{
		Name:        "zach",
		CurrentRoom: rootRoom,
	}

	for i := 0; i < 20; i++ {
		moveDir := rand.Intn(3)
		move(p, moveDir)
	}
}

func move(p *Player, direction int) {
	fmt.Printf("%s is in Room: %d\n", p.Name, p.CurrentRoom.roomNumber)
	switch direction {
	case 0:
		fmt.Printf("%s is checking to see if he can move to the North\n", p.Name)
		if p.CurrentRoom.GetNorthDoor() != nil {
			p.CurrentRoom = p.CurrentRoom.doorNorth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.roomNumber)
		} else {
			fmt.Printf("%s didn't move North because there isn't a door!\n", p.Name)
		}
	case 1:
		fmt.Printf("%s is checking to see if he can move to the East\n", p.Name)
		if p.CurrentRoom.GetEastDoor() != nil {
			p.CurrentRoom = p.CurrentRoom.doorEast
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.roomNumber)
		} else {
			fmt.Printf("%s didn't move East because there isn't a door!\n", p.Name)
		}
	case 2:
		fmt.Printf("%s is checking to see if he can move to the South\n", p.Name)
		if p.CurrentRoom.GetSouthDoor() != nil {
			p.CurrentRoom = p.CurrentRoom.doorSouth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.roomNumber)
		} else {
			fmt.Printf("%s didn't move South because there isn't a door!\n", p.Name)
		}
	case 3:
		fmt.Printf("%s is checking to see if he can move to the West\n", p.Name)
		if p.CurrentRoom.GetWestDoor() != nil {
			p.CurrentRoom = p.CurrentRoom.doorNorth
			fmt.Printf("%s moved to Room: %d\n", p.Name, p.CurrentRoom.roomNumber)
		} else {
			fmt.Printf("%s didn't move West because there isn't a door!\n", p.Name)
		}
	}
}

func TestDraw(t *testing.T) {
	r := image.Rect(0, 0, 500, 500)
	rImage := image.NewRGBA(r)

	rootR := image.Rect(100, 100, 104, 104)
	nRoom := image.Rect(100, 95, 104, 99)
	sRoom := image.Rect(100, 105, 104, 109)
	eRoom := image.Rect(105, 100, 109, 104)
	wRoom := image.Rect(95, 100, 99, 104)

	n2Room := image.Rect(100, 90, 104, 94)

	draw.Draw(rImage, rootR, image.White, rootR.Min, draw.Src)
	draw.Draw(rImage, nRoom, image.Black, nRoom.Min, draw.Src)
	draw.Draw(rImage, sRoom, image.Black, sRoom.Min, draw.Src)
	draw.Draw(rImage, eRoom, image.Black, eRoom.Min, draw.Src)
	draw.Draw(rImage, wRoom, image.Black, wRoom.Min, draw.Src)
	draw.Draw(rImage, n2Room, image.White, n2Room.Min, draw.Src)

	f, err := os.Create("dungeon.png")
	if err != nil {
		logrus.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, rImage)

}
