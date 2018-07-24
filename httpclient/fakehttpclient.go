package httpclient

import (
	"fmt"
	"io/ioutil"
	"time"
)

// LoadPage returns page from url
func LoadPage(url string) string {

	time.Sleep(50 * time.Millisecond)

	if url == "https://mycompany.com/" {
		return loadFile("./TestData/home.html")
	} else if url == "https://mycompany.com/about" {
		return loadFile("./TestData/about.html")
	}

	return "<html>Not found</html>"
}

func loadFile(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Could not load file. %s, %s", filename, err)
		return ""
	}
	return string(dat)

}
