package samples

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func main() {
	x, err := os.CreateTemp("", "/tmp/demo2")
	if err != nil {
		fmt.Println("Error while writing!")
	} else if err = x.Close(); err != nil {
		fmt.Println("Error while closing!")
	}

	err = ioutil.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}

	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}

	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
	err = os.WriteFile(path.Join("/var", x.Name()), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
	err = os.WriteFile(x.Name(), []byte("This is some data"), 0644)
	if err != nil {
		fmt.Println("Error while writing!")
	}
}
