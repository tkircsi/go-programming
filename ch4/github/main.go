// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"tkircsi/github/github"
	"tkircsi/github/report"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}

	report.PrintTextReport(os.Stdout, result)

	f, err := os.Create("report.html")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	report.PrintHTMLReport(f, result)
}
