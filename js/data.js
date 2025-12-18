function data_onError(msg) {
    root_showError(msg);
}

function data_onInfo(msg) {
    root_showInfo(msg);
}

function data_onRefresh(dataEle) {
    g.fetch('./data/').then(resp => {
        const {fields, data} = JSON.parse(resp);
        dataEle.fields = fields;
        dataEle.data = data;
    }).catch(e => {
        console.error(e);
    });
}

function data_onOpen(dataEle, id) {
    console.log(dataEle, id);
    const data = dataEle.dataByKey[id];
    if (data) {
        g.root.id = data.id;
        root_openProject(data.project, data.vendor);
        root_openNote(data.note);
    }
}