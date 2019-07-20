package main

import (
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"image/color"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
)

const (
	screenWidth  = 500
	screenHeight = 500
	actualheight = 600
	dpi          = 72
	fontSize     = 32
)

var (
	player   Snake
	apple    Apple
	gamefont font.Face
	counter  int
    paused   bool
)

func random(min, max float64) float64 {
	return (rand.Float64() * max) + min

}

func Game(screen *ebiten.Image) error {


    
    if ebiten.IsKeyPressed(ebiten.KeyP)  {
        paused=!paused
    }
  if !paused{	

    if !player.lose {
		ebiten.SetWindowTitle("Snake")
		screen.Fill(color.RGBA{0, 0, 0, 255})
		player.update()
		apple.update(&player)
		apple.draw(screen)
		player.draw(screen)
		text.Draw(screen, strconv.Itoa(player.score), gamefont, 0, screenHeight+50, color.White)
	} else {
		counter++
		screen.Fill(color.RGBA{0, 0, 0, 255})
		text.Draw(screen, strconv.Itoa(player.score), gamefont, screenWidth/2, screenHeight/2, color.White)
		if counter > 5*15 {
			
            player = NewSnake()
            apple = Apple{NewPoint(math.Round(random(0, screenWidth)), math.Round(random(0, screenHeight)))}
    
		}
	}
}else{
text.Draw(screen,"paused", gamefont, screenWidth/3, screenHeight/2, color.White)


}
	return nil

}

func main() {
	counter = 0
	f, err := ebitenutil.OpenFile("SourceCodePro-SemiBold.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetMaxTPS(13)
	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	gamefont = truetype.NewFace(tt, &truetype.Options{
		Size:    fontSize,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	player = NewSnake()
	apple = Apple{NewPoint(math.Round(random(0, screenWidth)), math.Round(random(0, screenHeight)))}
	if err := ebiten.Run(Game, screenWidth, actualheight, 1, "Snake"); err != nil {
		log.Fatal(err)
	}

}
