package routes

import (
	"encoding/json"
	"net/http"
	"errors"
	"time"
	_"strconv"
	"wx-spider/libs/spec.v1"
	"github.com/julienschmidt/httprouter"
)

type WxReader struct {
	Url string `json:"url"`
	Name string `json:"name"`
	WxId string `json:"wxId"`
	PostList []map[string]interface{} `json:"postList"`
}
func (s *Server) GetWxIdList(w http.ResponseWriter, r *http.Request,  params httprouter.Params) {
	resp, statusCode := spec.MarshalResponse(s.wxIdList, nil)
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func (s *Server) DoSearch(w http.ResponseWriter, r *http.Request,  params httprouter.Params) {
	values := r.URL.Query()
	wxid := values.Get("wxid")
	data, err := s.doSearch(wxid)
	resp, statusCode := spec.MarshalResponse(data, err)
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func (s *Server) doSearch(wxid string) (wxReader *WxReader, err error){
	if wxid == "" {
		err = errors.New("no wxid")
		return
	}
	wxReader = s.getWxReader(wxid)
	err = wxReader.setWxInfo()
	if err != nil {
		return
	}
	err = wxReader.setWxPostList()
	if err != nil {
		return
	}
	return
}

func(s *Server) GetNearlyPostList(w http.ResponseWriter, r *http.Request,  params httprouter.Params) {
	data, err := s.getNearlyPostList()
	resp, statusCode := spec.MarshalResponse(data, err)
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func(s *Server) getNearlyPostList() (posts []interface{}, err error) {
	for _, wx := range s.wxIdList {
		localData, err := s.readLocalPostById(wx["wxId"])
		if err != nil {
			return posts, err
		}
		var result map[string]interface{}
		resultStr := localData["result"].(string)
		if err = json.Unmarshal([]byte(resultStr), &result); err != nil {
			return posts, err
		}
		if result != nil && result["postList"] != nil{
			postList := result["postList"].([]interface{})
			posts = append(posts, s.getPostUnderTime(postList, 3)...) 
		}
	}
	return
}

func(s *Server) getPostUnderTime(postList []interface{}, day int) (results []interface{}) {
	now := time.Now()
	for _, post := range postList {
		postMap := post.(map[string]interface{})
		info := postMap["comm_msg_info"].(map[string]interface{})
		datetime := int64(info["datetime"].(float64))
		creatTime := time.Unix(datetime, 0)
		diff := now.Sub(creatTime)  
		if diff.Hours()/24 < float64(day) {
			results = append(results, post)
		}
	}
	return
}