package task3

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model          //内置模型(包含ID，CreatedAt，UpdatedAt，DeletedAt)
	Username     string `gorm:"size:50;unique;not null"`
	Password     string `gorm:"size:100;not null"`
	ArticleCount int    `gorm:"default: 0"`
	//一个用户发布的所有文章
	Post []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"size:255;unique;not null"` //标题
	Content string `gorm:"type:text;not null"`       //内容
	Status  int    `gorm:"default:1;not null"`       //状态(1草稿/2发布)
	UserID  uint   `gorm:"not null"`
	User    User   `gorm:"foreignKey:UserID"`

	Comment       []Comment `gorm:"foreignKey:PostID"`
	CommentCount  int       `gorm:"default: 0"`  //新增：统计文章的评论数量
	CommentStatus string    `gorm:"default:无评论"` //新增：评论状态
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null"` //评论内容
	Status  int    `gorm:"default:1;not null"` //状态(1待审核/2已发布)

	PostID uint `gorm:"not null"`
	Post   Post `gorm:"foreignKey:PostID"`

	UserID uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserID"`
}

/*
*
文章创建后自动更新用户的ArticleCount字段
*/
func (p *Post) AfterCreate(tx *gorm.DB) error {
	fmt.Println("Post AfterCreate >>>")
	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&User{}).
			Where("id = ?", p.UserID).
			Update("article_count", gorm.Expr("article_count + ?", 1)).
			Error
		if err != nil {
			panic(err)
			return err
		}
		return nil
	})
}

/*
*
删除评论之前更新文章的commentCount和commentStatus字段
*/
func (c *Comment) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("Comment BeforeDelete >>>")
	return tx.Transaction(func(tx *gorm.DB) error {
		var post Post
		err := tx.First(&post, c.PostID).Error
		if err != nil {
			return err
		}

		newCommentCount := post.CommentCount - 1
		if newCommentCount < 0 {
			newCommentCount = 0
		}

		err = tx.Model(&post).Update("comment_count", newCommentCount).Error

		if err != nil {
			return err
		}

		if newCommentCount == 0 {
			err := tx.Model(&post).Update("comment_status", "无评论").Error

			if err != nil {
				return err
			}
		}

		return nil
	})
}

/*
*
添加评论时更新文章评论数量
*/
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	fmt.Println("Comment AfterCreate >>>")

	return tx.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&Post{}).Where("id = ?", c.PostID).
			Update("comment_count", gorm.Expr("comment_count + ?", 1)).Error
		if err != nil {
			return err
		}
		return nil
	})

}

func RunGorm(db *gorm.DB) {

	//自动迁移模型(创建/更新)
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		return
	}

	fmt.Println("Table Created DONE")

	users, err := getUserPostsWithComments(1, db)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println(users.Post[0].Comment[0].Content)

	post, err := getMostCommentedPost(db)
	if err != nil {
		panic(err)
	}

	fmt.Println(post)

}

// 查询用户的所有文件及评论
func getUserPostsWithComments(userID uint, db *gorm.DB) (*User, error) {
	var user User
	err := db.Debug().
		Preload("Post.Comment").
		Where("id = ?", userID).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// 查询评论数量最多的文章
func getMostCommentedPost(db *gorm.DB) (*Post, error) {
	var post Post

	subQuery := db.Model(&Comment{}).
		Select("post_id, COUNT(*) as comment_count").
		Group("post_id")

	err := db.Debug().Table("posts").
		Joins(
			"JOIN (?) as comment_counts ON posts.id = comment_counts.post_id",
			subQuery).
		Select("posts.*, comment_counts.comment_count").
		Order("comment_counts.comment_count desc").
		Limit(1).Find(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}
