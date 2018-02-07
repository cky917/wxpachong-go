
package main

import (
	"flag"

	"wx-spider/spider.v1"
)

func main() {
	var port, path string
	flag.StringVar(&port, "port", "8080", "run on port")
	flag.StringVar(&path, "config", "./x.config", "config location")
	flag.Parse()

	spider.Run(port, path)
}