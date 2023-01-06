package repository

type Library struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Location string
}

// SQL REPO
func (repo *Repository) GetLibraryByID(libraryID int) (Library, error) {
	return Library{}, nil
}

func (repo *Repository) InsertLibrary(libraryID int) error {
	return nil
}

func (repo *Repository) UpdateLibrary(library Library) (Library, error) {
	return Library{}, nil
}

func (repo *Repository) DeleteLibrary(libraryID int) error {
	return nil
}

// Redis REPO
