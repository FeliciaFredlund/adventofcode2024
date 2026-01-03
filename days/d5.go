package days

import (
	"slices"
	"strconv"
	"strings"
)

func day5star1(data []byte) []byte {
	stringified := string(data)
	rules, reports, _ := strings.Cut(stringified, "\n\n")

	processedRules := processRules(rules)

	sumMiddlePages := 0
	goodReports, _ := sortReports(reports, processedRules)
	for _, report := range goodReports {
		pages := strings.Split(report, ",")
		middlePage, _ := strconv.Atoi(pages[len(pages)/2])
		sumMiddlePages += middlePage
	}

	return []byte(strconv.Itoa(sumMiddlePages))
}

func day5star2(data []byte) []byte {
	stringified := string(data)
	rules, reports, _ := strings.Cut(stringified, "\n\n")

	processedRules := processRules(rules)

	_, badReports := sortReports(reports, processedRules)

	sumMiddlePages := 0
	for _, report := range badReports {
		pages := strings.Split(report, ",")

		reducedRules := reduceRules(processedRules, pages)

		correctReport := make([]string, len(pages))
		for i := range correctReport {
			var page string
			for _, page = range pages {
				if _, exists := reducedRules[page]; !exists {
					break
				}
			}
			correctReport[i] = page

			// removing page from pages
			pageIndex := slices.Index(pages, page)
			if pageIndex != -1 {
				pages = slices.Delete(pages, pageIndex, pageIndex+1)
			}

			// removing page from all relevant rules in reduced rules
			for key, value := range reducedRules {
				rulesPageIndex := slices.Index(value, page)
				if rulesPageIndex == -1 {
					continue
				}
				if len(value) == 1 {
					delete(reducedRules, key)
					continue
				}
				reducedRules[key] = slices.Delete(reducedRules[key], rulesPageIndex, rulesPageIndex+1)
			}

		}
		slices.Reverse(correctReport)

		//calculating the sum
		middlePage, _ := strconv.Atoi(correctReport[len(correctReport)/2])
		sumMiddlePages += middlePage
	}

	return []byte(strconv.Itoa(sumMiddlePages))
}

func reduceRules(processedRules map[string][]string, pages []string) map[string][]string {
	reducedRules := map[string][]string{}
	for _, page := range pages {
		fullRule, exists := processedRules[page]
		if !exists {
			continue
		}
		rule := removedIrrelevantPages(slices.Clone(fullRule), pages)
		if len(rule) > 0 {
			reducedRules[page] = removedIrrelevantPages(slices.Clone(fullRule), pages)
		}
	}

	return reducedRules
}

func removedIrrelevantPages(rules []string, pages []string) []string {
	correctPages := []string{}

	for _, page := range pages {
		if slices.Contains(rules, page) {
			correctPages = append(correctPages, page)
		}
	}

	return correctPages
}

func processRules(rules string) map[string][]string {
	processedRules := map[string][]string{}
	for rule := range strings.FieldsSeq(rules) {
		left, right, found := strings.Cut(rule, "|")
		if !found {
			continue
		}
		processedRules[left] = append(processedRules[left], right)
	}
	return processedRules
}

func sortReports(reports string, processedRules map[string][]string) (goodReports []string, badReports []string) {
	for report := range strings.FieldsSeq(reports) {
		pages := strings.Split(report, ",")
		slices.Reverse(pages)
		good := true
	reportLoop:
		for i, page := range pages {
			for j := i + 1; j < len(pages); j++ {
				if slices.Contains(processedRules[page], pages[j]) {
					good = false
					break reportLoop
				}
			}

		}

		if good {
			goodReports = append(goodReports, report)
		} else {
			badReports = append(badReports, report)
		}
	}

	return goodReports, badReports
}

/*
TURNS OUT IN THE REAL DATA THERE IS NO TRUE ORDER TO THE PAGES. THEY EXIST IN A CIRCLE
I count use parts of this for star 2

func day5star1(data []byte) []byte {
	stringified := string(data)
	rules, reports, _ := strings.Cut(stringified, "\n\n")

	// processing rules
	countOfBeingToTheRight := map[string]int{}
	pagesToTheRightOfKey := map[string][]string{}
	for rule := range strings.FieldsSeq(rules) {
		left, right, found := strings.Cut(rule, "|")
		if !found {
			continue
		}
		if _, exists := countOfBeingToTheRight[left]; !exists {
			countOfBeingToTheRight[left] = 0
		}
		countOfBeingToTheRight[right]++
		pagesToTheRightOfKey[left] = append(pagesToTheRightOfKey[left], right)
	}

	// putting together the right order of all pages
	pageOrder := make([]string, len(countOfBeingToTheRight))

	for i := range pageOrder {
		var key string
		var value int
		for key, value = range countOfBeingToTheRight {
			if value <= 0 {
				break
			}
		}
		pageOrder[i] = key

		for _, page := range pagesToTheRightOfKey[key] {
			countOfBeingToTheRight[page]--
		}

		delete(countOfBeingToTheRight, key)
		delete(pagesToTheRightOfKey, key)
	}

	// checking reports
	sumMiddlePages := 0
	for report := range strings.FieldsSeq(reports) {
		reportPages := strings.Split(report, ",")
		rpIndex := 0
		for _, poPage := range pageOrder {
			if poPage != reportPages[rpIndex] {
				continue
			}
			rpIndex++
			if rpIndex == len(reportPages) {
				middlePage, _ := strconv.Atoi(reportPages[len(reportPages)/2])
				sumMiddlePages += middlePage
				break
			}
		}
	}

	return []byte(strconv.Itoa(sumMiddlePages))
}
*/
