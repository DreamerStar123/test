package main

import "fmt"

type Face [3][3]string

type Cube struct {
	U, D, F, B, L, R Face
}

func rotateFaceCW(f *Face) {
	tmp := *f
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f[j][2-i] = tmp[i][j]
		}
	}
}

func rotateFaceCCW(f *Face) {
	tmp := *f
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			f[2-j][i] = tmp[i][j]
		}
	}
}

func (c *Cube) rotateU() {
	rotateFaceCW(&c.U)
	topF := c.F[0]
	topR := c.R[0]
	topB := c.B[0]
	topL := c.L[0]
	c.F[0] = topR
	c.R[0] = topB
	c.B[0] = topL
	c.L[0] = topF
}

func (c *Cube) rotateUPrime() {
	rotateFaceCCW(&c.U)
	topF := c.F[0]
	topR := c.R[0]
	topB := c.B[0]
	topL := c.L[0]
	c.F[0] = topL
	c.L[0] = topB
	c.B[0] = topR
	c.R[0] = topF
}

func (c *Cube) rotateD() {
	rotateFaceCW(&c.D)
	botF := c.F[2]
	botR := c.R[2]
	botB := c.B[2]
	botL := c.L[2]
	c.F[2] = botL
	c.L[2] = botB
	c.B[2] = botR
	c.R[2] = botF
}

func (c *Cube) rotateDPrime() {
	rotateFaceCCW(&c.D)
	botF := c.F[2]
	botR := c.R[2]
	botB := c.B[2]
	botL := c.L[2]
	c.F[2] = botR
	c.R[2] = botB
	c.B[2] = botL
	c.L[2] = botF
}

func (c *Cube) rotateF() {
	rotateFaceCW(&c.F)
	top := [3]string{c.U[2][0], c.U[2][1], c.U[2][2]}
	right := [3]string{c.R[0][0], c.R[1][0], c.R[2][0]}
	bottom := [3]string{c.D[0][0], c.D[0][1], c.D[0][2]}
	left := [3]string{c.L[2][2], c.L[1][2], c.L[0][2]}

	c.U[2][0], c.U[2][1], c.U[2][2] = left[2], left[1], left[0]
	c.R[0][0], c.R[1][0], c.R[2][0] = top[0], top[1], top[2]
	c.D[0][0], c.D[0][1], c.D[0][2] = right[2], right[1], right[0]
	c.L[0][2], c.L[1][2], c.L[2][2] = bottom[0], bottom[1], bottom[2]
}

func (c *Cube) rotateFPrime() {
	rotateFaceCCW(&c.F)
	top := [3]string{c.U[2][0], c.U[2][1], c.U[2][2]}
	right := [3]string{c.R[0][0], c.R[1][0], c.R[2][0]}
	bottom := [3]string{c.D[0][0], c.D[0][1], c.D[0][2]}
	left := [3]string{c.L[2][2], c.L[1][2], c.L[0][2]}

	c.U[2][0], c.U[2][1], c.U[2][2] = right[0], right[1], right[2]
	c.R[0][0], c.R[1][0], c.R[2][0] = bottom[2], bottom[1], bottom[0]
	c.D[0][0], c.D[0][1], c.D[0][2] = left[0], left[1], left[2]
	c.L[0][2], c.L[1][2], c.L[2][2] = top[2], top[1], top[0]
}

func (c *Cube) rotateB() {
	rotateFaceCW(&c.B)
	top := [3]string{c.U[0][2], c.U[0][1], c.U[0][0]}
	right := [3]string{c.R[0][2], c.R[1][2], c.R[2][2]}
	bottom := [3]string{c.D[2][0], c.D[2][1], c.D[2][2]}
	left := [3]string{c.L[2][0], c.L[1][0], c.L[0][0]}

	c.U[0][0], c.U[0][1], c.U[0][2] = left[2], left[1], left[0]
	c.R[0][2], c.R[1][2], c.R[2][2] = top[2], top[1], top[0]
	c.D[2][2], c.D[2][1], c.D[2][0] = right[2], right[1], right[0]
	c.L[0][0], c.L[1][0], c.L[2][0] = bottom[0], bottom[1], bottom[2]
}

func (c *Cube) rotateBPrime() {
	rotateFaceCCW(&c.B)
	top := [3]string{c.U[0][2], c.U[0][1], c.U[0][0]}
	right := [3]string{c.R[0][2], c.R[1][2], c.R[2][2]}
	bottom := [3]string{c.D[2][0], c.D[2][1], c.D[2][2]}
	left := [3]string{c.L[2][0], c.L[1][0], c.L[0][0]}

	c.U[0][0], c.U[0][1], c.U[0][2] = right[0], right[1], right[2]
	c.R[0][2], c.R[1][2], c.R[2][2] = bottom[2], bottom[1], bottom[0]
	c.D[2][2], c.D[2][1], c.D[2][0] = left[0], left[1], left[2]
	c.L[0][0], c.L[1][0], c.L[2][0] = top[2], top[1], top[0]
}

func (c *Cube) rotateL() {
	rotateFaceCW(&c.L)
	top := [3]string{c.U[0][0], c.U[1][0], c.U[2][0]}
	front := [3]string{c.F[0][0], c.F[1][0], c.F[2][0]}
	bottom := [3]string{c.D[0][0], c.D[1][0], c.D[2][0]}
	back := [3]string{c.B[2][2], c.B[1][2], c.B[0][2]}

	c.U[0][0], c.U[1][0], c.U[2][0] = back[0], back[1], back[2]
	c.F[0][0], c.F[1][0], c.F[2][0] = top[0], top[1], top[2]
	c.D[0][0], c.D[1][0], c.D[2][0] = front[0], front[1], front[2]
	c.B[0][2], c.B[1][2], c.B[2][2] = bottom[2], bottom[1], bottom[0]
}

func (c *Cube) rotateLPrime() {
	rotateFaceCCW(&c.L)
	top := [3]string{c.U[0][0], c.U[1][0], c.U[2][0]}
	front := [3]string{c.F[0][0], c.F[1][0], c.F[2][0]}
	bottom := [3]string{c.D[0][0], c.D[1][0], c.D[2][0]}
	back := [3]string{c.B[2][2], c.B[1][2], c.B[0][2]}

	c.U[0][0], c.U[1][0], c.U[2][0] = front[0], front[1], front[2]
	c.F[0][0], c.F[1][0], c.F[2][0] = bottom[0], bottom[1], bottom[2]
	c.D[0][0], c.D[1][0], c.D[2][0] = back[2], back[1], back[0]
	c.B[0][2], c.B[1][2], c.B[2][2] = top[2], top[1], top[0]
}

func (c *Cube) rotateR() {
	rotateFaceCW(&c.R)
	top := [3]string{c.U[0][2], c.U[1][2], c.U[2][2]}
	front := [3]string{c.F[0][2], c.F[1][2], c.F[2][2]}
	bottom := [3]string{c.D[0][2], c.D[1][2], c.D[2][2]}
	back := [3]string{c.B[2][0], c.B[1][0], c.B[0][0]}

	c.U[0][2], c.U[1][2], c.U[2][2] = front[0], front[1], front[2]
	c.F[0][2], c.F[1][2], c.F[2][2] = bottom[0], bottom[1], bottom[2]
	c.D[0][2], c.D[1][2], c.D[2][2] = back[2], back[1], back[0]
	c.B[0][0], c.B[1][0], c.B[2][0] = top[2], top[1], top[0]
}

func (c *Cube) rotateRPrime() {
	rotateFaceCCW(&c.R)
	top := [3]string{c.U[0][2], c.U[1][2], c.U[2][2]}
	front := [3]string{c.F[0][2], c.F[1][2], c.F[2][2]}
	bottom := [3]string{c.D[0][2], c.D[1][2], c.D[2][2]}
	back := [3]string{c.B[2][0], c.B[1][0], c.B[0][0]}

	c.U[0][2], c.U[1][2], c.U[2][2] = back[0], back[1], back[2]
	c.F[0][2], c.F[1][2], c.F[2][2] = top[0], top[1], top[2]
	c.D[0][2], c.D[1][2], c.D[2][2] = front[0], front[1], front[2]
	c.B[0][0], c.B[1][0], c.B[2][0] = bottom[2], bottom[1], bottom[0]
}

func (c *Cube) rotate(oprs string) {
	for i, v := range oprs {
		counter := false
		if i+1 < len(oprs) && oprs[i+1] == '\'' {
			counter = true
		}
		switch v {
		case 'U':
			if counter {
				c.rotateUPrime()
			} else {
				c.rotateU()
			}
		case 'D':
			if counter {
				c.rotateDPrime()
			} else {
				c.rotateD()
			}
		case 'F':
			if counter {
				c.rotateFPrime()
			} else {
				c.rotateF()
			}
		case 'B':
			if counter {
				c.rotateBPrime()
			} else {
				c.rotateB()
			}
		case 'L':
			if counter {
				c.rotateLPrime()
			} else {
				c.rotateL()
			}
		case 'R':
			if counter {
				c.rotateRPrime()
			} else {
				c.rotateR()
			}
		}
	}
}

func fillFace(color string) Face {
	var f Face
	for i := range f {
		for j := range f[i] {
			f[i][j] = color
		}
	}
	return f
}

func NewCube() *Cube {
	return &Cube{
		U: fillFace("W"), D: fillFace("Y"),
		F: fillFace("G"), B: fillFace("B"),
		L: fillFace("O"), R: fillFace("R"),
	}
}

func main() {
	cube := NewCube()
	fmt.Println("Initial U face:", cube)
	// cube.rotate("LL'RR'FF'BB'UU'DD'")
	cube.rotate("RU'RURURU'R'U'RR")
	fmt.Println("After U rotation:", cube)
	cube.rotate("LL'RR'FF'BB'UU'DD'")
	fmt.Println("After U rotation:", cube)
}
