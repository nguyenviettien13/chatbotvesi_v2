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
	PAGEACCESS_TOKEN="EAAZA8vTvPfkcBANRZC3JKnHGXS4Ic8oZAFkQheEbeshbZB7y6YG80SGTRw2haVbxRDRWjerVAWeJ6PbLVrbsfro349Tgsl7rARUZBhZAq68FPJ3n7ZCZA4QZCHZCYEwhhQGLdFUmBwQFlwJenFVlUnlrp009njtWtHoTP1xBKsreOeE3GSigbTJuM0"
	VERIFY_TOKEN ="ratbimat"
	PORT =8080
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
		chaohoi:="xin chao " + msg.Sender.FullName()+", rất vui khi được đồng hành cùng bạn"
		m0:=fbbot.NewTextMessage(chaohoi)
		bot.Send(msg.Sender,m0)
		gioithieu:="Tên mình là neitteiv"
		m01:=fbbot.NewTextMessage(gioithieu)
		bot.Send(msg.Sender,m01)

		reply:="Khi nghi ngờ đường link của bạn nguy hiểm, hãy gửi nó cho tôi"
		m:=fbbot.NewTextMessage(reply)
		bot.Send(msg.Sender,m)


		reply1:="Hãy nhập đường link và gửi nó cho tôi: "
		m2:=fbbot.NewTextMessage(reply1)
		bot.Send(msg.Sender,m2)

	} else{
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