CREATE TABLE list_notes(
  id SERIAL PRIMARY KEY,
  checked BOOLEAN,
  listId INTEGER,
  noteId INTEGER,
  CONSTRAINT fk_list FOREIGN KEY(listId) REFERENCES list(id),
  CONSTRAINT fk_note FOREIGN KEY(noteId) REFERENCES note(id)
);