package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"go-course-workspace/library-app/library"
	"go-course-workspace/library-app/notifications"
	"go-course-workspace/library-app/storage"
)

const dbFile = "books.json"

func main() {
	lib, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не найден, создана новая библиотека.")
			lib = library.New()
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Библиотека загружена.")
	}
	run(lib, dbFile)
}

func run(lib *library.Library, dbPath string) {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nКоманды: add-book | add-reader | list | issue | return | save | load | notify | exit")
		fmt.Print("> ")
		if !sc.Scan() {
			return
		}
		parts := strings.Fields(sc.Text())
		if len(parts) == 0 {
			continue
		}
		switch parts[0] {
		case "add-book":
			lib.AddBook("Книга "+strconv.Itoa(len(lib.Books)+1), "Автор", 2024)
			fmt.Println("OK")
		case "add-reader":
			lib.AddReader("Иван", "Иванов")
			fmt.Println("OK")
		case "list":
			lib.ListAllBooks()
			lib.ListAllReaders()
		case "issue":
			if len(parts) < 3 {
				continue
			}
			b, _ := strconv.Atoi(parts[1])
			r, _ := strconv.Atoi(parts[2])
			if err := lib.IssueBookToReader(b, r); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "return":
			if len(parts) < 2 {
				continue
			}
			b, _ := strconv.Atoi(parts[1])
			if err := lib.ReturnBook(b); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "save":
			if len(parts) < 2 {
				continue
			}
			if err := storage.SaveBooksToCSV(parts[1], lib.Books); err != nil {
				fmt.Println("Ошибка:", err)
			}
		case "load":
			if len(parts) < 2 {
				continue
			}
			books, err := storage.LoadBooksFromCSV(parts[1])
			if err == nil {
				lib.Books = books
				lib.UpdateIDs()
			}
		case "notify":
			for _, n := range []notifications.Notifier{
				notifications.EmailNotifier{EmailAddress: "student@example.com"},
				notifications.SMSNotifier{PhoneNumber: "+79991234567"},
			} {
				n.Notify("Ваша книга просрочена!")
			}
		case "exit":
			storage.SaveLibraryToJSON(dbPath, lib)
			return
		}
	}
}