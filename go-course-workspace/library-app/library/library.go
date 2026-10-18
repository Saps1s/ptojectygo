package library

import (
	"fmt"

	"go-course-workspace/library-app/domain"
)

type Library struct {
	Books        []*domain.Book
	Readers      []*domain.Reader
	lastBookID   int
	lastReaderID int
}

func New() *Library {
	return &Library{Books: []*domain.Book{}, Readers: []*domain.Reader{}}
}

func (lib *Library) AddBook(title, author string, year int) error {
	for _, b := range lib.Books {
		if b.Title == title && b.Author == author {
			return fmt.Errorf("книга %q уже есть", title)
		}
	}
	lib.lastBookID++
	lib.Books = append(lib.Books, &domain.Book{
		ID: lib.lastBookID, Title: title, Author: author, Year: year,
	})
	return nil
}

func (lib *Library) AddReader(firstName, lastName string) *domain.Reader {
	lib.lastReaderID++
	r := &domain.Reader{
		ID: lib.lastReaderID, FirstName: firstName, LastName: lastName, IsActive: true,
	}
	lib.Readers = append(lib.Readers, r)
	return r
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, b := range lib.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}
	if err := book.IssueBook(reader); err != nil {
		return err
	}
	reader.AssignBook(book)
	return nil
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) ListAllBooks() {
	if len(lib.Books) == 0 {
		fmt.Println("Библиотека пуста.")
		return
	}
	for _, b := range lib.Books {
		fmt.Println(b)
	}
}

func (lib *Library) ListAllReaders() {
	for _, r := range lib.Readers {
		fmt.Println(r)
	}
}

func (lib *Library) UpdateIDs() {
	maxBookID := 0
	for _, b := range lib.Books {
		if b.ID > maxBookID {
			maxBookID = b.ID
		}
	}
	lib.lastBookID = maxBookID

	maxReaderID := 0
	for _, r := range lib.Readers {
		if r.ID > maxReaderID {
			maxReaderID = r.ID
		}
	}
	lib.lastReaderID = maxReaderID
}