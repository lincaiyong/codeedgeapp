function note_onSave() {
    let note = g.root.noteEle.content;
    note = note.replace(/@([^@]+)@\n```\n[\s\S]*?\n```/g, '@$1@');
    fetch('./note/', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            note,
            id: g.root.id,
        })
    }).then(() => {
        root_showInfo('note saved');
        data_onRefresh(g.root.dataEle);
    }).catch(error => console.error('Error:', error));
}