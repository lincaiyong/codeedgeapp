function project_collapseAll() {
    tree_collapseAll(g.root.treeEle);
}

function project_expandAll() {
    tree_expandAll(g.root.treeEle);
}

function project_locateItem() {
    tree_locate(g.root.treeEle, g.root.currentFilePath);
}

function project_clickItem(itemEle) {
    console.log('click: ' + JSON.stringify(itemEle.data));
    if (itemEle.data.leaf) {
        g.root.currentFilePath = itemEle.data.key;
        const relPath = itemEle.data.key;
        g.fetch('./file/' + relPath).then(res => {
            if (res.startsWith('diff:')) {
                res = res.substring(5);
                const data =JSON.parse(res);
                g.root.showCompare = true;
                const diffModel = g.root.compareEle._editor.getModel();
                if (diffModel) {
                    diffModel.original.setValue(data[0]);
                    diffModel.modified.setValue(data[1]);
                }
            } else {
                g.root.showCompare = false;
                g.root.currentFileContent = res;
                if (relPath.endsWith('.go')) {
                    g.root.currentFileLanguage = 'go';
                }
            }
        }).catch((err) => {
            console.error(err);
        });
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
    g.fetch(`./search?text=${text}&flag=${flags.join(' ')}`).then(resp => {
        const obj = JSON.parse(resp);
        const results = [];
        const resultMap = {};
        for (const item of obj) {
            const {path, line_text, line_number, match_index} = item;
            const startIdx = match_index[0];
            let line = line_text;
            if (startIdx > 25) {
                line = `...${line_text.substring(startIdx-20)}`;
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
        console.log(ele.data.key);
        const {path, line_number, match_index} = g.root.searchResultMap[ele.data.key];
        console.log(path, line_number);
        g.root.currentFilePath = path;
        g.fetch('./file/' + path).then(res => {
            g.root.currentFileContent = res;
            if (path.endsWith('.go')) {
                g.root.currentFileLanguage = 'go';
            }
            g.root.showCompare = false;
            setTimeout(() => {
                g.root.editorEle.selection = [line_number, match_index[0]+1, line_number, match_index[1]+1];
                g.root.editorEle.focusLine = line_number;
            }, 100);
        }).catch((err) => {
            console.error(err);
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
    navigator.clipboard.writeText(g.root.currentFilePath);
}

function editor_addBookmark() {

}