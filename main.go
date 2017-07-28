package main

import (
	"github.com/michlabs/fbbot"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"log"
	"net/url"
)

const (
	PAGEACCESS_TOKEN="get page access token from fb page"
	VERIFY_TOKEN ="string of password"
	PORT =4040
)
type Warning struct{
	
}

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
	user_url:=msg.Text
	ok:= isValidUrl(user_url)
	if !ok {
		reply:="Moi ban gui cho toi duong link can kiem tra/n Luu y la chi duong link"
		m:=fbbot.NewTextMessage(reply)
		bot.Send(msg.Sender,m)
	} else{
		fmt.Println(user_url)
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
			reply="nguy hiem"
		default:
			reply="an toan"
		}

		m:=fbbot.NewTextMessage(reply)
		bot.Send(msg.Sender,m)
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