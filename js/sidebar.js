function leftBar_onClickProject() {
    g.root.leftView = g.root.leftView === 'project' ? '' : 'project';
}

function leftBar_onClickData() {
    g.root.bottomView = g.root.bottomView === 'data' ? '' : 'data';
}

function rightbar_onClickNote() {
    g.root.rightView = g.root.rightView === 'note' ? '' : 'note';
}

function rightbar_onClickChat() {
    g.root.rightView = g.root.rightView === 'chat' ? '' : 'chat';
}