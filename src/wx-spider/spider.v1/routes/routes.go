package routes

import (
	"log"
	"wx-spider/libs/install.v1"
	"wx-spider/spider.v1/config"
)

func Register(r *install.Router, c *config.Config) {
	s, err := NewServer(c)
	if err != nil {
		log.Fatalln(err)
	}
	r.GET("/api/wxIdList", s.GetWxIdList)
	r.GET("/api/wxPostList", s.DoSearch)
	r.GET("/api/nearlyPost", s.GetNearlyPostList)
	go s.SetupSpiderPlan()
}