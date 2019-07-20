package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"math"
)

const (
	width float64 = 25
)

type Snake struct {

	//body[0]ition of the Head

	vel   Point
	body  []Point
	score int
	lose  bool
}

func NewSnake() Snake {

	return Snake{NewPoint(0, width), []Point{NewPoint(screenWidth/2, 20), NewPoint(screenWidth/2, 20-width)}, 0, false}
}

func (s *Snake) update() {


	if ebiten.IsKeyPressed(ebiten.KeyW) && s.body[0].y > 0 {
		s.vel = NewPoint(0, -width)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) && (s.body[0].y+width) < screenHeight {
		s.vel = NewPoint(0, width)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) && s.body[0].x > 0 {
		s.vel = NewPoint(-width, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) && (s.body[0].x+width) < screenWidth {
		s.vel = NewPoint(width, 0)
	}

	// Move the body
	for i := len(s.body) - 1; i > 0; i-- {

		if s.body[i] == s.body[0] {
			s.lose = true
		}

		s.body[i] = s.body[i-1]

	}

	// Move the head

	s.body[0] = Add(s.body[0], s.vel)
	if s.body[0].x+width < 0 {
		s.body[0].x = screenWidth
		s.vel = NewPoint(-width, 0)
	} else if s.body[0].x > screenHeight {
		s.body[0].x = 0
		s.vel = NewPoint(width, 0)
	} else if s.body[0].y+width < 0 {
		s.body[0].y = screenHeight
		s.vel = NewPoint(0, -width)
	} else if s.body[0].y > screenHeight {
		s.body[0].y = 0
		s.vel = NewPoint(0, width)
	}

}

func (s *Snake) draw(screen *ebiten.Image) {
	for i := 0; i < len(s.body); i++ {
		ebitenutil.DrawRect(screen, s.body[i].x, s.body[i].y, width, width, color.RGBA{255, 0, 1, 255})

	}
}

type Apple struct {
	pos Point
}

func (a *Apple) draw(screen *ebiten.Image) {

	ebitenutil.DrawRect(screen, a.pos.x, a.pos.y, width, width, color.RGBA{255, 255, 253, 255})

}

func (a *Apple) update(s *Snake) {

	//Check for collision
	if (((a.pos.x <= s.body[0].x) && (s.body[0].x <= (a.pos.x + width))) && ((a.pos.y <= s.body[0].y) && (s.body[0].y <= (a.pos.y + width)))) || (((s.body[0].x <= a.pos.x) && (a.pos.x <= (s.body[0].x + width))) && ((s.body[0].y <= a.pos.y) && (a.pos.y <= (s.body[0].y + width)))) {

		a.pos = NewPoint(math.Round(random(0, screenWidth)), math.Round(random(0, screenHeight)))
		s.score = s.score + 1
		if len(s.body) == 0 {
			s.body = append(s.body, NewPoint(s.body[0].x+width, s.body[0].y))
		} else {
			s.body = append(s.body, NewPoint(s.body[len(s.body)-1].x+width, s.body[len(s.body)-1].y))
		}

	}
}
