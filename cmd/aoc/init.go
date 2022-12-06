package main

import (
	"embed"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed tmpl
var templates embed.FS

var cmdInit = &cobra.Command{
	Use:       "init year",
	RunE:      runEInit,
	ValidArgs: []string{"year"},
	Args:      cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(cmdInit)
}

func runEInit(cmd *cobra.Command, args []string) error {
	year := args[0]

	if err := makeYearDir(year); err != nil {
		return err
	}

	return iterateDays(year)
}

func makeYearDir(year string) error {
	if _, err := os.Stat(year); os.IsNotExist(err) {
		mkerr := os.Mkdir(year, os.ModeDir)

		if mkerr != nil {
			return mkerr
		}
	} else if err != nil {
		return err
	} else {
		fmt.Println("Skipping year directory creation, already exists...")
	}

	return nil
}

func iterateDays(year string) error {
	failedDays := []int{}

	for i := 1; i <= 25; i++ {
		skip, err := makeDayDir(year, i)

		if err != nil {
			failedDays = append(failedDays, i)
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if !skip {
			err := populateDayDir(year, i)
			if err != nil {
				failedDays = append(failedDays, i)
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}

	if len(failedDays) != 0 {
		failedDaysStr := []string{}
		for _, failDay := range failedDays {
			failedDaysStr = append(failedDaysStr, strconv.Itoa(failDay))
		}

		return fmt.Errorf("failed to create/populate days: [%s]", strings.Join(failedDaysStr, ", "))
	}

	return nil
}

func makeDayDir(year string, day int) (skip bool, err error) {
	dayStr := fmt.Sprintf("%02d", day)
	path := fmt.Sprintf("%s/%s", year, dayStr)

	skip = true

	if _, err := os.Stat(path); os.IsNotExist(err) {
		mkerr := os.Mkdir(path, os.ModeDir)

		if mkerr != nil {
			return skip, mkerr
		} else {
			skip = false
			return skip, nil
		}
	} else if err != nil {
		return skip, err
	} else {
		fmt.Printf("Skipping day %02d directory creation, already exists...\r\n", day)
	}

	return
}

func populateDayDir(year string, day int) error {
	errorList := []error{}
	dayStr := fmt.Sprintf("%02d", day)
	dayPath, dayTestPath := fmt.Sprintf("%s/%s/%[2]s.go", year, dayStr), fmt.Sprintf("%s/%s/%[2]s_test.go", year, dayStr)
	testDataPath := fmt.Sprintf("%s/%s/testdata", year, dayStr)
	testDataSamplePath := fmt.Sprintf("%s/sample.txt", testDataPath)

	tmpl, err := template.ParseFS(templates, "tmpl/*.tmpl")
	if err != nil {
		return err
	}

	tmplData := map[string]string{
		"Year":         year,
		"Day":          dayStr,
		"DayTruncated": strconv.Itoa(day),
	}

	dayFile, err := os.Create(dayPath)
	if err != nil {
		errorList = append(errorList, err)
	} else {
		defer dayFile.Close()
	}

	err = tmpl.ExecuteTemplate(dayFile, "day.go.tmpl", tmplData)
	if err != nil {
		errorList = append(errorList, err)
	}

	dayTestFile, err := os.Create(dayTestPath)
	if err != nil {
		errorList = append(errorList, err)
	} else {
		defer dayTestFile.Close()
	}

	err = tmpl.ExecuteTemplate(dayTestFile, "day_test.go.tmpl", tmplData)
	if err != nil {
		errorList = append(errorList, err)
	}

	if err := os.Mkdir(testDataPath, os.ModeDir); err != nil {
		errorList = append(errorList, err)
	} else {
		err := os.WriteFile(testDataSamplePath, []byte(""), os.ModeAppend)
		if err != nil {
			errorList = append(errorList, err)
		}
	}

	if len(errorList) != 0 {
		errorMsg := fmt.Errorf("%d errors occurred whilst populating day %d directory: ", len(errorList), day)
		for _, err := range errorList {
			errorMsg = fmt.Errorf("%w; ", err)
		}

		return errorMsg
	}

	return nil
}
