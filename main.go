package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
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
	auth := smtp.PlainAuth("", from, password, smtpAddr)
	subject := fmt.Sprintf("%s başladı", itemName)
	bdy := fmt.Sprintf("Date:%v", time.Now())
	SendMail(from, to, subject, bdy, smtpAddr, auth)
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
			subject := fmt.Sprintf("%s fiyat:%v", itemName, lowPrice)
			bdy := fmt.Sprintf("Date:%v", time.Now())
			SendMail(from, to, subject, bdy, smtpAddr, auth)
			if err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Minute * time.Duration(interval))
	}
}

func SendMail(fromS, toS, subjS, bodyS, smtpAdd string, auth smtp.Auth) {
	from := mail.Address{"", fromS}
	to := mail.Address{"", toS}
	subj := subjS
	body := bodyS

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := smtpAdd + ":465"

	host, _, _ := net.SplitHostPort(servername)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()
}
