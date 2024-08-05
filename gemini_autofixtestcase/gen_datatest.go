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
		// "g106", "g107", "g108", "g109", "g110",
		"g111",
		//  "g112", "g113", "g114", "g115",
		// "g201", "g202", "g203", "g204",
		// "g301", "g302", "g303", "g304", "g305",
		"g306",
		//   "g307",
		// "g401", "g402", "g403", "g404", "g405", "g406",
		// "g501", "g502", "g503", "g504", "g505", "g506", "g507",
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
		"g111": testutils.SampleCodeG111,
		// "g112": testutils.SampleCodeG112,
		// "g113": testutils.SampleCodeG113,
		// "g114": testutils.SampleCodeG114,
		// "g115": testutils.SampleCodeG115,
		// "g201": testutils.SampleCodeG201,
		// "g202": testutils.SampleCodeG202,
		// "g203": testutils.SampleCodeG203,
		// "g204": testutils.SampleCodeG204,
		// "g301": testutils.SampleCodeG301,
		// "g302": testutils.SampleCodeG302,
		// "g303": testutils.SampleCodeG303,
		// "g304": testutils.SampleCodeG304,
		// "g305": testutils.SampleCodeG305,
		"g306": testutils.SampleCodeG306,
		// "g307": testutils.SampleCodeG307,
		// "g401": testutils.SampleCodeG401,
		// "g402": testutils.SampleCodeG402,
		// "g403": testutils.SampleCodeG403,
		// "g404": testutils.SampleCodeG404,
		// "g405": testutils.SampleCodeG405,
		// "g406": testutils.SampleCodeG406,
		// "g501": testutils.SampleCodeG501,
		// "g502": testutils.SampleCodeG502,
		// "g503": testutils.SampleCodeG503,
		// "g504": testutils.SampleCodeG504,
		// "g505": testutils.SampleCodeG505,
		// "g506": testutils.SampleCodeG506,
		// "g507": testutils.SampleCodeG507,
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
