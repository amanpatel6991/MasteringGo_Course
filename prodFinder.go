package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	"golang.org/x/net/html"
	"net/smtp"
	//"log"
	//"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
	"time"
	"os"
)

/*ENTER PASSWORD AS OS ARGS TO MAIL ALERTS*/
func main() {
	fmt.Println("starting..........")
	//urlRedmiFlipkart := "https://www.flipkart.com/redmi-note-7s-sapphire-blue-32-gb/p/itmfgpbacwygthha?pid=MOBFGGJ8Z7ZQVHSZ&lid=LSTMOBFGGJ8Z7ZQVHSZ8GXYCV&marketplace=FLIPKART&srno=b_1_3&otracker=hp_omu_Top%2BOffers_2_3.dealCard.OMU_88559FM66S42_3&otracker1=hp_omu_PINNED_neo%2Fmerchandising_Top%2BOffers_NA_dealCard_cc_2_NA_view-all_3&fm=neo%2Fmerchandising&iid=92fabc64-3e66-4eaa-80b5-c67a8e2a5cf6.MOBFGGJ8Z7ZQVHSZ.SEARCH&ppt=browse&ppn=browse&ssid=v00g4ytcyo0000001570965788836"
	urlDryerFlipkart := "https://www.flipkart.com/wahl-argan-care-wchd8-1324-hair-dryer/p/itmf8y3gnzzcztsg?pid=HDRF8XJQHGFWHEYF&cmpid=product.share.pp"
	//urlEarbudsClubFact := "https://www.clubfactory.com/amp/item-PID-5060169.html"

	password := ""

	lenArgs := len(os.Args)

	if lenArgs > 1 {
		password = os.Args[1]
	}

	/*Redmi Flipkart*/
	func() {
		for {
			priceChan := make(chan string)
			go getCurrprice(urlDryerFlipkart , priceChan)

			fmt.Println("waiting for product price...............")
			priceData := <- priceChan

			priceData = strings.Replace(priceData,"₹","",-1)
			priceData = strings.Replace(priceData,",","",-1)
			priceData = strings.Replace(priceData," ","",-1)

			priceDataVal, err := strconv.Atoi(priceData)
			if err != nil {
				fmt.Println("error in converting price value : " , err)
			}
			fmt.Println("Current Price of product : ",priceDataVal)

			if priceDataVal < 12	00 {
				fmt.Println("1st item PRice in your range :" , priceDataVal)
				SendMailAlert("Redmi note 7s",strconv.Itoa(priceDataVal) , password)
			} else {
				fmt.Println("still not in range !!")
			}

			time.Sleep(time.Second * 1)
			//break                              //REMOVE AFTER TESTING..

			fmt.Println()
			fmt.Println()
		}
	}()


	/*Earbuds Club factory*/
	//for {
	//	priceChan := make(chan string)
	//	go getCurrprice(urlEarbudsClubFact , priceChan)
	//
	//	fmt.Println("waiting for earbuds price...............")
	//	priceData := <- priceChan
	//
	//	priceData = strings.Replace(priceData,"₹","",-1)
	//	priceData = strings.Replace(priceData,",","",-1)
	//	priceData = strings.Replace(priceData," ","",-1)
	//
	//	priceDataVal, err := strconv.Atoi(priceData)
	//	if err != nil {
	//		fmt.Println("error in converting price value : " , err)
	//	}
	//	fmt.Println("Current Price of product : ",priceDataVal)
	//
	//	if priceDataVal < 200 {
	//		fmt.Println("Earbuds PRice in your range :" , priceDataVal)
	//		SendMailAlert("Earbuds" ,strconv.Itoa(priceDataVal) , password)
	//	} else {
	//		fmt.Println("still not in range !!")
	//	}
	//
	//	time.Sleep(time.Second * 1)
	//	//break                              //REMOVE AFTER TESTING..
	//
	//	fmt.Println()
	//	fmt.Println()
	//}

	fmt.Println("end!!!!!!!!!")
}


func getCurrprice(url string , priceChan chan string) {
	resp , _ := http.Get(url)

	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)

	depth := false
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return
		case html.TextToken:
			if depth {
				// emitBytes should copy the []byte it receives,
				// if it doesn't process it immediately.
				tempPrice := string(z.Text())
				// fmt.Println("iteration price ::",tempPrice)
				priceChan <- tempPrice
				//break
			}
		case html.StartTagToken, html.EndTagToken:
			t := z.Token()
			//tn, _ := z.TagName()
			//fmt.Println("fdsfdsf",t.DataAtom)
			//for _, a := range t.Attr {
			if t.Data == "div"{
				//fmt.Println("sas : ",t.Attr)
				for _, a := range t.Attr {
					if a.Key == "class" && a.Val == "_1vC4OE _3qQ9m1" {    // Flipkart Redmi
						//fmt.Println("att : " , a.Key, a.Val)
						if tt == html.StartTagToken {
							depth = true
						} else {
							depth = false
						}
					} else {
						depth = false
					}
				}
			} else if t.Data == "span" {                                      //Earbuds Clubfactory
				for _, a := range t.Attr {
					if a.Key == "class" && a.Val == "price_number detail-page-price" {    // Flipkart Redmi
						//fmt.Println("att : " , a.Key, a.Val)
						if tt == html.StartTagToken {
							depth = true
						} else {
							depth = false
						}
					} else {
						depth = false
					}
				}
			} else {
				depth = false
			}

		}
	}
}

func SendMailAlert(productName ,price ,password string) {
	auth := smtp.PlainAuth("", "amanpatel6991@gmail.com", password, "smtp.gmail.com")
	msg := []byte("To: amanpatel6991@gmail.com\r\n" +
		"Subject: CHECK PRICE NOW!\r\n" +
		"\r\n" +
		"Get Your product"+productName+" at :"+price+"   .\r\n")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"amanpatel6991@gmail.com",
		[]string{"amanpatel6991@gmail.com"},
		msg,
	)

	fmt.Println("err in sending mail.. :: " , err)
}