package main

import "fmt"

type Tile struct{}

type TileWalker interface {
	WalkTile(Tile)
}

type Updater interface {
	Update()
}

type TransForm struct {
	position int
}

type Enemy struct {
	TransForm
	tileWalker TileWalker
}

func (e *Enemy) checkTilesCollided() {
	// logic
	// collision checks
	fmt.Println("enemy walking on tile", e.position)
	e.tileWalker.WalkTile(Tile{})
}

func (e *Enemy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

func Update(u Updater) {
	u.Update()
}

type FireEnemy struct {
	*Enemy
}

func (f *FireEnemy) Update() {
	f.checkTilesCollided()
}

func (f *FireEnemy) WalkTile(tile Tile) {
	fmt.Println("fire enemy is walking on tile")
}

func main() {
	e := &FireEnemy{}
	e.Enemy = &Enemy{
		tileWalker: e,
	}
	for i := 0; i < 100; i++ {
		Update(e)
	}
}
