package loader

import (
	"os"
	"strconv"
	"text/template"

	"github.com/echojc/aocutil"
)

func Generate(year, day int) {
	err := GenerateFolders(year, day)

	if err != nil {
		println(err)
		os.Exit(1)
	}

	err = LoadData(year, day)

	if err != nil {
		panic(err)
	}

	os.Chdir("../..")
}

func GenerateFolders(year, day int) error {
	yearString := strconv.Itoa(year)

	err := generateFolder(yearString)

	if err != nil {
		return err
	}

	os.Chdir(yearString)

	dayString := strconv.Itoa(day)

	err = generateFolder(dayString)

	if err != nil {
		return err
	}

	os.Chdir(dayString)

	loadTemplate(year, day)

	return nil
}

func loadTemplate(year, day int) error {
	template, err := template.ParseFiles("../../template.tmp")

	if err != nil {
		return err
	}

	type TemplateData struct {
		Year int
		Day int
	}

	f, err := os.Create("main.go")

	if err != nil {
		return err
	}

	template.Execute(f, TemplateData{Year: year, Day: day})

	f.Close()

	return nil
}

func generateFolder(name string) error  {
	// Check if Folder exists
	if _, err := os.Stat(name); os.IsNotExist(err) {
		// Create Folder
		if err := os.Mkdir(name, 0755); err != nil {
			return err
		}
	}

	return nil
}

func LoadData(year, day int) (error) {
	// check if file session_id exists
	if _, err := os.Stat("../../session_id"); os.IsNotExist(err) {
		//print current folder
		dir, _ := os.Getwd()
		println(dir)
		return err
	}

	input, err := aocutil.NewInputFromFile("../../session_id")

	if err != nil {
		return err
	}

	_, err = input.Strings(year, day)

	if err != nil {
		return err
	}

	return nil
}