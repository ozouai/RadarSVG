package main

import (
	"context"
	"os"

	"github.com/ozouai/radarsvg"
)

func main() {
	output, err := os.OpenFile("output.svg", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(0777))
	if err != nil {
		panic(err)
	}
	defer output.Close()
	radarsvg.Generate(context.TODO(), 200, []*radarsvg.Data{
		{
			Label: "Test",
			Min:   0,
			Max:   10,
			Value: 5,
		},
		{
			Label: "Test1",
			Min:   0,
			Max:   10,
			Value: 3,
		},
		{
			Label: "Test2",
			Min:   0,
			Max:   10,
			Value: 5,
		},
		{
			Label: "Test3",
			Min:   0,
			Max:   10,
			Value: 5,
		},
		{
			Label: "Test4",
			Min:   0,
			Max:   10,
			Value: 5,
		},
		{
			Label: "Test4",
			Min:   0,
			Max:   10,
			Value: 10,
		},
	}, output)
}
