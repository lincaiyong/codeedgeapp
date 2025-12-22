package page

import . "github.com/lincaiyong/gui"

func ProjectView() *Element {
	return Div(NewOpt().V("root.leftView === 'project'"),
		Div(NewOpt().W("next.x + next.w/2"),
			Div(NewOpt(),
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
				Div(NewOpt().Y("prev.y2").H("24").V("!!root.project"),
					Svg(NewOpt().X(".y+10").Y("parent.h/2-.h/2").W("16").H(".w").Color(ColorGray110), SvgProjectDirectory),
					Text(NewOpt().X("prev.x2+4").Y("parent.h/2-.h/2").H("20"), "root.project"),
				),
				// tree
				Named("tree", Tree(NewTreeOpt().Y("prev.y2").H("parent.h-.y").OnClickItem("project_clickItem").Items("root.projectFiles").Indent("16"))),
			),
		),
		Named("leftViewBar", VBar(NewBarOpt().BgColor(ColorYellow).Opacity("0").X("200").V("prev.v"))),
		Named("mainView", Div(NewOpt().X("prev.x+prev.w/2").W("parent.w-.x").V("prev.v").BorderLeft(1).BorderColor(ColorGray235),
			Div(NewOpt().H("33").BorderBottom(1).BorderColor(ColorGray235).BgColor(ColorGray247),
				Button(NewButtonOpt().X("10").Y("parent.h/2-.h/2").Svg(SvgCopy).OnClick("editor_copyPath").V("root.currentFilePath !== ''")),
				Button(NewButtonOpt().X("prev.x2").Y("prev.y").W(".v ? prev.w : 0").Svg(SvgDiffWithClipboard).OnClick("editor_copyDiff").V("root.currentFilePath && root.currentPatch")),
				EllipsisText(NewOpt().X("prev.x2").H("20").Y("parent.h/2-.h/2"), "root.currentFilePath"),
			),
			Div(NewOpt().Y("prev.y2").H("parent.h-.y"),
				Named("editor", Editor(NewEditorOpt().V("!root.showCompare").OnCursorPositionChange("editor_onCursorChange").
					Value("root.currentFileContent").Language("root.currentFileLanguage"))),
				Named("compare", Compare(NewOpt().V("root.showCompare"))),
			),
		)),
	)
}
