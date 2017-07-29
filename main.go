package main

import (
	"github.com/michlabs/fbbot"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

const (
	PAGEACCESS_TOKEN="EAAZA8vTvPfkcBAJyhSj5KrpEBm9VNH6NkVsOHZAUuhCkkSJ82JtZC6qaUK4NXctks9vC0ulojVc6DyQQCGw5yaPGTBss3iFj2uMfqJDnfbPePB63B0q3dS8oCeEvbD6kZAn42IYKyqZBIKyzmYpYZBYnM845c9xZAyCEppQ1a175tq9YmmzXZA5h"
	VERIFY_TOKEN ="1234ratbimat"
	PORT =8080
)
type Warning struct{
	
}

type tennguoidung string
type solansudung int

var xuathien= make(map[tennguoidung]solansudung)

func main()  {
	var warning Warning
	
	//init a bot
	bot:=fbbot.New(PORT,VERIFY_TOKEN,PAGEACCESS_TOKEN)
	
	//bot call AddMessageHandle
	bot.AddMessageHandler(warning)
	bot.Run()

}

func(warning Warning) HandleMessage(bot *fbbot.Bot, msg *fbbot.Message)  {

	//khoi tao request

	if isFirstUsing(string(msg.Sender.FullName())) {
		xuathien[tennguoidung(msg.Sender.FullName())]=0


		chaohoi1:="xin chao " + msg.Sender.FullName()+", rất vui khi được đồng hành cùng bạn"
		m1:=fbbot.NewTextMessage(chaohoi1)
		bot.Send(msg.Sender,m1)
		chaohoi2:="Tên mình là neitteiv"
		m2:=fbbot.NewTextMessage(chaohoi2)
		bot.Send(msg.Sender,m2)

		chaohoi3:="Khi nghi ngờ đường link của bạn nguy hiểm, hãy gửi nó cho tôi"
		m3:=fbbot.NewTextMessage(chaohoi3)
		bot.Send(msg.Sender,m3)


		chaohoi4:="Hãy nhập đường link và gửi nó cho tôi: "
		m4:=fbbot.NewTextMessage(chaohoi4)
		bot.Send(msg.Sender,m4)
	}else {
		fmt.Println("Hãy nhập đường link và gửi nó cho tôi")
	}
		messenger:=msg.Text
		//tach lay url
		user_url:=checkStringHasurl(messenger)


		if strings.Compare(user_url,"")!=0 {//NEU TIN NHAN GUI DEN CO CHUA DUONG LINK
		req, err := http.NewRequest("GET", "http://api.openfpt.vn/cyradar?url="+user_url,nil)
		if err!=nil {
			fmt.Sprint("tao request hong")
		} else {
			fmt.Sprint("da tao xong request")
		}
		req.Header.Set("api_key","6631fdd937b547479fe036c5420863fc")

		//day request len server
		client := &http.Client{}
		res, err :=client.Do(req)



		defer res.Body.Close()

		body, _ := ioutil.ReadAll(res.Body)

		//giai ma file json tra ve
		var data Tracks

		err = json.Unmarshal(body, &data)

		if err != nil {
			log.Fatal("Failed to parse json", err.Error())
		}else{
			fmt.Println("parse xong du lieu")
			// fmt.Println(string(body))
			fmt.Println(data.Conclusion)
		}


		//tra loi nguoi dung
		var reply string
		switch data.Conclusion {
		case "danger":
			reply="warning: đường link NGUY HIỂM, bạn hãy cân nhắc trước khi click vào"
		default:
			reply="đường link AN TOÀN, bạn hãy yên tâm truy nhập"

		}

		m:=fbbot.NewTextMessage(reply)
		bot.Send(msg.Sender,m)

		tambiet:="Cảm ơn bạn đã sử dụng dịch vụ của chúng tôi"
		m0:=fbbot.NewTextMessage(tambiet)
		bot.Send(msg.Sender,m0)
	}
}

type Tracks struct {

	Conclusion string   `json: "conclusion"`
	Domain     string   `json: "domain"`
	Threat     []string `json: "threat"`
	Uri        string   `json: "uri`
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	} else {
		return true
	}
}

func checkStringHasurl(str string ) string {
	slice := strings.Fields(str)
	for _, val := range slice {
		if isValidUrl(val) {
			return val
		}
	}
	return ""

}
func isFirstUsing(tennguoidung string) bool {
	for value,_:=range xuathien{
		if strings.Compare(string(value),tennguoidung)==0 {
			return false
		}

	}
	return true
}