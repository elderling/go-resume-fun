package main

import (
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
	"gopkg.in/yaml.v3"
)

type Job struct {
	Title       string    `yaml:"title"`
	Company     string    `yaml:"company"`
	Location    string    `yaml:"location"`
	Description string    `yaml:"description"`
	StartDate   time.Time `yaml:"start_date"`
	EndDate     time.Time `yaml:"end_date"`
}

func main() {

	austin, err := time.LoadLocation("America/Chicago")

	if err != nil {
		panic(err)
	}

	job := Job{
		Title:       "Principal Software Architect",
		Company:     "Switch",
		Location:    "Austin, TX",
		Description: "Transition one billing and account management system into multiple systems.",
		StartDate:   time.Date(2022, time.January, 0, 0, 0, 0, 0, austin),
		EndDate:     time.Date(2023, time.March, 0, 0, 0, 0, 0, austin),
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, job.Title)
	err2 := pdf.OutputFileAndClose("hello.pdf")

	if err2 != nil {
		panic(err2.Error())
	}

	theYaml, err3 := yaml.Marshal(job)

	if err3 != nil {
		panic(err3.Error())
	}

	fmt.Println(string(theYaml))
}
