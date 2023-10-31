package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Room struct {
	name      int
	doorNorth *Room
	doorSouth *Room
	doorEast  *Room
	doorWest  *Room
	monster   *Monster
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

func (r *Room) CreateRooms(theRoom *Room, depth int, mnstrs *Monsters) {
	depth--

	if depth <= 0 {
		fmt.Println("DEPTH IS 0")
		return
	}
	//if theRoom == nil {
	//	return
	//}

	m := &Monster{Name: fmt.Sprintf("monster%d", rand.Intn(100))}
	mnstrs.addMonster(m)

	chooseDoor := rand.Intn(3)
	switch chooseDoor {
	case 0:
		if theRoom.doorNorth == nil {
			theRoom.SetNorthDoor(&Room{name: chooseDoor + rand.Intn(100), doorSouth: theRoom, monster: m})
			theRoom.CreateRooms(theRoom.doorNorth, depth, mnstrs)
		} else {
			depth++
			theRoom.CreateRooms(theRoom, depth, mnstrs)
		}
	case 1:
		if theRoom.doorEast == nil {
			theRoom.SetEastDoor(&Room{name: chooseDoor + rand.Intn(100), doorWest: theRoom, monster: m})
			theRoom.CreateRooms(theRoom.doorEast, depth, mnstrs)
		} else {
			depth++
			theRoom.CreateRooms(theRoom, depth, mnstrs)
		}
	case 2:
		if theRoom.doorSouth == nil {
			theRoom.SetSouthDoor(&Room{name: chooseDoor + rand.Intn(100), doorNorth: theRoom, monster: m})
			theRoom.CreateRooms(theRoom.doorSouth, depth, mnstrs)
		} else {
			depth++
			theRoom.CreateRooms(theRoom, depth, mnstrs)
		}
	case 3:
		if theRoom.doorWest == nil {
			theRoom.SetWestDoor(&Room{name: chooseDoor + rand.Intn(100), doorEast: theRoom, monster: m})
			theRoom.CreateRooms(theRoom.doorWest, depth, mnstrs)
		} else {
			depth++
			theRoom.CreateRooms(theRoom, depth, mnstrs)
		}
	}

}

type Monsters struct {
	Monsters []*Monster
	Lock     sync.Mutex
}

type Monster struct {
	Name string
}

func (m *Monsters) addMonster(mnstr *Monster) {
	m.Lock.Lock()
	m.Monsters = append(m.Monsters, mnstr)
	m.Lock.Unlock()
}

func (m *Monsters) removeMonster() *Monster {
	return &Monster{}
}

func (m *Monsters) createMonsters(monsterAmount int) {
	for i := 0; i <= monsterAmount; i++ {
		mnst := Monster{Name: fmt.Sprintf("monster%d", i)}
		m.Monsters = append(m.Monsters, &mnst)
	}
}

func (m *Monsters) listMonsters() {
	fmt.Println(m.Monsters)
}

func move(p *Player, direction int) {

	fmt.Printf("%s is in Room: %d\n", p.Name, p.CurrentRoom.name)
	switch direction {
	case 0:
		fmt.Printf("%s is checking to see if he can move to the North\n", p.Name)
		if p.CurrentRoom.doorNorth != nil {
			p.CurrentRoom = p.CurrentRoom.doorNorth
			fmt.Printf("%s moved to Room: %d and the monster in there is %s\n", p.Name, p.CurrentRoom.name, p.CurrentRoom.monster)
		}
		fmt.Printf("%s didn't move North because there isn't a door!\n", p.Name)
		return
	case 1:
		fmt.Printf("%s is checking to see if he can move to the East\n", p.Name)
		if p.CurrentRoom.doorEast != nil {
			p.CurrentRoom = p.CurrentRoom.doorEast
			fmt.Printf("%s moved to Room: %d and the monster in there is %s\n", p.Name, p.CurrentRoom.name, p.CurrentRoom.monster)
		}
		fmt.Printf("%s didn't move East because there isn't a door!\n", p.Name)
		return
	case 2:
		fmt.Printf("%s is checking to see if he can move to the South\n", p.Name)
		if p.CurrentRoom.doorSouth != nil {
			p.CurrentRoom = p.CurrentRoom.doorSouth
			fmt.Printf("%s moved to Room: %d and the monster in there is %s\n", p.Name, p.CurrentRoom.name, p.CurrentRoom.monster)
		}
		fmt.Printf("%s didn't move South because there isn't a door!\n", p.Name)
		return
	case 3:
		fmt.Printf("%s is checking to see if he can move to the West\n", p.Name)
		if p.CurrentRoom.doorWest != nil {
			p.CurrentRoom = p.CurrentRoom.doorWest
			fmt.Printf("%s moved to Room: %d and the monster in there is %s\n", p.Name, p.CurrentRoom.name, p.CurrentRoom.monster)
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
	for i := 0; i < 100; i++ {
		dungeon := &Room{name: i}
		mnstrs := &Monsters{}
		dungeon.CreateRooms(dungeon, 200, mnstrs)
	}
	//spew.Dump(dungeon)

	//p := &Player{
	//	Name:        "Zach",
	//	CurrentRoom: dungeon,
	//}
	//
	//for i := 0; i <= 20; i++ {
	//	move(p, rand.Intn(3))
	//}

}
