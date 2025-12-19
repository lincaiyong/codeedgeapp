package page

import . "github.com/lincaiyong/gui"

func SearchView() *Element {
	return Div(NewOpt().V("root.bottomView === 'search'"),
		Div(NewOpt(),
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
	)
}
