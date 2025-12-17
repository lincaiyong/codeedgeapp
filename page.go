package codeedgeapp

import (
	"embed"
	_ "embed"
	"github.com/gin-gonic/gin"
	. "github.com/lincaiyong/gui"
)

//go:embed js/*.js
var pageJsFS embed.FS

func handlePage(c *gin.Context) {
	root := Div(NewOpt().OnCreated("root_onCreated"),
		Named("bottomBarElse", Div(NewOpt().H("parent.h-next.h"),
			Named("leftBar", LeftBar(NewOpt().W("33").BgColor(ColorGray247).BorderRight(1).BorderColor(ColorGray235))),
			Named("leftRightBarElse", Div(NewOpt().X("prev.x2").W("parent.w-prev.w-next.w"),
				Named("bottomViewElse", Div(NewOpt().H("next.y+next.h/2").V("!!root.leftView || !!root.rightView"),
					Named("leftView", Div(NewOpt().W("next.x+next.w/2").V("!!root.leftView").BorderColor(ColorGray235).BorderRight(1),
						ProjectView(),
					)),
					Named("rightViewBar", VBar(NewBarOpt().X("prev.v ? (next.v ? parent.w*3/4 : parent.w-.w/2) : -.w/2").V("next.v && prev.v"))),
					Named("rightView", Div(NewOpt().X("prev.x+prev.w/2").W("parent.w-.x").V("!!root.rightView").BgColor(ColorGray247).BorderLeft(1).BorderColor(ColorGray235),
						NotebookView(),
					)),
				)),
				Named("bottomViewBar", HBar(NewBarOpt().Y("prev.v ? (next.v ? parent.h*3/5 : parent.h-.h/2) : -.h/2").V("next.v && prev.v"))),
				Named("bottomView", Div(NewOpt().Y("prev.y2-prev.h/2").H("parent.h-.y").V("!!root.bottomView").BorderTop(1).BorderColor(ColorGray235),
					DataView(),
				)),
			)),
			Named("rightBar", RightBar(NewOpt().X("parent.w-.w").W("33").BgColor(ColorGray247).BorderColor(ColorGray235).BorderLeft(1))),
		)),
		Named("bottomBar", BottomBar(NewOpt().Y("parent.h-.h").H("24").BgColor(ColorGray247).BorderColor(ColorGray235).BorderTop(1))),
	)
	// view
	root.SetProperty("leftView", "''")
	root.SetProperty("bottomView", "''")
	root.SetProperty("rightView", "''")
	// project
	root.SetProperty("projectFiles", "[]")
	root.SetProperty("projectName", "'?'")
	// search
	root.SetProperty("searchInputText", "''")
	root.SetProperty("searchResults", "[]")
	root.SetProperty("searchResultMap", "{}")
	// editor
	root.SetProperty("showCompare", "false")
	root.SetProperty("currentFilePath", "''")
	root.SetProperty("currentFileContent", "''")
	root.SetProperty("currentFileLanguage", "'go'")
	//
	root.SetProperty("message", "''")
	// js
	var jsCode []string
	items, _ := pageJsFS.ReadDir("js")
	for _, item := range items {
		b, _ := pageJsFS.ReadFile("js/" + item.Name())
		jsCode = append(jsCode, string(b))
	}
	HandlePage(c, "CodeEdge App", root, jsCode...)
}
