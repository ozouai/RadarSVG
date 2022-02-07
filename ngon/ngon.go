package ngon

import (
	"bytes"
	"context"
	"math"
	"strconv"

	svg "github.com/ajstarks/svgo"
)

func NGon(ctx context.Context, canvas *svg.SVG, radius int, sides int, style ...string) {
	path := CalculateVertices(ctx, radius, sides)
	canvas.Path(VerticesToPath(path), style...)
}

func CalculateVertices(ctx context.Context, radius int, sides int) []*Vertix {
	var path []*Vertix
	for i := 0; i < sides; i++ {
		path = append(path, CalculateVertix(ctx, radius, sides, i))
	}
	return path
}

func CalculateVertix(ctx context.Context, radius int, sides int, side int) *Vertix {
	angle := (2 * math.Pi) / float64(sides)
	return &Vertix{
		X: math.Sin(angle*float64(side)) * float64(radius) / 2.0,
		Y: math.Cos(angle*float64(side)) * float64(radius) / 2.0,
	}
}

type Vertix struct {
	X float64
	Y float64
}

func (m Vertix) IntX() int {
	return int(math.Floor(m.X))
}
func (m Vertix) IntY() int {
	return int(math.Floor(m.Y))
}

func VerticesToPath(vertices []*Vertix) string {
	buf := bytes.NewBufferString("")
	buf.WriteString("M ")
	buf.WriteString(strconv.Itoa(vertices[0].IntX()))
	buf.WriteByte(',')
	buf.WriteString(strconv.Itoa(vertices[0].IntY()))
	buf.WriteString(" ")
	for _, v := range vertices[1:] {
		buf.WriteString("L ")
		buf.WriteString(strconv.Itoa(v.IntX()))
		buf.WriteByte(',')
		buf.WriteString(strconv.Itoa(v.IntY()))
		buf.WriteString(" ")
	}
	buf.WriteString("L ")
	buf.WriteString(strconv.Itoa(vertices[0].IntX()))
	buf.WriteByte(',')
	buf.WriteString(strconv.Itoa(vertices[0].IntY()))
	buf.WriteString(" ")
	return buf.String()
}
