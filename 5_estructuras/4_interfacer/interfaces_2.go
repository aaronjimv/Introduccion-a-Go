package main 

import(
	"fmt"
	"math"
)

type Forma interface{
	area() float64
}

type Circulo struct{
	x,y,r float64
}

func (c *Circulo) area() float64 {
	return math.Pi * c.r*c.r
}

type Rectangulo struct{
	x1,y1,x2,y2 float64
}

func distancia(x1,y1,x2,y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangulo) area() float64 {
	l := distancia(r.x1, r.y1, r.x1, r.y2)
	w := distancia(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type MultiForma struct{
	formas []Forma
}

func (m *MultiForma) area() float64{
	var areaToatal float64
	for _, f := range m.formas {
		areaToatal += f.area()
	}
	return areaToatal
}

func main() {
	
	c := &Circulo{0,0,5}
	r := &Rectangulo{0,0,10,10}

	multiForma := MultiForma{
		formas : []Forma{
			c,r,
		},
	}

	fmt.Println(c.area())
	fmt.Println(r.area())

	fmt.Println(multiForma.area())
}