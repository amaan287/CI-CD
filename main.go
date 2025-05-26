package main
import (
	"net/http"
	"fmt"
)
func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte ("Hello world"))
	})
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000",nil)
}
