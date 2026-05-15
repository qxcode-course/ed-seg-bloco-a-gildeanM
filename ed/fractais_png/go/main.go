package main

import (
	"fmt"
	"math/rand"
)

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func desenhaQuadrado(pen *Pen, size float64) {
	for i := 0; i < 4; i++ {
		pen.Walk(size)
		pen.Right(90)
	}
}

func tapete(pen *Pen, len float64) {

	// fator := 0.35
	// if dirr < 1 {
	// 	return
	// }

	// angulodir := float64(60)

	// pen.DrawCircle(dirr)
	// pen.Right(angulodir / 2)
	// for i := 0; i < 6; i++ {
	// 	pen.Right(angulodir)
	// 	pen.Up()
	// 	pen.Walk(dirr)
	// 	pen.Down()
	// 	arvore(pen, dirr*fator)
	// 	// pen.DrawCircle(dirr * fator)
	// 	pen.Up()
	// 	pen.Walk(-dirr)
	// }
	// pen.Left(angulodir / 2)

	var pos float64
	pos = 1200
	// parent_pos := pos / 3
	pos /= 3
	len /= 3
	pen.SetPosition(pos, pos)

	pen.FillSquare(len, len)
	var child_pos float64
	child_pos = pos / 9
	len /= 3
	pen.SetPosition(child_pos*2, pos)
	pen.FillSquare(len, len)

	pen.SetPosition(child_pos*5, pos)
	pen.FillSquare(len, len)

	pen.SetPosition(child_pos*8, pos)
	pen.FillSquare(len, len)

}

func main() {
	pen := NewPen(1200, 1200) // cria um canvas de 1000 de largura por 1000 de altura
	pen.SetRGB(0, 0, 0)       // muda a cor do pincel para vermelho
	pen.SetPosition(0, 0)     // move o pincel para x 250, y 500
	pen.SetHeading(90)

	pen.FillSquare(1200, 1200)
	pen.SetRGB(255, 255, 255)
	tapete(pen, 1200)

	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
