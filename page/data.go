package page

import . "github.com/lincaiyong/gui"

func DataView() *Element {
	return Div(NewOpt().V("root.bottomView === 'data'"),
		Div(NewOpt().H("33").BgColor(ColorGray247),
			Text(NewOpt().X(".y+4").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Data'"),
			Button(NewButtonOpt().X("parent.w-parent.h+.y").Y("parent.h/2-.h/2").Svg(SvgHide).OnClick("() => g.root.bottomView = ''")),
		),
		Div(NewOpt().Y("prev.y2").H("parent.h-prev.h"),
			Named("data", Table(NewTableOpt().OnError("data_onError").OnInfo("data_onInfo").OnRefresh("data_onRefresh").OnOpen("data_onOpen"))),
		),
	)
}
