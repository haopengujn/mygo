package main

import (
	"fmt"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box

func (b Box) volume() float64 {
	return b.width * b.height * b.depth
}

func (b *Box) setClour(c Color) {
	b.color = c
}

func (bl BoxList) biggestClour() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range bl {
		if b.volume() > v {
			v = b.volume()
			k = b.color
		}
	}
	return k
}

func (bl BoxList) paintAllBlack() {
	for i, _ := range bl {
		bl[i].setClour(BLACK)
	}
	// for _, b := range bl {
	// 	b.setClour(BLACK)
	// }
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func main() {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 1, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}
	fmt.Printf("We have %d boxs in our set\n", len(boxes))
	fmt.Println("The volume of the frist one is ", boxes[0].volume())
	fmt.Println("The color of the last one is ", boxes[len(boxes)-1].color.String())
	fmt.Println("The biggest one is ", boxes.biggestClour().String())
	fmt.Println("Let's paint them all black")
	boxes.paintAllBlack()
	fmt.Println("The second one's color now is ", boxes[1].color.String())
	fmt.Println("Obviously, now, the biggest one's color is ", boxes.biggestClour().String())
}
