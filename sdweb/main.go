package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sd"
	"strconv"
	"time"
)

var (
	f = new(sd.Form)
)

func main() {
	http.HandleFunc("/", handler(get_index))
	http.HandleFunc("/answer", handler(get_sd))
	http.ListenAndServe(":8080", nil)
}

func get_index(w http.ResponseWriter, r *http.Request, tmpl template.Template) {
	tmpl.Execute(w, f.WriteList(false))
	fmt.Println(time.Now().String() + "  初始")
}

func get_sd(w http.ResponseWriter, r *http.Request, tmpl template.Template) {
	r.ParseForm()
	var buf = sd.FList{}
	for i, _ := range buf {
		figure, _ := strconv.Atoi(r.FormValue("Box" + strconv.Itoa(i)))
		buf[i] = uint8(figure)
	}
	fmt.Println(buf)
	f.InitShudu(sltos(buf))
	if f.CheckSd() {
		if !f.Answer() {
			fmt.Println(time.Now().String() + "  无解")
		}
	} else {
		fmt.Println(time.Now().String() + "数独题错误")
	}
	fmt.Println(f.WriteList(true))
	tmpl.Execute(w, f.WriteList(true))
	fmt.Println(time.Now().String() + "  解答")
}

func handler(fn func(http.ResponseWriter, *http.Request, template.Template)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r, writectx(readtpl("index.tpl")))
	}
}

func readtpl(filename string) string {
	readbuf, _ := ioutil.ReadFile(filename)
	return string(readbuf)
}

func writectx(tpl string) template.Template {
	tmpl := template.New("")
	//template.Must(t, err)
	tmpl.Parse(tpl)
	return *tmpl
}

func sltos(buf sd.FList) string {
	s := ""
	for _, v := range buf {
		s += string(v + 48)
	}
	return s
}
