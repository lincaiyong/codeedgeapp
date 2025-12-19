function search_onKeyDown(ele, event) {
    if (event.key === 'Enter') {
        search_doSearch();
    }
}

function search_doSearchProject(text, flags, project) {
    return new Promise((resolve, reject) => {
        g.fetch(`./search/?text=${text}&flag=${flags.join(' ')}&project=${project}`).then(resp => {
            resolve(resp);
        }).catch(err => {
            reject(err);
        });
    })
}

async function search_doSearch() {
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
    let resp = await search_doSearchProject(text, flags, g.root.project);
    if (g.root.vendor) {
        for (const v of g.root.vendor.split(',')) {
            const tmp = await search_doSearchProject(text, flags, v);
            tmp.forEach(t => t.path = `@vendor/${v.replaceAll('/', '%2f')}/${t.path}`);
            resp = resp.concat(tmp);
        }
    }
    const results = [];
    const resultMap = {};
    for (const item of resp) {
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
}

function search_clickItem(ele) {
    if (ele.data.leaf) {
        const {path, line_number, match_index} = g.root.searchResultMap[ele.data.key];
        project_openFile({
            filePath: path,
            lineNumber: line_number,
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
