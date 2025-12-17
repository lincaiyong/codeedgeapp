function leftBar_onClickProject() {
    g.root.leftView = g.root.leftView ? '' : 'project';
}

function leftBar_onClickTable() {
    g.root.bottomView = g.root.bottomView ? '' : 'table';
}

function rightbar_onClickNotebook() {
    g.root.rightView = g.root.rightView === 'notebook' ? '' : 'notebook';
}