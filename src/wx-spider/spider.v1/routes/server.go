package routes

import (
	"sync"
	"wx-spider/spider.v1/config"
)

func NewServer(config *config.Config) (server *Server, err error) {
	server = &Server{
		ids: make(map[int64]bool),
		wxIdList: config.WxIdList,
	}
	return
}

type Server struct {
	sync.Mutex
	ids  map[int64]bool
	wxIdList []map[string]string
}
