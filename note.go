package codeedgeapp

import . "github.com/lincaiyong/gui"

func NotebookView() *Element {
	return Div(NewOpt(),
		Div(NewOpt().H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
			Text(NewOpt().X("10").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Notebook'"),
		),
		Div(NewOpt().Y("prev.y2").H("parent.h-.y"),
			Named("note", Note(NewNoteOpt().Content("root.noteContent"))),
		),
	)
}
