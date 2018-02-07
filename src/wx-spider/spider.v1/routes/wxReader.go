package routes
import (
	"fmt"
	"errors"
	"regexp"
	"encoding/json"
	"strings"
	"log"
	"github.com/PuerkitoBio/goquery"
)
func (s *Server) getWxReader(wxid string) (wxReader *WxReader){
	wxReader = &WxReader{ WxId: wxid }
	return
}

func (wx *WxReader) setWxInfo() (err error) {
	request := fmt.Sprintf("http://weixin.sogou.com/weixin?type=1&query=%s&ie=utf8&_sug_=y&_sug_type_=1", wx.WxId)
	doc, err := goquery.NewDocument(request)
	if err != nil {
		log.Fatal(err)
		return
	}
	wechatObj := doc.Find("#sogou_vr_11002301_box_0 a")
	url, exists := wechatObj.Attr("href")
	if !exists {
		url = ""
	}
	wechatName := doc.Find("#sogou_vr_11002301_box_0 [uigs=account_name_0]").First().Text()
	wx.Url = url
	wx.Name = wechatName
	return nil
}

/*
 * 根据微信公众号信息获取微信公众号文章列表
 */
func (wx *WxReader) setWxPostList() (err error) {
	doc, err := goquery.NewDocument(wx.Url)
	if err != nil {
		log.Fatal("获取图文信息列表失败:" + err.Error())
		err = errors.New("获取图文信息列表失败:" + err.Error())
		return
	}
	docHtml, _:= doc.Html()
	if strings.Index(docHtml, "为了保护你的网络安全，请输入验证码") > 0 {
		fmt.Println("验证码")
	} else {
		wx.PostList, err = wx.getPostList(docHtml)
		if err != nil {
			return
		}
	}
	return  nil
}

func (wx *WxReader) getPostList(postHtml string) (resp []map[string]interface{}, err error) {
	msgReg := regexp.MustCompile(`var msgList = ({.+}}]});?`)
	strReg := regexp.MustCompile(`(var msgList = |;$)|(&nbsp;)`)
	quotReg := regexp.MustCompile(`(&quot;)`)
	ampReg := regexp.MustCompile(`(amp;)|(\\")`)
	
	//文章数组,页面上是没有的,在js中,通过正则截取出来
	msgList := msgReg.FindAllString(postHtml, -1)
	if len(msgList) == 0 {
		err = errors.New("-没有搜索到"+ wx.Name + "的文章,只支持订阅号,服务号不支持!");
		return
	}
	//返回的最近20条，第二项是最近10条
	msgListStr := msgList[0]
	msgListStr = strReg.ReplaceAllString(msgListStr, "")
	msgListStr = quotReg.ReplaceAllString(msgListStr, "\\\"")

	type MsgListObj struct {
		List []map[string]interface{} `json:"list"`
	}
	var msgListObj MsgListObj
	if err = json.Unmarshal([]byte(msgListStr), &msgListObj); err != nil {
		return 
	}
	if len(msgListObj.List) == 0{
		err = errors.New("-没有搜索到"+ wx.Name + "的文章,只支持订阅号,服务号不支持!");
		return
	}
	var articles = make([]map[string]interface {}, len(msgListObj.List))
	for index, post:= range msgListObj.List {
		if post != nil {
			extInfo := post["app_msg_ext_info"].(map[string]interface {})
			contentUrl := extInfo["content_url"].(string)
			articleUrl := "http://mp.weixin.qq.com" + ampReg.ReplaceAllString(contentUrl, "")
			post["articleUrl"] = articleUrl
			post["wxName"] = wx.Name
			articles[index] = post
		}
	}
	resp = articles
	return
}