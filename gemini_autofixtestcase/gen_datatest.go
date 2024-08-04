package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/securego/gosec/v2/testutils"
)

// writeFile writes data to a file at the specified path.
func writeFile(filePath, data string) error {
	err := ioutil.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}

// createOrReplaceFolder creates a new folder or replaces it if it already exists.
func createOrReplaceFolder(folderPath string) error {
	if _, err := os.Stat(folderPath); !os.IsNotExist(err) {
		err := os.RemoveAll(folderPath)
		if err != nil {
			return fmt.Errorf("failed to remove existing folder: %w", err)
		}
	}
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}
	return nil
}

func main() {
	folders := []string{
		// "g101", "g102", "g103",
		// "g104",
		// "g106", "g107", "g108", "g109", "g110", "g111", "g112", "g113", "g114", "g115",
		"g201", "g202", "g203", "g204",
	}
	data := map[string][]testutils.CodeSample{
		// "g101": testutils.SampleCodeG101,
		// "g102": testutils.SampleCodeG102,
		// "g103": testutils.SampleCodeG103,
		// "g104": testutils.SampleCodeG104,
		// "g106": testutils.SampleCodeG106,
		// "g107": testutils.SampleCodeG107,
		// "g108": testutils.SampleCodeG108,
		// "g109": testutils.SampleCodeG109,
		// "g110": testutils.SampleCodeG110,
		// "g111": testutils.SampleCodeG111,
		// "g112": testutils.SampleCodeG112,
		// "g113": testutils.SampleCodeG113,
		// "g114": testutils.SampleCodeG114,
		// "g115": testutils.SampleCodeG115,
		"g201": testutils.SampleCodeG201,
		"g202": testutils.SampleCodeG202,
		"g203": testutils.SampleCodeG203,
		"g204": testutils.SampleCodeG204,
	}

	for _, folder := range folders {
		// Create a new folder. If the folder already exists, it will be deleted.
		err := createOrReplaceFolder(fmt.Sprintf("./gemini_autofixtestcase/%s", folder))
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Folder created successfully")
		}

		// Create a new file
		for idx, samples := range data[folder] {
			filePath := fmt.Sprintf("./gemini_autofixtestcase/%s/%d_samples.go", folder, idx)
			file, err := os.Create(filePath)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("File created successfully")
			}
			defer file.Close()

			// Write the code samples to the file
			err = writeFile(filePath, strings.Join(samples.Code, ""))
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Code samples written successfully")
			}
		}
	}
}
