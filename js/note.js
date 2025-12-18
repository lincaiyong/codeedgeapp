function note_onSave() {
    fetch('./note/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            note: g.root.noteEle.content,
            id: g.root.id,
        })
    }).then(() => {
        root_showInfo('note saved');
        data_onRefresh(g.root.dataEle);
    }).catch(error => console.error('Error:', error));
}