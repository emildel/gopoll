package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func createBar() {
	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: "Results",
	}))

	data := make([]opts.BarData, 0)
	data = append(data, opts.BarData{Value: 5})
	data = append(data, opts.BarData{Value: 2})
	data = append(data, opts.BarData{Value: 4})

	bar.SetXAxis([]string{"Yes", "No", "Maybe"}).
		AddSeries("Count", data)

	f, _ := os.Create("myBarChart.html")
	bar.Render(f)
}
