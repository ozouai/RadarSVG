package radarsvg

import (
	"context"
	"io"
	"math"
	"strconv"

	svg "github.com/ajstarks/svgo"
	"github.com/ozouai/radarsvg/ngon"
)

type Data struct {
	Label string
	Max   int
	Min   int
	Value int
}

func Generate(ctx context.Context, radius int, data []*Data, output io.Writer) {
	canvas := svg.New(output)
	canvas.Start(radius+30, radius+30)
	canvas.Translate(radius/2+15, radius/2+15)
	canvas.Style("text/css", `
.radarsvg-outline {
	fill: white;
	stroke: lightgrey;
}
.radarsvg-subline {
	fill: none;
	stroke: lightgrey;
}
.radarsvg-datumline {
	stroke: lightgrey;
}
.radarsvg-dataline {
	fill: rgba(255, 0, 0, 0.5);
}
.radarsvg-hovercircle {
	fill: none;
	stroke: none;
}
.radarsvg-point {
	fill: rgb(50, 125, 75);
	text-anchor: middle;
	alignment-baseline: central;
}
.radarsvg-label {
	text-anchor: middle;
	alignment-baseline: central;
}
.radarsvg-labelleft {
	text-anchor: start;
}
.radarsvg-labelright {
	text-anchor: end;
}
.radarsvg-labeltop {
	alignment-baseline: before-edge;
}
.radarsvg-labelbottom {
	alignment-baseline: after-edge;
}`)
	ngon.NGon(ctx, canvas, radius, len(data), `class="radarsvg-outline"`)
	radiusStep := radius / 4
	for i := 1; i <= 4; i++ {
		ngon.NGon(ctx, canvas, radius-(radiusStep*i), len(data), `class="radarsvg-subline"`)
	}

	vertices := ngon.CalculateVertices(ctx, radius, len(data))
	for _, v := range vertices {
		canvas.Line(0, 0, v.IntX(), v.IntY(), `class="radarsvg-datumline"`)

	}

	var datumPoints []*ngon.Vertix

	for i, d := range data {
		normalized := mapNumber(float64(d.Value), float64(d.Min), float64(d.Max), 0, 1)
		length := normalized * float64(radius)
		datumPoints = append(datumPoints, ngon.CalculateVertix(ctx, int(math.Floor(length)), len(data), i))
	}

	canvas.Path(ngon.VerticesToPath(datumPoints), `class="radarsvg-dataline"`)
	for i, v := range vertices {
		textClassList := "radarsvg-label "
		if v.IntX() < -2 {
			textClassList += "radarsvg-labelleft "
		} else if v.IntX() > 2 {
			textClassList += "radarsvg-labelright "
		}
		if v.IntY() < -2 {
			textClassList += "radarsvg-labeltop "
		} else if v.IntY() > 2 {
			textClassList += "radarsvg-labelbottom "
		}
		canvas.Text(v.IntX(), v.IntY(), data[i].Label, `class="`+textClassList+`"`)
	}
	for i, p := range datumPoints {
		canvas.Circle(p.IntX(), p.IntY(), 1, `class="radarsvg-point"`)
		canvas.Circle(p.IntX(), p.IntY(), 3, `class="radarsvg-hovercircle"`, `data-radarsvg-i="`+strconv.Itoa(i)+`"`)
	}

	canvas.Gend()
	canvas.End()
}

func mapNumber(num float64, min float64, max float64, targetMin float64, targetMax float64) float64 {
	return targetMin + ((targetMax-targetMin)/(max-min))*(num-min)
}
