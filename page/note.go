package page

import . "github.com/lincaiyong/gui"

func NoteView() *Element {
	return Div(NewOpt().V("root.rightView === 'note'"),
		Div(NewOpt().H("33").BgColor(ColorGray247).BorderBottom(1).BorderColor(ColorGray235),
			Text(NewOpt().X("10").Y("parent.h/2-.h/2").H("20").FontWeight("500"), "'Note'"),
			Button(NewButtonOpt().X("next.x-.w").Y("next.y").Svg(SvgSave).OnClick("note_onSave").V("root.admin")),
			Button(NewButtonOpt().X("parent.w-parent.h+.y").Y("parent.h/2-.h/2").Svg(SvgHide).OnClick("() => g.root.rightView = ''")),
		),
		Div(NewOpt().Y("prev.y2").H("parent.h-.y"),
			Named("note", Note(NewNoteOpt().OnOpenLink("note_onOpenLink"))),
		),
	)
}
