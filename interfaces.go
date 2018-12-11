//=============================================================
// interfaces.go
//-------------------------------------------------------------
// Interfaces
//=============================================================
package main

import (
	"github.com/faiface/pixel"
)

type Entity interface {
	hit(x, y, vx, vy float64, power int) bool
	explode()
	getMass() float64
	getType() entityType
	draw(dt, elapsed float64)
	move(x, y float64)
	getPosition() pixel.Vec
	setPosition(x, y float64)
	getBounds() *Bounds
}
