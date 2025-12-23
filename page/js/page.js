function root_onCreated() {
    setTimeout(function () {
        root_showInfo("welcome");
        // g.root.leftView = 'project';
        // g.root.rightView = 'note';
        g.root.bottomView = 'data';
        data_onRefresh(g.root.dataEle);
    });
    if (g.root.admin) {
        let serverPid = 0;
        setInterval(() => g.fetch('./status/').then(resp => {
            const {pid} = resp || {};
            if (serverPid && pid && pid !== serverPid) {
                location.reload();
            }
            serverPid = pid;
        }).catch(e => {
            console.error(e);
        }), 1000);

        document.addEventListener('keydown', function (e) {
            if ((e.metaKey || e.ctrlKey) && e.key === 's') {
                e.preventDefault();
                note_onSave();
            }
        });

        document.addEventListener('keydown', function (e) {
            if ((e.metaKey || e.ctrlKey) && e.shiftKey && (e.key === 'x' || e.key === 'd')) {
                e.preventDefault();
                if (g.root.currentFilePath) {
                    const editor = g.root.editorEle?._editor;
                    if (editor) {
                        const selectedText = editor.getModel().getValueInRange(editor.getSelection());
                        const v = `${selectedText}@${g.root.currentFilePath}:${g.root.editorEle.currentLine}`;
                        if (e.key === 'x') {
                            editor_appendValue(g.root.noteEle.editorEle, `\n${v}`);
                        } else {
                            navigator.clipboard.writeText(v);
                        }
                    }
                }
            }
        });
    }
    document.addEventListener('keydown', function (e) {
        if ((e.metaKey || e.ctrlKey) && e.key === 'f') {
            e.preventDefault();
            g.root.bottomView = 'search';
            g.root.searchInputEle.ref.focus();
        }
    });
    document.addEventListener('keydown', function (e) {
        if ((e.metaKey || e.ctrlKey) && e.key === 'd') {
            e.preventDefault();
            g.root.bottomView = 'data';
        }
    });
}

let root_messageTimer = null;

function root_onUpdated(ele, k, v) {
    if (k === 'message') {
        clearTimeout(root_messageTimer);
        if (v) {
            root_messageTimer = setTimeout(() => {
                ele.message = '';
            }, 2000);
        }
    }
    if (k === 'id') {
        window.document.title = `CodeEdge App ${v}`;
    }
}

function root_showWarn(msg) {
    g.root.message = `⚠️ ${msg}`;
}

function root_showError(msg) {
    g.root.message = `❌ ${msg}`;
}

function root_showInfo(msg) {
    g.root.message = `✅ ${msg}`;
}

function root_fetchProjectFiles(project) {
    return new Promise((resolve, reject) => {
        g.fetch(`./files/?project=${project}`).then(resp => {
            resolve(resp);
        }).catch(e => {
            reject(e);
        });
    })
}

async function root_openProject(project, vendor) {
    g.root.project = project;
    g.root.showCompare = false;
    g.root.currentFilePath = '';
    g.root.currentFileContent = '';
    g.root.currentFileLanguage = '';
    g.root.vendor = vendor;
    g.root.leftView = 'project';
    let files = await root_fetchProjectFiles(project);
    if (vendor) {
        for (const v of vendor.split(',')) {
            if (v.endsWith('}')) {
                files.push(`@vendor/${v}`);
            } else {
                let vendorFiles = await root_fetchProjectFiles(v);
                vendorFiles = vendorFiles.map(f => `@vendor/${v.replaceAll('/', '%2f')}/${f}`);
                files = files.concat(vendorFiles);
            }
        }
    }
    g.root.projectFiles = files;
}

function root_openNote(text) {
    g.root.rightView = 'note';
    g.root.noteEle.content = text;
}
