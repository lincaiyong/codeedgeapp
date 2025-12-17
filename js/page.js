function root_onCreated() {
    setTimeout(function () {
        root_showSucceed("welcome");
        g.root.bottomView = 'data';
        g.root.dataEle.fields = ['id', 'name', 'age', 'height'];
        g.root.dataEle.data = [['1', 'andy', '12', '189'],['2', 'bob', '20', '177']];
        g.fetch('./files').then(resp => {
            g.root.projectFiles = JSON.parse(resp);
        }).catch(err => {
            console.error(err);
        })
    });
}

function root_showWarn(msg) {
    g.root.message = `⚠️ ${msg}`;
}

function root_showSucceed(msg) {
    g.root.message = `✅ ${msg}`;
}

function root_showError(msg) {
    g.root.message = `❌ ${msg}`;
}

function root_showInfo(msg) {
    g.root.message = `${msg}`;
}