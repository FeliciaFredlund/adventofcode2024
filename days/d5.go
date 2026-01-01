package days

func day5star1(data []byte) []byte {
	return data

	//sumMiddlePages := 0

	//return []byte(strconv.Itoa(sumMiddlePages))
}

func day5star2(data []byte) []byte {
	return data
}

/*

TURNS OUT IN THE REAL DATA THERE IS NO TRUE ORDER TO THE PAGES. THEY EXIST IN A CIRCLE

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
