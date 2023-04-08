package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type Object struct {
	x, y       int
	shape      rune
	directionX int
	directionY int
}

func (obj *Object) draw() {
	termbox.SetCell(obj.x, obj.y, obj.shape, termbox.ColorDefault, termbox.ColorDefault)
}

func (obj *Object) clear() {
	termbox.SetCell(obj.x, obj.y, ' ', termbox.ColorDefault, termbox.ColorDefault)
}

func (obj *Object) move() {
	obj.clear()
	obj.x += obj.directionX
	obj.y += obj.directionY
	obj.draw()
}

func (obj *Object) collide(other *Object) bool {
	return obj.x == other.x && obj.y == other.y
}

func createObject(x, y int, shape rune, directionX, directionY int) *Object {
	return &Object{
		x:          x,
		y:          y,
		shape:      shape,
		directionX: directionX,
		directionY: directionY,
	}
}

func listenKeyboard(keyboardChan chan termbox.Event) {
	for {
		keyboardChan <- termbox.PollEvent()
	}
}

func initTermbox() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}

func main() {
	initTermbox()
	defer termbox.Close()

	player := createObject(20, 20, '^', 0, 0)
	bullets := []*Object{}
	enemies := []*Object{}

	keyboardChan := make(chan termbox.Event)
	go listenKeyboard(keyboardChan)

	// Move player based on keyboard input
	go func() {
		for {
			select {
			case ev := <-keyboardChan:
				switch ev.Key {
				case termbox.KeyArrowUp:
					player.directionY = -1
				case termbox.KeyArrowDown:
					player.directionY = 1
				case termbox.KeyArrowLeft:
					player.directionX = -1
				case termbox.KeyArrowRight:
					player.directionX = 1
				case termbox.KeySpace:
					bullets = append(bullets, createObject(player.x, player.y-1, '|', 0, -1))
				}
			case <-time.After(time.Millisecond * 50):
				player.directionX, player.directionY = 0, 0
			}
		}
	}()

	// Generate enemies at random locations
	go func() {
		for {
			enemies = append(enemies, createObject(rand.Intn(50), 2, 'V', 0, 1))
			time.Sleep(time.Second)
		}
	}()

loop:
	for {
		err := termbox.Flush()
		if err != nil {
			panic(err)
		}

		// Update bullets positions and remove out of bound bullets
		for i := 0; i < len(bullets); i++ {
			bullets[i].move()
			if bullets[i].y < 1 {
				bullets[i].clear()
				bullets = append(bullets[:i], bullets[i+1:]...)
				i--
			}
		}

		// Update enemies positions and remove out of bound enemies
		for i := 0; i < len(enemies); i++ {
			enemies[i].move()
			if enemies[i].y > 25 {
				enemies[i].clear()
				enemies = append(enemies[:i], enemies[i+1:]...)
				i--
			}
		}

		// Check for bullet-enemy collisions
		for i := 0; i < len(bullets); i++ {
			for j := 0; j < len(enemies); j++ {
				if bullets[i].collide(enemies[j]) {
					bullets[i].clear()
					bullets = append(bullets[:i], bullets[i+1:]...)
					i--
					enemies[j].clear()
					enemies = append(enemies[:j], enemies[j+1:]...)
					j--
				}
			}
		}

		// Check for player-enemy collisions
		for i := 0; i < len(enemies); i++ {
			if player.collide(enemies[i]) {
				println("Game Over")
				break loop
			}
		}

		player.move()
		time.Sleep(time.Millisecond * 50)
	}
}
