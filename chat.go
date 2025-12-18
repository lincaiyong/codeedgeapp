package codeedgeapp

import . "github.com/lincaiyong/gui"

func ChatView() *Element {
	return Div(NewOpt().V("root.rightView === 'chat'"),
		Div(NewOpt().H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
			Text(NewOpt().X("10").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Chat'"),
			Button(NewButtonOpt().X("next.x-parent.h+.y").Y("next.y").Svg(SvgDelete).OnClick("() => g.root.chatEle.outputEle.value = ''")),
			Button(NewButtonOpt().X("parent.w-parent.h+.y").Y("parent.h/2-.h/2").Svg(SvgHide).OnClick("() => g.root.rightView = ''")),
		),
		Div(NewOpt().Y("prev.y2").H("parent.h-.y").BgColor("'white'"),
			Named("chat", Chat(NewChatOpt().EventUrl("'/chat'"))),
		),
	)
}
