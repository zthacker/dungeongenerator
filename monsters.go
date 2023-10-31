package dungeoncreator

import (
	"fmt"
	"sync"
)

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
