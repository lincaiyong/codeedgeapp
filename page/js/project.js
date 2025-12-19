function project_collapseAll() {
    tree_collapseAll(g.root.treeEle);
}

function project_expandAll() {
    tree_expandAll(g.root.treeEle);
}

function project_locateItem() {
    tree_locate(g.root.treeEle, g.root.currentFilePath);
}

function project_openFile({filePath, lineNumber, selection, patch, rhs}) {
    g.root.currentPatch = '';
    g.root.currentFilePath = filePath;
    let project = g.root.project;
    let relPath = filePath;
    if (relPath.startsWith('@vendor/')) {
        const items = filePath.split('/');
        project = items[1];
        relPath = items.slice(2).join('/');
    }
    g.fetch(`./file/${relPath}?project=${project}&patch=${patch ? encodeURIComponent(patch) : ''}&rhs=${rhs || ''}`).then(res => {
        if (res instanceof Array) {
            g.root.showCompare = true;
            g.root.compareEle.lhs = res[0];
            g.root.compareEle.rhs = res[1];
            if (res.length === 3) {
                g.root.currentPatch = res[2].trimEnd();
            }
        } else {
            g.root.showCompare = false;
            g.root.currentFileContent = res;
            g.root.currentFileLanguage = relPath.endsWith('.go') ? 'go' : '';
            if (lineNumber) {
                setTimeout(() => {
                    if (selection) {
                        g.root.editorEle.selection = selection;
                    } else {
                        g.root.editorEle.selection = [lineNumber, 1, lineNumber + 1, 1];
                    }
                    g.root.editorEle.focusLine = lineNumber;
                }, 100);
            }
        }
    }).catch((err) => {
        console.error(err);
    });
}

function project_clickItem(itemEle) {
    console.log('click: ' + JSON.stringify(itemEle.data));
    if (itemEle.data.leaf) {
        const filePath = itemEle.data.key;
        project_openFile({filePath});
    }
}

function editor_onCursorChange(lineNo, charNo) {
    // editor_pushHistory(g.root.currentFilePath, lineNo);
}

function editor_copyPath() {
    let v;
    if (g.root.currentPatch) {
        v = `---+++ ${g.root.currentFilePath}\n${g.root.currentPatch}\n---+++\n`;
    } else {
        v = `@${g.root.currentFilePath}:${g.root.editorEle.currentLine}`;
    }
    navigator.clipboard.writeText(v);
    // editor_appendValue(g.root.noteEle.editorEle, `\n${v}`);
}

function editor_addBookmark() {

}