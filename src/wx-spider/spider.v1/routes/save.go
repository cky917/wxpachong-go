package routes

import(
	"fmt"
	"sync"
	"time"
	"github.com/robfig/cron"
)
func(s *Server) SetupSpiderPlan() {
	c:= cron.New()
	spec := "1 26 1,7,11,14,16,20 * * *"
	c.AddFunc(spec, func() {
		go s.DoSave()
	})
	c.Start()
	select{}
}

func(s *Server) DoSave() {
	fmt.Println("开始执行定时任务：", time.Now())
	var mutex = &sync.Mutex{}
	index := 0
	c:= cron.New()
	spec := "*/30 * * * * ?"
	c.AddFunc(spec, func() {
		mutex.Lock()
		fmt.Println(index, len(s.wxIdList))
		if index == len(s.wxIdList) - 1{
			c.Stop()
		}
		go s.doSave(s.wxIdList[index])
		index++
		mutex.Unlock()
	})  
	c.Start()  
	select{}
}

func(s *Server) doSave(wxInfo map[string]string) {
	wxId := wxInfo["wxId"]
	wxName := wxInfo["name"]
	fmt.Println("开始爬取" + wxName + "的文章")
	wxReader, err:= s.doSearch(wxId)
	if err != nil {
		fmt.Println("获取" + wxName + "文章失败", err.Error())
	}
	err = s.savePostToLocal(wxReader)
	if err != nil {
		fmt.Println("保存" + wxName + "文章失败", err.Error())
	}
}