package main

import (
	"fmt"
	"github.com/golang/snappy"
	"gopkg.in/olivere/elastic.v2"
	"github.com/go-echarts/go-echarts/charts"
	"log"
	"math/rand"
	"os"
	"time"
)

const test = `{"Ag(T+D)":{"instID":"Ag(T+D)","name":"白银延期","last":"4141","upDown":"21","upDownRate":"0.51","quoteDate":"20170328","quoteTime":"22:34:29"},"Au(T+D)":{"instID":"Au(T+D)","name":"黄金延期","last":"280.55","upDown":"0.88","upDownRate":"0.31","quoteDate":"20170328","quoteTime":"22:34:15"},"mAu(T+D)":{"instID":"mAu(T+D)","name":"Mini黄金延期","last":"280.5","upDown":"0.7","upDownRate":"0.25","quoteDate":"20170328","quoteTime":"22:34:10"}}`

func main() {
	elastics()
}
func msp(){
	ages := make(map[string]int)
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
		"hanhah" : 3242,
		"sb"	: 2332,
	}
	for k, v := range ages{
		fmt.Println(k,v)
	}

}

func elastics(){
	type Tweet struct {
		User    string
		Message string
	}
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.225.130:9200/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es succ")

	for i := 21; i < 20; i++ {
		tweet := Tweet{User: "olivere", Message: "Take Five"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do()
		if err != nil {
			// Handle error
			panic(err)
			return
		}
	}

	fmt.Println("insert succ")
}

func snappys(){
	fmt.Println("source len:", len(test))

	got := snappy.Encode(nil, []byte(test))
	fmt.Println("compressed len:", len(got))

	a, _ := snappy.Decode(nil, got)
	fmt.Println("uncompressed len:", len(a))
	fmt.Println("--------------------------------------------")
}
var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63()) % 50)
	}
	return r
}
func run(){
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	f, err := os.Create("bar.html")
	if err != nil {
		log.Println(err)
	}
	bar.Render(f)
}