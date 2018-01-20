package main

import (
	"github.com/StevenZack/naif"
	. "github.com/lxn/walk/declarative"
)

func main() {
	naif.Start("./views/")
	MainWindow{
		Title:   "Test",
		Icon:    Bind("'./music.ico'"),
		MinSize: Size{400, 400},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			WebView{
				Name: "wv",
				URL:  "http://127.0.0.1:10246",
			},
		},
	}.Run()
}
