function project_collapseAll() {
    tree_collapseAll(g.root.treeEle);
}

function project_expandAll() {
    tree_expandAll(g.root.treeEle);
}

function project_locateItem() {
    tree_locate(g.root.treeEle, g.root.currentFilePath);
}

function project_openFile({filePath, lineNumber, selection, patch}) {
    g.root.currentFilePath = filePath;
    let project = g.root.project;
    let relPath = filePath;
    if (relPath.startsWith('@vendor/')) {
        const items = filePath.split('/');
        project = items[1];
        relPath = items.slice(2).join('/');
    }
    g.fetch(`./file/${relPath}?project=${project}&patch=${encodeURIComponent(patch)}`).then(res => {
        if (res instanceof Array) {
            g.root.showCompare = true;
            g.root.compareEle.lhs = res[0];
            g.root.compareEle.rhs = res[1];
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

function search_onKeyDown(ele, event) {
    if (event.key === 'Enter') {
        search_doSearch();
    }
}

function search_doSearch() {
    const text = encodeURIComponent(g.root.searchInputEle.ref.value);
    if (!text) {
        return;
    }
    root_showInfo(`search: ${text}`)
    const flags = [];
    if (!g.root.searchCaseBtnEle.selected) {
        flags.push('-i')
    }
    if (g.root.searchWordBtnEle.selected) {
        flags.push('-w')
    }
    if (!g.root.searchRegexBtnEle.selected) {
        flags.push('-F')
    }
    g.fetch(`./search/?text=${text}&flag=${flags.join(' ')}&project=${g.root.project}&vendor=${g.root.vendor}`).then(resp => {
        const obj = JSON.parse(resp);
        const results = [];
        const resultMap = {};
        for (const item of obj) {
            const {path, line_text, line_number, match_index} = item;
            const startIdx = match_index[0];
            let line = line_text;
            if (startIdx > 25) {
                line = `...${line_text.substring(startIdx - 20)}`;
            }
            const key = `${path.replaceAll('/', '%2f')}/L${line_number}: ${line.replaceAll('/', '%2f')}`;
            results.push(key);
            resultMap[key] = {path, line_number, match_index};
        }
        g.root.searchResults = results;
        g.root.searchResultMap = resultMap;
        tree_expandAll(g.root.searchTreeEle);
    }).catch(err => {
        console.error(err);
    });
}

function search_clickItem(ele) {
    if (ele.data.leaf) {
        const {path, line_number, match_index} = g.root.searchResultMap[ele.data.key];
        project_openFile({
            path,
            line_number,
            selection: [line_number, match_index[0] + 1, line_number, match_index[1] + 1]
        });
    }
}

function search_onComputeIcon(ele) {
    const {leaf, key} = ele.data;
    if (leaf === undefined || key === undefined) {
        return;
    }
    if (!leaf) {
        const ext = key.substring(key.lastIndexOf('.') + 1);
        switch (ext) {
            case 'go':
                return 'res/svg/go.svg';
            case 'js':
                return 'res/svg/js.svg';
            case 'py':
                return 'res/svg/python.svg';
        }
        if (key === 'go.mod' || key.endsWith('/go.mod')) {
            return 'res/svg/goMod.svg'
        }
        if (key === '.gitignore' || key.endsWith('/.gitignore')) {
            return 'res/svg/ignored.svg'
        }
        return 'res/svg/text.svg';
    } else {
        return 'res/svg/textArea.svg';
    }
}

function editor_onCursorChange(lineNo, charNo) {
    // editor_pushHistory(g.root.currentFilePath, lineNo);
}

function editor_copyPath() {
    const v = `@${g.root.currentFilePath}:${g.root.editorEle.currentLine}`;
    navigator.clipboard.writeText(v);
    editor_appendValue(g.root.noteEle.editorEle, `\n${v}`);
}

function editor_addBookmark() {

}