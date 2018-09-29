/**
*	Author : Muhammad Arizal Saputro
*	StudentId: 1301178325
*/

package model

import "math"

type City struct {
	ID int
	X float64
	Y float64
}

func (c *City)SetLocation(x,y float64){
	c.X = x
	c.Y = y
}

func (c *City)GetX()float64 {
	return c.X
}

func (c *City)GetY()float64{
	return c.Y
}

func (c *City)DistanceTo(d City)float64{
	return math.Sqrt(math.Pow( c.X - d.X,2) + math.Pow(c.Y - d.Y,2))
}

