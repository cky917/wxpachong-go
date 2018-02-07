package spider
import(
	"log"
	"wx-spider/spider.v1/config"
	"wx-spider/spider.v1/routes"
	"wx-spider/libs/install.v1"
)

func Run(port string, path string) {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c, err := config.New(path)
	if err != nil {
		log.Fatal(err)
	}
	i := install.New()
	r := install.NewRouter()
	routes.Register(r, c)
	i.WithRouter(r)
	i.ListenAndServe(":" + port)
}