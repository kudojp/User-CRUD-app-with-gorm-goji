package models

import(
	"time"
	"regexp"
)

type User struct {
	Id int64
	Name string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// gormはDeletedAtをつけると自動でsoft deleteしてくれる
	DeletedAt time.Time
}

func UserValidate(user User) (error){
	Validator := valval.Object(valval.M{
		"Name": valval.String(
			valval.MaxLength(20),
			// MustCompileは正規表現の解析に失敗したらpanic(ランタイムエラー)
			// $は行の末尾
			valval.Regexp(regexp.MustCompile(`^[a-z ]+$`))
		),
	})
	return Validator.Validate(user)
}