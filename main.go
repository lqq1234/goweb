package main
import (
	"fmt"
	"net/http"
	"goweb/handler"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conf := handler.GetConfigContent()
		fmt.Fprintf(w, string(conf))
	})
	handler.Info.Println("服务启动在8888端口")
	handler.Error.Println(http.ListenAndServe(":8888", nil))
}

