package ticker

import (
	// "fmt"
	"fmt"
	"math/rand"
	"robit_wechat/tools/global"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func SendToMyself(){

	fh := global.WxSelf.FileHelper()
	global.WxSelf.SendTextToFriend(fh,"不准休息")

}

func SendMessageToLover(txt string){
	var(
		err error
	)

	// 初始化随机数生成器
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	// 定义一个包含情话的切片
	loveMessages := []string{
		"答案很长，我准备用一生时间来回答，你准备听了吗？",
		"春风再美也比不上你的笑，没见过你的人不会明了",
		"夜阑卧听风吹雨，铁马是你，冰河也是你。",
		"于千万人之中遇见你所要遇见的人，于千万年之中，时间的无涯的荒野里，没有早一步，也没有晚一步，刚巧赶上了，那也没有别的话可说，惟有轻轻地问一声：“噢，你也在这里吗？",
		"我走过许多地方的路，行过许多地方的桥，看过许多次数的云，喝过许多种类的酒，却只爱过一个正当最好年龄的人。",
	}

	// 随机选择一句情话
	message := loveMessages[rng.Intn(len(loveMessages))]

	content := txt + message

	err = global.WxFriends.SearchByRemarkName(1,"熙宝").SendText(content)
	if err != nil {
		err = errors.Wrapf(err, "SendMessageToLover err")
		logrus.Error(err.Error())
	}

}

func SendMessageToConstomer(){
	var err error

	err = global.WxGroups.SearchByNickName(1,"深井烧腊(工商店)").SendText("饭点到啦，深井烧腊欢迎您的光临")
	if err != nil {
		err = errors.Wrapf(err, "SendMessageToCUS err")
		logrus.Error(err.Error())
	}
}

func Msg_OntimeTicker(){
	SendToMyselfTicker := time.NewTicker(10*time.Minute)
	defer SendToMyselfTicker.Stop()

	for {
		select{
		case t := <-time.After(1*time.Minute):
			nowTime := t.Format("15:04")

			if nowTime == "8:00"{
				SendMessageToLover("早安！ ")
			}

			if nowTime == "11:40"{
				SendMessageToConstomer()
			}

			if nowTime == "17:40"{
				SendMessageToConstomer()
			}

			if nowTime == "23:59"{
				SendMessageToLover("晚安！ 我爱你。")
			}

		case <- SendToMyselfTicker.C:
			fmt.Println("SendToMyself function called")
			go SendToMyself()
		}	
	}
}