package models

import (
    "database/sql/driver"
    "errors"
    "fmt"
    "strings"
	"time"

	"github.com/goravel/framework/database/orm"
)

type StringSlice []string

type BlogsPost struct {
    orm.Model
    ID          uint        `gorm:"primaryKey"`
    Title       string      `gorm:"not null"`
    Slug        string      `gorm:"uniqueIndex;not null"`
    Content     string      `gorm:"type:text;not null"`
    Date        time.Time   `gorm:"not null"`
    Description string
    Tags        []string    `gorm:"-"`
    Categories  []string    `gorm:"-"`
    Draft       bool        `gorm:"default:false"`
    CreatedAt   time.Time   `gorm:"not null"`
    UpdatedAt   time.Time   `gorm:"not null"`  
}

type BlogPost struct {
    orm.Model
    ID          uint         `json:"id"`
    Title       string       `json:"title"`
    Slug        string       `json:"slug"`
    Content     string       `json:"content"`
    Date        time.Time    `json:"date"`
    Description string       `json:"description"`
    Tags        StringSlice  `json:"tags"`
    Categories  StringSlice  `json:"categories"`
    Draft       bool         `json:"draft"`
    CreatedAt   time.Time    `json:"created_at"`
    UpdatedAt   time.Time    `json:"updated_at"`
}

func (p *BlogPost) TableName() string {
    return "blog_posts"
}

func (ss *StringSlice) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return errors.New(fmt.Sprint("Failed to unmarshal StringSlice value:", value))
    }

    *ss = strings.Split(string(bytes), ",")
    return nil
}

func (ss StringSlice) Value() (driver.Value, error) {
    if len(ss) == 0 {
        return nil, nil
    }
    return strings.Join(ss, ","), nil
}
