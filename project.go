package codeedgeapp

import . "github.com/lincaiyong/gui"

func ProjectView() *Element {
	return Div(NewOpt().V("root.leftView === 'project'"),
		Div(NewOpt().W("next.x + next.w/2"),
			Div(NewOpt().H("next.y+next.h/2"),
				// header
				Div(NewOpt().H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
					Text(NewOpt().X(".y+4").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Project'"),
					Div(NewOpt(),
						Button(NewButtonOpt().X("next.x-parent.h+.y").Y("next.y").Svg(SvgLocate).OnClick("project_locateItem")),
						Button(NewButtonOpt().X("next.x-parent.h+.y").Y("next.y").Svg(SvgExpandAll).OnClick("project_expandAll")),
						Button(NewButtonOpt().X("next.x-parent.h+.y").Y("next.y").Svg(SvgCollapseAll).OnClick("project_collapseAll")),
						Button(NewButtonOpt().X("parent.w-parent.h+.y").Y("parent.h/2-.h/2").Svg(SvgHide).OnClick("() => g.root.leftView = ''")),
					),
				),
				// sub header
				Div(NewOpt().Y("prev.y2").H("24"),
					Svg(NewOpt().X(".y+10").Y("parent.h/2-.h/2").W("16").H(".w").Color(ColorGray110), SvgProjectDirectory),
					Text(NewOpt().X("prev.x2+4").Y("parent.h/2-.h/2").H("20"), "root.projectName"),
				),
				// tree
				Named("tree", Tree(NewTreeOpt().Y("prev.y2").H("parent.h-.y").OnClickItem("project_clickItem").Items("root.projectFiles").Indent("16"))),
			),
			HBar(NewBarOpt().Y("parent.h/2").Opacity("1")),
			Div(NewOpt().Y("prev.y+prev.h/2").H("parent.h-.y"),
				// header
				Div(NewOpt().H("33").BgColor(ColorGray247),
					Text(NewOpt().X(".y+4").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Search'"),
				),
				// sub header
				Div(NewOpt().Y("prev.y2").H("33").BorderBottom(1).BorderColor(ColorGray235),
					Named("searchInput", Input(NewInputOpt().W("parent.w-2*.x").H("20").Y("parent.h/2-.h/2").X(".y").OnKeyDown("search_onKeyDown"), "'search...'")),
					Named("searchCaseBtn", SearchToolButton(NewButtonOpt().X("next.x-parent.h+4+.y").Y("next.y").Svg(SvgMatchCase))),
					Named("searchWordBtn", SearchToolButton(NewButtonOpt().X("next.x-parent.h+4+.y").Y("next.y").Svg(SvgExactWords))),
					Named("searchRegexBtn", SearchToolButton(NewButtonOpt().X("parent.w-parent.h+4").Y("parent.h/2-.w/2").Svg(SvgRegex))),
				),
				// tree
				Named("searchTree", Tree(NewTreeOpt().Y("prev.y2").H("parent.h-.y").Items("root.searchResults").OnClickItem("search_clickItem").OnComputeIcon("search_onComputeIcon"))),
			),
		),
		Named("leftViewBar", VBar(NewBarOpt().BgColor(ColorYellow).Opacity("0").X("parent.w/3").V("prev.v"))),
		Named("mainView", Div(NewOpt().X("prev.x+prev.w/2").W("parent.w-.x").V("prev.v").BorderLeft(1).BorderColor(ColorGray235),
			Div(NewOpt().H("33").BorderBottom(1).BorderColor(ColorGray235).BgColor(ColorGray247),
				Text(NewOpt().X("10").H("20").Y("parent.h/2-.h/2"), "root.currentFilePath"),
				Button(NewButtonOpt().X("prev.x2").Y("parent.h/2-.h/2").Svg(SvgCopy).OnClick("editor_copyPath").V("root.currentFilePath !== ''")),
				Button(NewButtonOpt().X("prev.x2").Y("parent.h/2-.h/2").Svg(SvgBookMarks).OnClick("editor_addBookmark").V("prev.v")),
			),
			Div(NewOpt().Y("prev.y2").H("parent.h-.y"),
				Named("editor", Editor(NewEditorOpt().V("!root.showCompare").OnCursorPositionChange("editor_onCursorChange").
					Value("root.currentFileContent").Language("root.currentFileLanguage"))),
				Named("compare", Compare(NewOpt().V("root.showCompare"))),
			),
		)),
	)
}
