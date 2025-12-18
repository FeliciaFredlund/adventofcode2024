package internals

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// func getDataFromUrl(url, filename/path) make http request and save to file
func getDataFromUrl(url, filepath string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request from url (%s): %w", url, err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response from url (%s): %w", url, err)
	}

	SaveToFile(data, filepath)
	return data, nil
}

// func getDataFromFile(filename/path) get data from cached file
func getDataFromFile(filepath string) ([]byte, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading file (%s): %w", filepath, err)
	}
	return data, nil
}

// func isCached(filename/path) checks if input have already been cached
func isCached(filepath string) bool {
	_, err := os.Stat(filepath)
	if err != nil {
		// either the file doesn't exist so it isn't cached, or something went wrong
		// recaching is fine in the second case
		return false
	}
	return true
}

func GetData(url, filepath string) ([]byte, error) {
	if isCached(filepath) {
		data, err := getDataFromFile(filepath)
		if err != nil {
			return nil, fmt.Errorf("error getting data from cached file: %w", err)
		}
		return data, nil
	}

	// get data from url instead since not cached
	data, err := getDataFromUrl(url, filepath)
	if err != nil {
		return nil, fmt.Errorf("error getting data from url: %w", err)
	}
	return data, nil
}

// func SaveToFile(data, filename/path) to save the data to a file
func SaveToFile(data []byte, filepath string) error {
	err := os.WriteFile(filepath, data, 0664)
	if err != nil {
		return fmt.Errorf("error saving data to file: %w", err)
	}
	return nil
}
