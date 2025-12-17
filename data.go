package codeedgeapp

import . "github.com/lincaiyong/gui"

func DataView() *Element {
	return Div(NewOpt(),
		Div(NewOpt().H("33").BgColor(ColorGray247),
			Text(NewOpt().X(".y+4").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Data'"),
		),
		Div(NewOpt().Y("prev.y2"),
			Named("data", Table(NewTableOpt().OnError("data_onError").OnInfo("data_onInfo"))),
		),
	)
}
