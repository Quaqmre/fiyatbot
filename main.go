package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var item string
var itemName string
var smtpAddr string
var from string
var to string
var password string
var price int
var interval int

func init() {
	flag.StringVar(&item, "item", "none", "select item url")
	flag.StringVar(&itemName, "itemName", "none", "select item url")
	flag.StringVar(&smtpAddr, "smtp", "none", "select smtp server addres")
	flag.StringVar(&from, "from", "none", "select from")
	flag.StringVar(&to, "to", "none", "select to")
	flag.StringVar(&password, "pass", "none", "select password")
	flag.IntVar(&price, "price", 0, "select price")
	flag.IntVar(&interval, "interval", 1, "select price")

}

func main() {
	flag.Parse()
	// if item == "none" || itemName == "none" || smtpAddr == "none" || from == "none" || to == "none" || password == "none" || price == 0 {
	log.Printf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", item, itemName, smtpAddr, from, to, password, price, interval)
	// }
	for {
		res, err := http.Get(item)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		// Load the HTML document
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		list := make([]int, 0)

		// Find the review items
		doc.Find(".urun_fiyat").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			str, _ := s.Attr("data-sort")
			price, er := strconv.Atoi(str)
			if err != nil {
				fmt.Println(er)
			}
			list = append(list, price)
		})
		lowPrice := list[0]
		for _, v := range list {
			if v < lowPrice {
				lowPrice = v
			}
		}
		fmt.Printf("%v bulunan low price:%v\n", time.Now(), lowPrice)
		if lowPrice < price {
			auth := smtp.PlainAuth("", from, password, smtpAddr)
			toList := []string{to}
			msg := fmt.Sprintf("To:%s\r\nSubject:%s fiyat:%v\r\n", to, itemName, lowPrice)
			err = smtp.SendMail(smtpAddr+":587", auth, from, toList, []byte(msg))
			if err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Minute * time.Duration(interval))
	}
}
