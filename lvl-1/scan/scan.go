package scan

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func Int64(w io.Writer, r io.Reader) int64 {
	var input string
	for {
		_, err := fmt.Fscanln(r, input)
		if err != nil {
			fmt.Fprintf(w, "Ошибка ввода %v\nпопробуйте ввести число еще раз", err)
			continue
		}
		if input == "q" {
			os.Exit(0)
		}
		number, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Fprintf(w, "Ошибка: необходимо ввести целое число")
			continue
		}
		return number
	}
}

func Float64(w io.Writer, r io.Reader) float64 {
	var input string
	for {
		_, err := fmt.Fscanln(r, input)
		if err != nil {
			fmt.Fprintf(w, "Ошибка ввода %v\nпопробуйте ввести число еще раз", err)
			continue
		}
		if input == "q" {
			os.Exit(0)
		}
		number, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintln(w, "Ошибка: необходимо ввести число")
			continue
		}
		return number
	}
}
