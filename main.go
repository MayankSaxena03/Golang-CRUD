package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	Id     int    `json:"Id"`
	Name   string `json:"Name"`
	Author string `json:"Author"`
	Year   int    `json:"Year"`
}

func home(w http.ResponseWriter, r *http.Request) {
	t := template.New("home")
	t, _ = t.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Home</title>
		</head>
		<body>
			<h1>Home</h1>
			<a href="/books">Books</a>
			<a href="/books/create">Create</a>
			<a href="/books/update">Update</a>
			<a href="/books/delete">Delete</a>
		</body>
		</html>
	`)
	t.Execute(w, nil)
}

func books(w http.ResponseWriter, r *http.Request) {
	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatalf("Error opening csv file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}

	var books []Book
	for i, row := range csvData {
		if len(row) != 4 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		bookYear, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("Error converting Year to int at line: %d", i+1)
			continue
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatalf("Error converting Id to int at line: %d", i+1)
			continue
		}
		books = append(books, Book{
			Id:     Id,
			Name:   row[1],
			Author: row[2],
			Year:   bookYear,
		})
	}

	m := map[string]interface{}{
		"books": books,
	}

	t := template.New("books")
	t, _ = t.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Books</title>
		</head>
		<body>
			<h1>Books</h1>
			<table>
				<tr>
					<th>Id</th>
					<th>Name</th>
					<th>Author</th>
					<th>Year</th>
				</tr>
				{{range .books}}
				<tr>
					<td>{{.Id}}</td>
					<td>{{.Name}}</td>
					<td>{{.Author}}</td>
					<td>{{.Year}}</td>
				</tr>
				{{end}}
			</table>
		</body>
		</html>
	`)

	if err := t.Execute(w, m); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func bookInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		log.Fatalf("Error converting Id to int: %v", err)
	}

	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatalf("Error opening csv file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}

	var books []Book
	for i, row := range csvData {
		if len(row) != 4 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		bookYear, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("Error converting Year to int at line: %d", i+1)
			continue
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatalf("Error converting Id to int at line: %d", i+1)
			continue
		}
		books = append(books, Book{
			Id:     Id,
			Name:   row[1],
			Author: row[2],
			Year:   bookYear,
		})
	}

	var book Book
	found := false
	for _, b := range books {
		if b.Id == Id {
			book = b
			found = true
			break
		}
	}

	if !found {
		t := template.New("book not found")
		t, _ = t.Parse(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Book not found</title>
			</head>
			<body>
				<h1>Book not found</h1>
			</body>
			</html>
		`)
		t.Execute(w, nil)
		return
	}

	m := map[string]interface{}{
		"book": book,
	}

	t := template.New("book info")
	t, _ = t.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Book info</title>
		</head>
		<body>
			<h1>Book info</h1>
			<table>
				<tr>
					<th>Id</th>
					<th>Name</th>
					<th>Author</th>
					<th>Year</th>
				</tr>
				<tr>
					<td>{{.book.Id}}</td>
					<td>{{.book.Name}}</td>
					<td>{{.book.Author}}</td>
					<td>{{.book.Year}}</td>
				</tr>
			</table>
		</body>
		</html>
	`)
	if err := t.Execute(w, m); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	Name := r.FormValue("Name")
	Author := r.FormValue("Author")
	Year := r.FormValue("Year")

	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatalf("Error opening csv file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}

	var books []Book
	for i, row := range csvData {
		if len(row) != 4 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		bookYear, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("Error converting Year to int at line: %d", i+1)
			continue
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatalf("Error converting Id to int at line: %d", i+1)
			continue
		}
		books = append(books, Book{
			Id:     Id,
			Name:   row[1],
			Author: row[2],
			Year:   bookYear,
		})

	}

	maxId := 0
	for _, b := range books {
		if b.Id > maxId {
			maxId = b.Id
		}
	}

	bookYear, err := strconv.Atoi(Year)
	if err != nil {
		log.Fatalf("Error converting Year to int: %v", err)
	}

	books = append(books, Book{
		Id:     maxId + 1,
		Name:   Name,
		Author: Author,
		Year:   bookYear,
	})

	file, err := os.Create("books.csv")
	if err != nil {
		log.Fatalf("Error creating csv file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, b := range books {
		err := writer.Write([]string{strconv.Itoa(b.Id), b.Name, b.Author, strconv.Itoa(b.Year)})
		if err != nil {
			log.Fatalf("Error writing to csv file: %v", err)
		}
	}

	t := template.New("book created")
	t, _ = t.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Book Created</title>
		</head>
		<body>
			<h1>Book Created</h1>
		</body>
		</html>
	`)
	t.Execute(w, nil)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		log.Fatalf("Error converting Id to int: %v", err)
	}
	Name := r.FormValue("Name")
	Author := r.FormValue("Author")
	Year := r.FormValue("Year")

	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatalf("Error opening csv file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}

	var books []Book
	for i, row := range csvData {
		if len(row) != 4 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		bookYear, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("Error converting Year to int at line: %d", i+1)
			continue
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatalf("Error converting Id to int at line: %d", i+1)
			continue
		}
		books = append(books, Book{
			Id:     Id,
			Name:   row[1],
			Author: row[2],
			Year:   bookYear,
		})

	}

	for i, b := range books {
		if b.Id == Id {
			bookYear, err := strconv.Atoi(Year)
			if err != nil {
				log.Fatalf("Error converting Year to int: %v", err)
			}
			books[i] = Book{
				Id:     Id,
				Name:   Name,
				Author: Author,
				Year:   bookYear,
			}
		}
	}

	file, err := os.Create("books.csv")
	if err != nil {
		log.Fatalf("Error creating csv file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, b := range books {
		err := writer.Write([]string{strconv.Itoa(b.Id), b.Name, b.Author, strconv.Itoa(b.Year)})
		if err != nil {
			log.Fatalf("Error writing to csv file: %v", err)
		}
	}

	t := template.New("book updated")
	t, _ = t.Parse(`

		<!DOCTYPE html>
		<html>
		<head>
			<title>Book Updated</title>
		</head>
		<body>
			<h1>Book Updated</h1>
		</body>	
		</html>
	`)
	t.Execute(w, nil)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id, err := strconv.Atoi(vars["Id"])
	if err != nil {
		log.Fatalf("Error converting Id to int: %v", err)
	}

	csvFile, err := os.Open("books.csv")
	if err != nil {
		log.Fatalf("Error opening csv file: %v", err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading csv file: %v", err)
	}
	fmt.Println(csvData)

	var books []Book
	for i, row := range csvData {
		if len(row) != 4 {
			log.Fatalf("Error at line: %d", i+1)
			continue
		}

		bookYear, err := strconv.Atoi(row[3])
		if err != nil {
			log.Fatalf("Error converting Year to int at line: %d", i+1)
			continue
		}
		Id, err := strconv.Atoi(row[0])
		if err != nil {
			log.Fatalf("Error converting Id to int at line: %d", i+1)
			continue
		}
		books = append(books, Book{
			Id:     Id,
			Name:   row[1],
			Author: row[2],
			Year:   bookYear,
		})
	}

	for i, b := range books {
		if b.Id == Id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	file, err := os.Create("books.csv")
	if err != nil {
		log.Fatalf("Error creating csv file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, b := range books {
		err := writer.Write([]string{strconv.Itoa(b.Id), b.Name, b.Author, strconv.Itoa(b.Year)})
		if err != nil {
			log.Fatalf("Error writing to csv file: %v", err)
		}
	}

	t := template.New("book deleted")
	t, _ = t.Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>Book Deleted</title>
		</head>
		<body>
			<h1>Book Deleted</h1>
		</body>
		</html>
	`)
	t.Execute(w, nil)
}

func main() {
	fmt.Println("Server started on port 8080")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/books/", books)
	r.HandleFunc("/books/{Id}/", bookInfo)
	r.HandleFunc("/books/create/", createBook)
	r.HandleFunc("/books/update/{Id}", updateBook)
	r.HandleFunc("/books/delete/{Id}", deleteBook)
	http.ListenAndServe(":8080", r)
}

/** URLS for testing **/
// http://localhost:8080/books/
// http://localhost:8080/books/1/
// http://localhost:8080/books/create/?Name=The%20Hobbit&Author=J.R.R.%20Tolkien&Year=1937
// http://localhost:8080/books/update/1?Name=The%20Hobbit&Author=J.R.R.%20Tolkien&Year=1937
// http://localhost:8080/books/delete/1
