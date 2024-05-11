package global

import(
	"github.com/eatmoreapple/openwechat"
	"robit_wechat/tools/conf"
)

var(
	Conf *conf.Conf
	WxSelf *openwechat.Self
	WxFriends openwechat.Friends
	WxGroups openwechat.Groups
)