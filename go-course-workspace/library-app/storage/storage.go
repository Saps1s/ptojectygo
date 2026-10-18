package storage

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"

	"go-course-workspace/library-app/domain"
	"go-course-workspace/library-app/library"
)

func SaveBooksToCSV(filePath string, books []*domain.Book) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	w.Write([]string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"})
	for _, b := range books {
		rid := ""
		if b.ReaderID != nil {
			rid = strconv.Itoa(*b.ReaderID)
		}
		w.Write([]string{
			strconv.Itoa(b.ID), b.Title, b.Author,
			strconv.Itoa(b.Year), strconv.FormatBool(b.IsIssued), rid,
		})
	}
	return nil
}

func LoadBooksFromCSV(filePath string) ([]*domain.Book, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil || len(records) < 2 {
		return []*domain.Book{}, err
	}
	books := make([]*domain.Book, 0, len(records)-1)
	for _, rec := range records[1:] {
		id, _ := strconv.Atoi(rec[0])
		year, _ := strconv.Atoi(rec[3])
		isIssued, _ := strconv.ParseBool(rec[4])
		var rid *int
		if rec[5] != "" {
			v, _ := strconv.Atoi(rec[5])
			rid = &v
		}
		books = append(books, &domain.Book{
			ID: id, Title: rec[1], Author: rec[2], Year: year,
			IsIssued: isIssued, ReaderID: rid,
		})
	}
	return books, nil
}

type storageData struct {
	Books   []*domain.Book   `json:"books"`
	Readers []*domain.Reader `json:"readers"`
}

func SaveLibraryToJSON(filePath string, lib *library.Library) error {
	data := storageData{Books: lib.Books, Readers: lib.Readers}
	raw, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, raw, 0644)
}

func LoadLibraryFromJSON(filePath string) (*library.Library, error) {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var data storageData
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, err
	}
	lib := library.New()
	lib.Books = data.Books
	lib.Readers = data.Readers
	lib.UpdateIDs()
	return lib, nil
}