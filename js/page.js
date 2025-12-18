function root_onCreated() {
    setTimeout(function () {
        root_showInfo("welcome");
        // g.root.leftView = 'project';
        // g.root.rightView = 'note';
        g.root.bottomView = 'data';
        data_onRefresh(g.root.dataEle);
    });
    let serverPid = 0;
    setInterval(() => g.fetch('./status').then(resp => {
        const {pid} = JSON.parse(resp) || {};
        if (serverPid && pid && pid !== serverPid) {
            location.reload();
        }
        serverPid = pid;
    }).catch(e => {
        console.error(e);
    }), 1000);
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

function root_openProject(project, vendor) {
    g.root.project = project;
    g.root.currentFilePath = '';
    g.root.currentFileContent = '';
    g.root.currentFileLanguage = '';
    g.root.vendor = vendor;
    g.root.leftView = 'project';
    g.fetch(`./files/?project=${project}&vendor=${vendor}`).then(resp => {
        g.root.projectFiles = JSON.parse(resp);
    });
}

function root_openNote(text) {
    g.root.rightView = 'note';
    g.root.noteEle.content = text;
}
