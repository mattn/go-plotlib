package plotlib

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/font"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"

	"github.com/mattn/go-pairplot"
	"github.com/olekukonko/tablewriter"
	"golang.org/x/image/font/opentype"
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
	p := plot.New()
	pp, err := pairplot.NewPairPlotCSV(filename)
	if err != nil {
		log.Fatal(err)
	}
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
	ttf, err := opentype.Parse(bb)
	if err != nil {
		log.Fatal(err)
	}
	mfont := font.Font{Typeface: font.Typeface(fontname)}
	font.DefaultCache.Add([]font.Face{
		{
			Font: mfont,
			Face: ttf,
		},
	})
	plot.DefaultFont = mfont
	plotter.DefaultFont = mfont
}
