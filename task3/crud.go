package task3

import (
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	ID    uint64
	Name  string
	Age   int
	Grade string
}

func Run(db *gorm.DB) {
	db.AutoMigrate(Student{})

	s := &Student{
		Name:  "tet",
		Age:   1,
		Grade: "三年级",
	}

	db.Create(s)

	var student []Student
	db.Debug().Find(&student, "age <= ?", s.Age)

	fmt.Println(student)

	db.Model(&Student{}).Where("name", "张三").Update("grade", "四年级")

	db.Debug().Find(&student)
	fmt.Println(student)

	db.Debug().Where("age <= ?", 15).Delete(&Student{})

	db.Debug().Find(&student)
	fmt.Println(student)

}
