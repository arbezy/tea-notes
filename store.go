package main

type Note struct {
	ID int64
	Title string
	Body string
}

type Store struct {

}

func (s *Store) Init() error {
	return nil
}

func (s *Store) GetNotes() ([]Note, error) {
	var notes []Note
	notes = append(notes, Note{1, "title1", "bodytextbodytext"})
	notes = append(notes, Note{2, "title2", "bodytextbodytext"})
	notes = append(notes, Note{3, "title3", "bodytextbodytext"})
	return notes, nil
}

func (s *Store) SaveNote(note Note) error {
	return nil
}

