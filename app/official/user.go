package official

import (
	"encoding/json"

	"github.com/godcong/wego/config"
	"github.com/godcong/wego/core"
	"github.com/godcong/wego/log"
	"github.com/godcong/wego/net"
	"github.com/godcong/wego/util"
)

type User struct {
	config.Config
	*OfficialAccount
}

func newUser(account *OfficialAccount) *User {
	return &User{
		Config:          defaultConfig,
		OfficialAccount: account,
	}
}

func NewUser() *User {
	return newUser(account)
}

// http请求方式: POST（请使用https协议）
// https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN
// POST数据格式：JSON
// POST数据例子：
// {
// "openid":"oDF3iY9ffA-hqb2vVvbr7qxf6A0Q",
// "remark":"pangzi"
// }
// 成功:
// {"errcode":0,"errmsg":"ok"}
// 失败:
// {"errcode":40013,"errmsg":"invalid appid"}
func (u *User) UpdateRemark(openid, remark string) *net.Response {
	log.Debug("User|UpdateRemark", openid, remark)
	p := u.token.GetToken().KeyMap()
	resp := u.client.HttpPostJson(
		u.client.Link(UserInfoUpdateRemarkUrlSuffix),
		util.Map{
			"openid": openid,
			"remark": remark,
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	return resp
}

// 接口调用请求说明
// http请求方式: GET
// https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
// 成功:
// {"subscribe":1,"openid":"o6_bmjrPTlm6_2sgVt7hMZOPfL2M","nickname":"Band","sex":1,"language":"zh_CN","city":"广州","province":"广东","country":"中国","headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0","subscribe_time":1382694957,"unionid":"o6_bmasdasdsad6_2sgVt7hMZOPfL""remark":"","groupid":0,"tagid_list":[128,2],"subscribe_scene":"ADD_SCENE_QR_CODE","qr_scene":98765,"qr_scene_str":""}
func (u *User) UserInfo(openid, lang string) *core.UserInfo {
	log.Debug("User|UpdateRemark", openid, lang)
	p := u.token.GetToken().KeyMap()
	p.Set("openid", openid)
	if lang != "" {
		p.Set("lang", lang)
	}

	resp := u.client.HttpGet(
		u.client.Link(UserInfoUrlSuffix),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})
	var info core.UserInfo
	json.Unmarshal(resp.ToBytes(), &info)

	return &info
}

// http请求方式: POST
// https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN
// 成功:
// {"user_info_list":[{"subscribe":1,"openid":"oLyBi0tDnybg0WFkhKsn5HRetX1I","nickname":"sean","sex":1,"language":"zh_CN","city":"浦东新区","province":"上海","country":"中国","headimgurl":"http:\/\/thirdwx.qlogo.cn\/mmopen\/anblvjPKYbMGjBnTVxw5gEZiasF6LiaMHheNxN4vWJcfCLRl8gEX0L6M7sNjtMkFYx8PJRCS1lr9RGxadkFlBibpA\/132","subscribe_time":1521022410,"remark":"nishi123","groupid":101,"tagid_list":[101],"subscribe_scene":"ADD_SCENE_PROFILE_CARD","qr_scene":0,"qr_scene_str":""},{"subscribe":1,"openid":"oLyBi0lCK5rQPuo0_cHJrjQ4J9XE","nickname":"🎀曉青青💋baby💞","sex":2,"language":"zh_CN","city":"浦东新区","province":"上海","country":"中国","headimgurl":"http:\/\/thirdwx.qlogo.cn\/mmopen\/ajNVdqHZLLAiae3G7CGiaF8I6nxDiczQIHSpEFSXwFQoP2v923ficqHdxnRoeZC1BAibXcQNkBOFsibBicMydnLE0UnKw\/132","subscribe_time":1521012452,"remark":"","groupid":0,"tagid_list":[],"subscribe_scene":"ADD_SCENE_QR_CODE","qr_scene":0,"qr_scene_str":""}]}
// 失败:
// {"errcode":40013,"errmsg":"invalid appid"}
func (u *User) BatchGet(openids []string, lang string) []*core.UserInfo {
	log.Debug("User|BatchGet", openids, lang)
	p := u.token.GetToken().KeyMap()
	var list []*core.UserId

	for _, v := range openids {
		if lang != "" {
			list = append(list, &core.UserId{
				OpenId: v,
				Lang:   lang,
			})
		} else {
			list = append(list, &core.UserId{
				OpenId: v,
			})
		}

	}
	resp := u.client.HttpPostJson(
		u.client.Link(UserInfoBatchGetUrlSuffix),
		util.Map{
			"user_list": list,
		},
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): p,
		})

	m := make(map[string][]*core.UserInfo)
	e := json.Unmarshal(resp.ToBytes(), &m)
	if e == nil {
		if v, b := m["user_info_list"]; b {
			return v
		}
	}
	return nil
}

//http请求方式: GET（请使用https协议）
//https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID
func (u *User) Get(nextOpenid string) *net.Response {
	log.Debug("User|Get", nextOpenid)
	query := u.token.GetToken().KeyMap()
	if nextOpenid != "" {
		query.Set("next_openid", nextOpenid)
	}

	resp := u.client.HttpGet(
		u.client.Link(UserGetUrlSuffix),
		util.Map{
			net.REQUEST_TYPE_QUERY.String(): query,
		})

	return resp
}