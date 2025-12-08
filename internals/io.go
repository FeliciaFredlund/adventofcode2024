package internals

// func GetDataFromUrl(url, filename/path) make http request and save to file
func getDataFromUrl(url, filepath string) ([]byte, error) {
	// IMPLEMENT

	// call SaveToFile(data, filepath)
	return nil, nil
}

// func GetDataFromFile(filename/path) get data from cached file
func getDataFromFile(filepath string) ([]byte, error) {
	// IMPLEMENT

	return nil, nil
}

// func IsCached(filename/path) checks if input have already been cached
func isCached(filepath string) bool {
	// IMPLEMENT

	return false
}

func GetData(url, filepath string) ([]byte, error) {
	if isCached(filepath) {
		// get data from file
		data, err := getDataFromFile(filepath)
		if err != nil {
			//DO STUFF
		}
		return data, nil
	}

	// make http request
	data, err := getDataFromUrl(url, filepath)
	if err != nil {
		//DO STUFF
	}

	return data, nil
}

// func SaveToFile(data, filename/path) to save the data to a file
func SaveToFile(data []byte, filepath string) error {
	// IMPLEMENT

	return nil
}
