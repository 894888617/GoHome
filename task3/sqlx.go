package task3

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Employees struct {
	ID         uint64    `db:"id"`
	Name       string    `db:"name"`
	Department string    `db:"department"`
	Salary     int       `db:"salary"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type Books struct {
	ID        uint64    `db:"id"`
	Title     string    `db:"title"`
	Author    string    `db:"author"`
	Price     int       `db:"price"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func RunSqlx(db *sqlx.DB) {

	var books []Books
	err := db.Select(&books, "select * from books where price > ? order by price desc ", 50)
	if err != nil {
		panic(err)
	}
	fmt.Printf("找到 %d 本价格大于50元的书籍", len(books))

	for _, book := range books {
		fmt.Println(book)
	}

	return
}
