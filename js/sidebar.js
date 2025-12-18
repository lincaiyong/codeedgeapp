function sidebar_onClickProject() {
    g.root.leftView = g.root.leftView === 'project' ? '' : 'project';
}

function sidebar_onClickSearch() {
    g.root.bottomView = g.root.bottomView === 'search' ? '' : 'search';
    if (g.root.bottomView === 'search') {
        g.root.searchInputEle.ref.focus();
    }
}

function sidebar_onClickData() {
    g.root.bottomView = g.root.bottomView === 'data' ? '' : 'data';
}

function sidebar_onClickNote() {
    g.root.rightView = g.root.rightView === 'note' ? '' : 'note';
}

function sidebar_onClickChat() {
    g.root.rightView = g.root.rightView === 'chat' ? '' : 'chat';
}