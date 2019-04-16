package plotlib

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"

	"github.com/golang/freetype/truetype"
	"github.com/mattn/go-pairplot"
	"github.com/olekukonko/tablewriter"
)

func MarkdownCSV(filename string) string {
	var buf bytes.Buffer
	table, err := tablewriter.NewCSV(&buf, filename, true)
	if err != nil {
		log.Fatal(err)
	}
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.Render()
	return buf.String()
}

func PairPlotCSV(filename string) []byte {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	pp, err := pairplot.NewPairPlot(filename)
	if err != nil {
		log.Fatal(err)
	}
	pp.Hue = "Name"
	p.HideAxes()
	p.Add(pp)
	w, err := p.WriterTo(4*vg.Inch, 4*vg.Inch, "png")
	if err != nil {
		log.Fatal(err)
	}
	var b bytes.Buffer
	writer := bufio.NewWriter(&b)
	w.WriteTo(writer)
	return b.Bytes()
}

func DefaultFont(fontpath, fontname string) {
	bb, err := ioutil.ReadFile(fontpath)
	if err != nil {
		log.Fatal(err)
	}
	ttf, err := truetype.Parse(bb)
	if err != nil {
		log.Fatal(err)
	}
	vg.AddFont(fontname, ttf)
	plot.DefaultFont = fontname
}
