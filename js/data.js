function data_onError(msg) {
    root_showError(msg);
}

function data_onInfo(msg) {
    root_showInfo(msg);
}

function data_onRefresh(tableEle) {
    g.fetch('./data').then(resp => {
        const {fields, data} = JSON.parse(resp);
        tableEle.fields = fields;
        tableEle.data = data;
    }).catch(e => {
        console.error(e);
    });
}

function data_onOpen(tableEle, id) {
    console.log(tableEle, id);
    root_openProject('??');
    root_openNote('@data.go:10\n---\n---');
}