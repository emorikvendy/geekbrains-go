package error_with_date

import (
	"fmt"
	"os"
	"time"
)

type ErrorWithDate struct {
	text string
	time time.Time
}

func New(text string) error {
	return &ErrorWithDate{
		text: text,
		time: time.Now(),
	}
}

func (e *ErrorWithDate) Error() string {
	return fmt.Sprintf("ErrorWithDate: %s\ntime:\n%s", e.text, e.time)
}

func CreateFile(filename string) (*os.File, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, New(err.Error())
	}
	return file, nil
}
