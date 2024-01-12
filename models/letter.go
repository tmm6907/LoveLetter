package models

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

type Letter struct {
	gorm.Model
	Title string
	Body  string
}

func validate(title string, body string) error {

	fmt.Println("Length:", len(body))
	if len(title) == 0 {
		return fmt.Errorf("error: no title provided")
	}

	if len(title) > 48 {
		return fmt.Errorf("error: title too long")
	}

	if len(body) == 0 {
		fmt.Println("body too short!!!")
		return fmt.Errorf("error: no body provided")
	}

	if len(body) > 240 {
		fmt.Println("Failed")
		return fmt.Errorf("error: no body provided")
	}
	return nil
}
func clean(title string, body string) (string, string) {
	titler := cases.Title(language.English)
	cleanTitle := titler.String(title)
	return cleanTitle, body
}

func NewLetter(subject string, body string) (*Letter, error) {
	err := validate(subject, body)
	// fmt.Println("validation error: ", err)
	if err != nil {
		return nil, err
	}
	subject, body = clean(subject, body)
	return &Letter{
		Title: subject,
		Body:  body,
	}, nil

}
