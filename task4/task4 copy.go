package main // Треугольники в многоугольниках

import (
	"fmt"
	"math"
)

func main() {

	var N int

	N = 14
	// fmt.Scanln(&N)
	Poly := newPolygon(N)

	for _, v := range Poly.points {
		fmt.Printf("\t(%f; %f)", v.x, v.y)
	}
	fmt.Printf("\t(%f; %f)\n", Poly.points[0].x, Poly.points[0].y)

	// fmt.Printf("\n\n\t%s\n\t%f\n\n", "S=", MaxSOfTrianglesInPolygon(Poly))
	fmt.Printf("%f", MaxSOfTrianglesInPolygon(Poly))

	// for i := 3; i < 21; i++ {
	// 	Poly := newPolygon(i)
	// 	fmt.Printf("\n%v: %f", i, MaxSOfTrianglesInPolygon(Poly))
	// }

	var T [4][3]point
	T[0][0].x = 0.000000
	T[0][0].y = 2.246980
	T[0][1].x = -2.190643
	T[0][1].y = -0.500000
	T[0][2].x = 1.756759
	T[0][2].y = -1.400969

	T[1][0].x = -0.974928
	T[1][0].y = 2.024459
	T[1][1].x = -1.756759
	T[1][1].y = 1.400969
	T[1][2].x = -2.190643
	T[1][2].y = 0.500000

	T[2][0].x = -1.756759
	T[2][0].y = -1.400969
	T[2][1].x = 0.000000
	T[2][1].y = -2.246980
	T[2][2].x = 0.974928
	T[2][2].y = -2.024459

	T[3][0].x = 2.190643
	T[3][0].y = -0.500000
	T[3][1].x = 1.756759
	T[3][1].y = 1.400969
	T[3][2].x = 0.974928
	T[3][2].y = 2.024459

	sum := float64(0)
	for cur := range T {
		var t Triangle
		t.points[0] = T[cur][0]
		t.points[1] = T[cur][1]
		t.points[2] = T[cur][2]

		t.a = newVector(t.points[0], t.points[1])
		t.b = newVector(t.points[1], t.points[2])
		t.c = newVector(t.points[2], t.points[0])

		p := (t.a.length + t.b.length + t.c.length) / 2 // полупериметр

		t.Square = math.Sqrt((p) * (p - t.a.length) * (p - t.b.length) * (p - t.c.length))
		fmt.Printf("\n\n\n\nSQUARE=\t %v", t.Square)
		sum += t.Square
	}
	fmt.Printf("\n\n\n\nSUM=\t %v", sum)

}

func MaxSOfTrianglesInPolygon(p Polygon) float64 {
	var res float64
	// TheBigTr := newTriangle(p.points[len(p.points)/3], p.points[2*len(p.points)/3], p.points[3*len(p.points)/3])
	var Ai, Bi, Ci int
	Ai = 0
	Bi = len(p.points) / 3
	Ci = 2 * len(p.points) / 3
	TheBigTr := newTriangle(p.points[Ai], p.points[Bi], p.points[Ci])
	res += TheBigTr.Square

	// fmt.Printf("\t%s\t%v\t%v\t%v", "Triangle points", 0, len(p.points)/3, 2*len(p.points)/3)

	for _, v := range TheBigTr.points {
		fmt.Printf("\t(%f; %f)", v.x, v.y)
	}
	fmt.Printf("\t(%f; %f)\n\n", TheBigTr.points[0].x, TheBigTr.points[0].y)

	var Sector1, Sector2, Sector3 []point
	for i := 1; i < Bi; i++ {
		Sector1 = append(Sector1, p.points[i])
	}
	for i := Bi + 1; i < Ci; i++ {
		// fmt.Printf("\t%v", i)
		Sector2 = append(Sector2, p.points[i])
	}
	for i := Ci + 1; i < len(p.points); i++ {
		Sector3 = append(Sector3, p.points[i])
	}

	// fmt.Printf("\nSector1=\n%f\n\nSector2=\n%f\n\nSector3=\n%f\n\n", Sector1, Sector2, Sector3)

	res += MaxSOfTrianglesInSector(Sector1) + MaxSOfTrianglesInSector(Sector2) + MaxSOfTrianglesInSector(Sector3)

	return res
}

func MaxSOfTrianglesInSector(p []point) float64 {

	if len(p) < 3 {
		return 0
	}

	// if len(p) == 3 {
	// 	newT := newTriangle(p[0], p[1], p[2])
	// 	return newT.Square
	// }

	// var arr [3]point
	var arr1, arr2 []point

	newT := newTriangle(p[0], p[(len(p)/2)], p[len(p)-1])

	for _, v := range newT.points {
		fmt.Printf("\t(%f; %f)", v.x, v.y)
	}
	fmt.Printf("\t(%f; %f)\n\n", newT.points[0].x, newT.points[0].y)

	for i := 1; i < (len(p) / 2); i++ {
		arr1 = append(arr1, p[i])
	}
	for i := (len(p) / 2) + 1; i < len(p)-1; i++ {
		arr2 = append(arr2, p[i])
	}

	return (newT.Square + MaxSOfTrianglesInSector(arr1) + MaxSOfTrianglesInSector(arr2))
}

func SOfTriangles(t []Triangle) float64 {
	var Sum float64
	Sum = 0
	for _, v := range t {
		Sum += v.Square
	}
	return Sum
}

type point struct {
	x, y float64
}
type Polygon struct {
	a      float64
	points []point
	center point
}

type Triangle struct {
	Polygon
	points  [3]point
	a, b, c vector
	Square  float64
}

type vector struct {
	point
	length float64
}

func newVector(A point, B point) vector {
	var v vector
	v.x = B.x - A.x
	v.y = B.y - A.y
	v.length = float64(math.Sqrt(math.Pow(float64(v.x), 2) + math.Pow(float64(v.y), 2)))
	return v
}

func newPolygon(N int) Polygon {
	var p Polygon
	p.a = 1
	p.center.x = 0
	p.center.y = 0

	Radius := float64(p.a) / (2 * math.Sin(float64(math.Pi)/float64(N)))

	// fmt.Println(Radius)
	// angle := (float64(math.Pi) * float64(N-2)) / float64(N)
	angle := math.Pi - ((float64(math.Pi) * float64(N-2)) / float64(N))

	p.points = append(p.points, point{})
	p.points[0].y = Radius
	for i := 1; i < N; i++ {
		p.points = append(p.points, point{})
		p.points[i].x = p.center.x + p.points[i-1].x*math.Cos(angle) - p.points[i-1].y*math.Sin(angle)
		p.points[i].y = p.center.y + p.points[i-1].x*math.Sin(angle) + p.points[i-1].y*math.Cos(angle)
	}

	return p
}

func newTriangle(A, B, C point) Triangle {
	var t Triangle
	t.points[0] = A
	t.points[1] = B
	t.points[2] = C

	t.a = newVector(t.points[0], t.points[1])
	t.b = newVector(t.points[1], t.points[2])
	t.c = newVector(t.points[2], t.points[0])

	p := (t.a.length + t.b.length + t.c.length) / 2 // полупериметр

	t.Square = math.Sqrt((p) * (p - t.a.length) * (p - t.b.length) * (p - t.c.length))

	return t
}
