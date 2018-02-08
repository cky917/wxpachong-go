package routes
import(
	"fmt"
	"bufio"
	"encoding/json"
	"time"
	"io/ioutil"
	"os"
	"wx-spider/spider.v1/models"
)

func(s *Server) savePostToLocal(wxReader *WxReader) (err error) {
	savePostDir := "./data/" + wxReader.WxId + ".json"
	file, err := os.Create(savePostDir)
	if err != nil {
		return
	}
	defer file.Close()
	data, _:= json.Marshal(wxReader)
	saveMap := &models.File{
		UpdateTime: fmt.Sprintf("%d",time.Now().Unix()), 
		Result: string(data),
	}
	// 调用 `Sync` 来将缓冲区的信息写入磁盘。
	file.Sync()
	// `bufio` 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(file)
	writeData, _:= json.Marshal(saveMap)
	bytes, err := w.WriteString(string(writeData))
	fmt.Printf("wrote %d bytes\n", bytes)

	// 使用 `Flush` 来确保所有缓存的操作已写入底层写入器。
	w.Flush()
	return nil
}

func(s *Server) readLocalPostById(wxId string) (posts map[string]interface{}, err error){
	postDir := "./data/" + wxId + ".json"
	file, err := os.Open(postDir)
	if err != nil {
		fmt.Println(err)
		return posts, err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	if err = json.Unmarshal(bytes, &posts); err != nil {
		return
	}
	return
}