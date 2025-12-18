package codeedgeapp

import . "github.com/lincaiyong/gui"

func LeftBar(opt *Opt) *Element {
	return Div(opt,
		Named("projectBtn", ToolButton(NewButtonOpt().Svg(SvgProject).X("parent.w/2-.w/2-0.5").Y(".x").Selected("root.leftView === 'project'").OnClick("leftBar_onClickProject"))),
		ToolButton(NewButtonOpt().Svg(SvgSearch).X("prev.x").Y("prev.y2 + 8").Selected("root.showSearch").OnClick("leftBar_onClickSearch")),
		//ToolButton(NewButtonOpt().Svg(SvgPullRequests).X("prev.x").Y("prev.y2 + 8")),
		//HDivider(NewOpt().X("prev.x").Y("prev.y2 + 9").W("prev.w").BgColor(ColorGray201)),
		//ToolButton(NewButtonOpt().Svg(SvgStructure).X("prev.x").Y("prev.y2 + 9")),
		//ToolButton(NewButtonOpt().Svg(SvgMoreHorizontal).X("prev.x").Y("prev.y2 + 8")),

		//ToolButton(NewButtonOpt().Svg(SvgPythonPackages).X("next.x").Y("next.y-8-.h")),
		//ToolButton(NewButtonOpt().Svg(SvgServices).X("next.x").Y("next.y-8-.h")),
		//ToolButton(NewButtonOpt().Svg(SvgTerminal).X("next.x").Y("next.y-8-.h")),
		//ToolButton(NewButtonOpt().Svg(SvgProblems).X("next.x").Y("next.y-8-.h")),
		ToolButton(NewButtonOpt().Svg(SvgImportDataCell).X("parent.w/2-.w/2-0.5").Y("parent.h-.h-.x").Selected("root.bottomView === 'data'").OnClick("leftBar_onClickData")),
	)
}

func RightBar(opt *Opt) *Element {
	return Div(opt,
		ToolButton(NewButtonOpt().Svg(SvgChangedFile).X("parent.w/2-.w/2-0.5").Y(".x").Flag("false").Selected("root.rightView == 'note'").OnClick("rightbar_onClickNote")),
		ToolButton(NewButtonOpt().Svg(SvgAIChat).X("prev.x").Y("prev.y2 + 8").Selected("root.rightView == 'chat'").OnClick("rightbar_onClickChat")),
		//ToolButton(NewButtonOpt().Svg(SvgDatabase).X("prev.x").Y("prev.y2 + 8")),
		//ToolButton(NewButtonOpt().Svg(SvgBookMarks).X("prev.x").Y("prev.y2 + 8").Selected("root.rightView === 'bookmark'").OnClick("rightBar_clickBookmark")),
	)
}

func BottomBar(opt *Opt) *Element {
	return Div(opt, Named("message", Text(NewOpt().X("10").H("18").Y("parent.h/2-.h/2"), "root.message")))
	//SourceRootButton(NewButtonOpt().Text("'page'").X("9").Y("1")),
	//Svg(NewOpt().X("prev.x2").Y("parent.h/2-.h/2-1").W("17").H(".w").Color(ColorGray110), SvgArrowRight),
	//SourceDirButton(NewButtonOpt().Text("'example'").X("prev.x2 + 1").Y("1")),
	//Svg(NewOpt().X("prev.x2").Y("parent.h/2-.h/2-1").W("17").H(".w").Color(ColorGray110), SvgArrowRight),
	//SourceFileButton(NewButtonOpt().Text("'goland.go'").X("prev.x2 + 1").Y("1")),
	//Named("img", Img(NewOpt().V("0"), "'res/img/bot.png'"))
}
