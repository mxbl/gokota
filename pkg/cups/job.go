package cups

import "fmt"

type Job struct {
	id      int
	user    string
	title   string
	copies  int
	options string
	file    string
}

func (j *Job) String() string {
	return "Job{" +
		"id: " + fmt.Sprint(j.id) +
		", user: " + j.user +
		", title: " + j.title +
		", copies: " + fmt.Sprint(j.copies) +
		", options: " + j.options +
		", file: " + j.file +
		"}"
}
