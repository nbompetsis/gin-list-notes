CREATE TABLE list_notes(
  listId INTEGER,
  noteId INTEGER,
  checked BOOLEAN,
  CONSTRAINT fk_list FOREIGN KEY(listId) REFERENCES lists(id),
  CONSTRAINT fk_note FOREIGN KEY(noteId) REFERENCES notes(id)
);