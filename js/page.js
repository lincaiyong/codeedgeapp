function root_onCreated() {
    setTimeout(function () {
        root_showInfo("welcome");
        g.root.leftView = 'project';
        g.root.rightView = 'chat';
        g.root.bottomView = 'data';
        g.root.dataEle.fields = ['id', 'project', 'note', 'patch'];
        g.root.dataEle.data = [['1', 'xx', '@data.go:10\n', ''], ['2', 'xx', '@run.go:20', '']];
        g.fetch('./files').then(resp => {
            g.root.projectFiles = JSON.parse(resp);
        }).catch(err => {
            console.error(err);
        })
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

function root_onUpdated(ele, k) {
    if (k === 'message') {

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

function root_openProject(key) {
    g.root.leftView = 'project';
}

function root_openNote(text) {
    g.root.rightView = 'note';
    g.root.noteContent = text;
}

function root_test() {
    g.event('/test', (data, reason) => {
        if (reason === 'close' || reason === 'error') {
            console.log('done');
        } else {
            console.log(data);
        }
    });
}