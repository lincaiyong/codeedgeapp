function project_collapseAll() {
    tree_collapseAll(g.root.treeEle);
}

function project_expandAll() {
    tree_expandAll(g.root.treeEle);
}

function project_locateItem() {
    tree_locate(g.root.treeEle, g.root.currentFilePath);
}

function project_gotoLine(lineNumber, selection) {
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

function jsonToMarkdown(json) {
    let markdown = '';

    function flattenJson(obj, prefix = '') {
        const result = [];
        if (Array.isArray(obj)) {
            obj.forEach((item, index) => {
                const path = `${prefix}[${index}]`;
                if (item !== null && typeof item === 'object' && !Array.isArray(item)) {
                    result.push(...flattenJson(item, path));
                } else if (Array.isArray(item)) {
                    result.push(...flattenJson(item, path));
                } else if (typeof item === 'string' && item.includes('\n')) {
                    result.push({path, value: item});
                }
            });
            return result;
        }
        for (const [key, value] of Object.entries(obj)) {
            const path = prefix ? `${prefix}/${key}` : key;
            if (Array.isArray(value)) {
                result.push(...flattenJson(value, path));
            } else if (value !== null && typeof value === 'object') {
                result.push(...flattenJson(value, path));
            } else if (typeof value === 'string' && value.includes('\n')) {
                result.push({path, value});
            }
        }
        return result;
    }

    const flatData = flattenJson(json);
    for (const {path, value} of flatData) {
        const pathParts = path.split('/');
        let depth = 0;
        for (const part of pathParts) {
            depth++;
            const arrayDepth = (part.match(/\[/g) || []).length;
            depth += arrayDepth;
        }
        markdown += `${path}\n`;
        markdown += '```go\n' + value + '\n```\n\n';
    }

    return markdown.trim();
}


function project_fmtFileContent(text) {
    try {
        const obj = JSON.parse(text);
        text = jsonToMarkdown(obj, {
            level: 1,
            maxLevel: 4,
            arrayAsList: true,
            codeBlock: true,
        })
        return text;
    } catch (e) {
        console.error(e);
        return text;
    }
}

function project_openFile({filePath, lineNumber, selection, patch, rhs}) {
    g.root.currentPatch = '';
    g.root.currentFilePath = filePath;
    let project = g.root.project;
    let relPath = filePath;
    if (relPath.endsWith('}')) {
        const idx = relPath.indexOf('{');
        let extra = relPath.substring(idx + 1, relPath.length - 1);
        let format = false;
        if (extra.endsWith('-fmt')) {
            format = true;
            extra = extra.substring(0, extra.length - 4);
        }
        fetch(`./object/${extra}/`).then(async resp => {
            g.root.showCompare = false;
            let text = await resp.text();
            if (format) {
                text = project_fmtFileContent(text);
                g.root.currentFileLanguage = 'markdown';
            } else {
                g.root.currentFileLanguage = relPath.endsWith('.go') ? 'go' : '';
            }
            g.root.currentFileContent = text;
            project_gotoLine(lineNumber, selection);
        }).catch(err => {
            console.error(err);
        });
        return;
    }
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
            project_gotoLine(lineNumber, selection);
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
    let v = `@${g.root.currentFilePath}:${g.root.editorEle.currentLine}`;
    navigator.clipboard.writeText(v);
}


function editor_copyDiff() {
    let v = `---+++ ${g.root.currentFilePath}\n${g.root.currentPatch}\n---+++\n`;
    navigator.clipboard.writeText(v);
}

function editor_addBookmark() {

}