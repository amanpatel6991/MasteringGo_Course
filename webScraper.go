package main

import (
	"net/http"
	//"io/ioutil"
	"fmt"
	"golang.org/x/net/html"
	//"log"
	//"github.com/PuerkitoBio/goquery"
	"strings"
	"strconv"
)

func main() {
	fmt.Println("starting..........")
	url := "https://www.flipkart.com/redmi-note-7s-sapphire-blue-32-gb/p/itmfgpbacwygthha?pid=MOBFGGJ8Z7ZQVHSZ&lid=LSTMOBFGGJ8Z7ZQVHSZ8GXYCV&marketplace=FLIPKART&srno=b_1_3&otracker=hp_omu_Top%2BOffers_2_3.dealCard.OMU_88559FM66S42_3&otracker1=hp_omu_PINNED_neo%2Fmerchandising_Top%2BOffers_NA_dealCard_cc_2_NA_view-all_3&fm=neo%2Fmerchandising&iid=92fabc64-3e66-4eaa-80b5-c67a8e2a5cf6.MOBFGGJ8Z7ZQVHSZ.SEARCH&ppt=browse&ppn=browse&ssid=v00g4ytcyo0000001570965788836"
	priceChan := make(chan string)
	go getCurrprice(url , priceChan)

	fmt.Println("waiting for price ........ .......")
	priceData := <- priceChan

	priceData = strings.Replace(priceData,"₹","",-1)
	priceData = strings.Replace(priceData,",","",-1)
	priceDataVal, err := strconv.Atoi(priceData)
	if err != nil {
		fmt.Println("error in converting price value : " , err)
	}
	fmt.Println("in main::",priceDataVal)

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
				fmt.Println("iteration price ::",tempPrice)
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
					if a.Key == "class" && a.Val == "_1vC4OE _3qQ9m1" {
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