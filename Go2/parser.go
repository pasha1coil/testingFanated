package main

import (
	"encoding/csv"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"log"
	"os"
)

type InstagramData struct {
	Rank      string
	Instagram string
	Name      string
	Category  string
	Followers string
	Country   string
	EngAuth   string
	EngAvg    string
}

func main() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{"https://hypeauditor.com/top-instagram-all-russia/"},
		ParseFunc: parseInst,
	}).Start()
}

func parseInst(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("div.row").Each(func(i int, s *goquery.Selection) {

		data := InstagramData{
			Rank:      s.Find("div.rank span").Eq(0).Text(),
			Instagram: s.Find("div.contributor__name-content").Text(),
			Name:      s.Find("div.contributor__title").Text(),
			Category:  s.Find("div.category").Text(),
			Followers: s.Find("div.subscribers").Text(),
			Country:   s.Find("div.audience").Text(),
			EngAuth:   s.Find("div.authentic").Text(),
			EngAvg:    s.Find("div.engagement").Text(),
		}
		err := exportToCSV(data)
		if err != nil {
			log.Fatal("Failed to export to csv file", err)
		}
	})
}

func exportToCSV(data InstagramData) error {
	file, err := os.OpenFile("InstagramData.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open file", err)
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var row []string
	row = append(row, data.Rank, data.Instagram, data.Name, data.Category, data.Followers, data.Country, data.EngAuth, data.EngAvg)
	err = writer.Write(row)
	if err != nil {
		log.Fatal("Failed to write to file", err)
		return err
	}
	return nil
}
