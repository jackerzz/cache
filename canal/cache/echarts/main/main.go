package main

import (
	"cache/canal/cache/echarts"
	"log"
	"net/http"

)

const (
	host   = "http://127.0.0.1:8080"
)





func main() {
	// Avoid "404 page not found".
	http.HandleFunc("/", echarts.LogTracing(echarts.BarHandler))

	http.HandleFunc("/bar", echarts.LogTracing(echarts.BarHandler))
	http.HandleFunc("/bar3D", echarts.LogTracing(echarts.Bar3DHandler))
	http.HandleFunc("/boxPlot", echarts.LogTracing(echarts.BoxPlotHandler))
	http.HandleFunc("/kline", echarts.LogTracing(echarts.KlineHandler))
	http.HandleFunc("/line", echarts.LogTracing(echarts.LineHandler))
	http.HandleFunc("/wordCloud", echarts.LogTracing(echarts.WcHandler))
	http.HandleFunc("/map", echarts.LogTracing(echarts.MapHandler))
	//http.HandleFunc("/effectScatter", logTracing(esHandler))
	//http.HandleFunc("/funnel", logTracing(funnelHandler))
	//http.HandleFunc("/gauge", logTracing(gaugeHandler))
	//http.HandleFunc("/geo", logTracing(geoHandler))
	//http.HandleFunc("/graph", logTracing(graphHandler))
	//http.HandleFunc("/heatMap", logTracing(heatMapHandler))
	//http.HandleFunc("/line3D", logTracing(line3DHandler))
	//http.HandleFunc("/liquid", logTracing(liquidHandler))

	//http.HandleFunc("/overlap", logTracing(overlapHandler))
	//http.HandleFunc("/parallel", logTracing(parallelHandler))
	//http.HandleFunc("/pie", logTracing(pieHandler))
	//http.HandleFunc("/radar", logTracing(radarHandler))
	//http.HandleFunc("/sankey", logTracing(sankeyHandler))
	//http.HandleFunc("/scatter", logTracing(scatterHandler))
	//http.HandleFunc("/scatter3D", logTracing(scatter3DHandler))
	//http.HandleFunc("/surface3D", logTracing(surface3DHandler))
	//http.HandleFunc("/themeRiver", logTracing(themeRiverHandler))

	//http.HandleFunc("/page", logTracing(pageHandler))

	log.Println("Run server at " + host)
	http.ListenAndServe(":8080", nil)
}
