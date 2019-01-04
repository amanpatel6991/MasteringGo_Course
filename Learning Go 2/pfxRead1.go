package main

import (
	"fmt"
	"io/ioutil"
	"golang.org/x/crypto/pkcs12"
)

func main() {
	b, err := ioutil.ReadFile("idsrv3test.pfx")
	//b, err:= ioutil.ReadFile("1571753451.p12")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(string(b))

	key, cert, err1 := pkcs12.DecodeAll(b, "idsrv3test")
	//key , cert , err1 := pkcs12.Decode(b,"test")
	if err1 != nil {
		fmt.Println("ERROR :", err1)
	}

	for _, v := range key {
		fmt.Println("KEY ::", v)
	}

	for _, v := range cert {
		fmt.Println("CERT ::", v)
		fmt.Println("CERT DATAA::", string(v.RawSubject))
		fmt.Println("CERT ISSUER::", string(v.RawIssuer))
		fmt.Println("?/////////////////////////////////")
	}

	//fmt.Println("??????????????????????")
	//
	//pem , err2 := pkcs12.ToPEM(b,"test")
	//if err2!=nil {
	//	fmt.Println("ERROR :",err2)
	//}
	//
	//fmt.Println("PEM ::",pem)


}
