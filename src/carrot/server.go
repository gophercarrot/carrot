package carrot

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wcharczuk/go-chart"
)

func makeRange(min, max int) []float64 {
	a := make([]float64, max-min+1)
	for i := range a {
		a[i] = float64(min) + float64(i)
	}
	return a
}

func DrawChart(res http.ResponseWriter, req *http.Request, latency []float64, timeSeries []time.Time) {

	graph := chart.Chart{
		XAxis: chart.XAxis{
			Name:      "Requests Index",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxis: chart.YAxis{
			Name:      "Latency Count",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Style: chart.Style{
					Show:        true,
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: makeRange(1, len(latency)),
				YValues: latency,
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func renderHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello Carrot!</h1>")
}

func StartHTTPServer(port string, latency []float64, timeSeries []time.Time) {
	http.HandleFunc("/", renderHTML)
	http.HandleFunc("/latency", func(w http.ResponseWriter, r *http.Request) {
		DrawChart(w, r, latency, timeSeries)
	})
	fmt.Printf("HTTP Server Listening at... %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
