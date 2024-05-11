package msg

import (
	"strings"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	// "robit_wechat/tools/global"
	"robit_wechat/tools/mysql"
	"time"
)

func HandleMsg(msg *openwechat.Message) {
	if msg.IsSendBySelf(){
		return
	}

	var(
		contentText = ""
		err 	error
	)


	if msg.IsJoinGroup(){
		parts := strings.Split(msg.Content, `"`)

		// 提取用户名
		invitor := parts[1]
		invitee := parts[3]	
		mysql.Addinvitor(invitor)
		number := mysql.GetInvitedNumber(invitor)
		
		if(number >=5 ){
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02")
			_,err = msg.ReplyText(fmt.Sprintf("%s\n@%s 小主你好，经过小深检索你已经邀请了 %d 人,您的邀请人数已达标，请在使用的时候查询你的邀请人数，并把记录给收银人员看即可使用\n@%s，欢迎加入深井烧腊粉丝群！我们会不定时发放福利哦！开业活动：邀请五个人进入群聊即可获得五元现金消费券",formattedTime,invitor,number, invitee)) 	
			if err != nil{
				err = errors.Wrapf(err,"err about reply of invitor")
				logrus.Errorf(err.Error())
			}
		}else{
			_,err = msg.ReplyText(fmt.Sprintf("@%s，您已经邀请了%d人，感谢支持\n@%s，欢迎加入深井烧腊粉丝群！我们会不定时发放福利哦！开业活动：邀请五个人进入群聊即可获得五元现金消费券", invitor,number, invitee)) 
			
			if err != nil{
					err = errors.Wrapf(err,"err about reply of invitor")
					logrus.Errorf(err.Error())
			}
		}
		return
	}


	if msg.IsText(){
		contentText = trimMsgContent(msg.Content)// 去除空格

		handleTextReplyBypass(msg,contentText)		
	}


}

func handleTextReplyBypass(msg *openwechat.Message,txt string){
	var(
		err error
		content = ""
	)

	if txt == "小深井" {
		_, err = msg.ReplyText(
			`你好！我是小深井，我是一个简单的微信机器人程序🤖🤖
			------------------------------
			输入【活动】，即可回复店铺开业活动
			输入【店铺介绍】，即可了解店铺背景
			输入【群内昵称+已经邀请了多少人？】，即可回复已邀请人数(中文问号)
			输入【菜单】，即可回复店铺菜单
			------------------------------
			注意：为了防止优惠券被复用，仅限当日需要使用优惠券时，输入【群内昵称+已经邀请了多少人？】，出示当日机器人回答记录给收银员即可使用优惠，如果邀请人数满足五人，机器人回答后邀请记录将清零哟～
			！！！！谨慎输入【群内昵称+已经邀请了多少人？】 感谢您的理解
			`)

		if err != nil{
			err = errors.Wrapf(err,"reply err about 功能介绍 ")
			logrus.Errorf(err.Error())
			return
		}
		return
	}
	
	if txt == "菜单" {
		_, err = msg.ReplyText(`
		卤味系列：
			潮卤鸭腿 港式香肠饭
			黄金猪手 潮卤耳尖
			香辣牛肚 潮卤小牛肉
		烧腊系列：
			招牌烧鸭 秘汁叉烧
			烧鸭拼叉烧	烧鸭拼牛肉
			烧鸭拼牛肚	烧鸭拼口试鸡
			牛肉拼牛肚
		特色系列：
			香辣鸡腿 秘制酱鸭
			口水鸡   风味烤肉
			湘西腊肉 泡椒牛蛙
			东坡肉	 卤肉饭
			`)

		if err != nil{
			err = errors.Wrapf(err,"reply err about 菜单 ")
			logrus.Errorf(err.Error())
			return
		}
		return
	}

	if txt == "活动"{
		_,err = msg.ReplyText(`新店开业！福利多多
			新店活动：邀请五人进群，可得五元现金抵用券！
			日常活动：每日红包手气最佳和0.01的同学可以五折消费一次！
			新店开业，欢迎您的光临！`)

		if err != nil{
			err = errors.Wrapf(err,"reply text msg err about activity",)
			logrus.Error(err.Error())
			return
		}
		return
	}

	if txt  == "店铺介绍"{
		_,err = msg.ReplyText(`✨🌟 欢迎来到深井烧腊 —— 大学生的美食首选！ 🌟✨
		在湘潭大学，湖南农业大学都有我们的身影，深受学生喜爱，被学生们誉为“排队王”，在湖南农大餐饮受欢迎榜第一名！😎😎
		
		在忙碌的学习生活中，寻找一顿既美味又实惠的饭菜？深井烧腊就是您的不二之选！我们位于贤德公寓8栋北侧，以传统的烧腊技艺和湘菜的口味结合，为您带来无与伦比的美食体验。
		
		特色烧鸭 —— 我们的招牌菜，选用优质鸭肉，经过精心腌制和独家秘方烤制，皮脆肉嫩，香气四溢。
		
		量大价廉 —— 我们明白每一分钱的价值，所以我们承诺提供的不只是一顿美味的饭菜，更是一份对您经济预算的深切体贴。在深井烧腊，您可以享受到既满足味蕾又不压钱包的美食体验。
		
		便捷打包 —— 无论是要赶回宿舍复习，还是在校园里享受悠闲午后，我们的打包服务让您随时随地都能享受到深井烧腊的美味。我们不收打包费！！！
		
		来享受美味，留下满足。在深井烧腊，每一口都是对传统美味的致敬，每一餐都是对日常生活的小确幸。期待您的光临，让我们一起分享美食的快乐！🥰🥰🥰
		`)

		if err != nil{
			err = errors.Wrapf(err,"reply text msg about introduction")
			logrus.Error(err.Error())
			return 
		}
		return
	}

	if strings.Contains(txt,"已经邀请了多少人？"){
		parts := strings.Split(txt, "已经邀请了")
    // parts[0] 将包含"userNickname"
    	userNickname := parts[0]

		number := mysql.GetInvitedNumber(userNickname)
		if number >= 5{
			currentTime := time.Now()
			formattedTime := currentTime.Format("2006-01-02")
			content = fmt.Sprintf("当前时间：%s\n @%s,经过小深检索您已经邀请了%d人，您的邀请人数已经达标。\n 当您看到这条消息您的邀请人数已清零",formattedTime,userNickname,number)
			mysql.Setzero(userNickname)
		}else{
			content = fmt.Sprintf("@%s,经过小深检索您已经邀请了%d人，感谢您的支持！还需要邀请%d人就可以换取优惠啦。",userNickname,number,5-number)
		}

		_,err = msg.ReplyText(content)
		if err != nil{
			errors.Wrapf(err,"error about reply invited")
			logrus.Errorf(err.Error())
			return
		}
		return
	}
}


func trimMsgContent(content string) string {
	content = strings.TrimLeft(content, " ")
	content = strings.TrimRight(content, " ")
	return content
}