package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/go-pdf/fpdf"
	"gopkg.in/yaml.v3"
)

type Job struct {
	Title        string    `yaml:"title"`
	Company      string    `yaml:"company"`
	Location     string    `yaml:"location"`
	Description  string    `yaml:"description"`
	StartDate    time.Time `yaml:"start_date"`
	EndDate      time.Time `yaml:"end_date"`
	Skills       []string  `yaml:"skills"`
	SkillDomains []string  `yaml:"skill_domains"`
}
type Experience struct {
	Jobs []Job `yaml:"experience"`
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
		Description: "Transition one billing and account management system into multiple systems. Maintain and extend custom, in-house billing and account management systems in HTML, JavaScript, Perl/Moose, and MySQL. Estimate development effort for feature designs. Implement new feature designs and assist with testing. Mentor junior developer through code review and direction.",
		StartDate:   time.Date(2022, time.January, 0, 0, 0, 0, 0, austin),
		EndDate:     time.Date(2023, time.March, 0, 0, 0, 0, 0, austin),
	}

	skills := []string{"Computer Programming", "Software Design", "Leadership", "Mentoring"}
	sort.Strings(skills)

	job.Skills = skills

	skillDomains := []string{
		"Applications",
		"Back-end Services",
		"Salesforce",
		"Jira Service Desk",
		"Billing",
		"Account Management",
		"Informatica",
		"Leadership",
		"Credit Card Processing",
		"Datacenter",
	}

	job.SkillDomains = skillDomains

	sort.Strings(job.SkillDomains)

	jobs := Experience{}

	jobs.Jobs = append(jobs.Jobs, job)

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, job.Title)
	err2 := pdf.OutputFileAndClose("hello.pdf")

	if err2 != nil {
		panic(err2.Error())
	}

	theYaml, err3 := yaml.Marshal(jobs)

	if err3 != nil {
		panic(err3.Error())
	}

	fmt.Println(string(theYaml))

	var theExperience Experience
	err4 := yaml.Unmarshal(theYaml, &theExperience)
	if err4 != nil {
		panic(err4.Error())
	}

}
