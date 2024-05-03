CREATE TABLE list_notes(
  list_id SERIAL,
  note_id SERIAL,
  checked BOOLEAN NOT NULL DEFAULT false,
  CONSTRAINT fk_list FOREIGN KEY(list_id) REFERENCES lists(id),
  CONSTRAINT fk_note FOREIGN KEY(note_id) REFERENCES notes(id)
);