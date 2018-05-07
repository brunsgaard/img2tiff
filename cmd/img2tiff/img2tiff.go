package main

import (
	"bufio"
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"net/http"

	"github.com/otiai10/gosseract"
	"golang.org/x/image/tiff"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		img, _, err := image.Decode(r.Body)
		if err != nil {
			// replace this with real error handling
			panic(err)
		}

		var b bytes.Buffer
		writer := bufio.NewWriter(&b)
		tiff.Encode(writer, img, &tiff.Options{})
		writer.Flush()

		client := gosseract.NewClient()
		defer client.Close()
		client.SetImageFromBytes(b.Bytes())
		client.SetLanguage("dan")
		client.SetPageSegMode(4)
		text, _ := client.HOCRText()
		fmt.Fprintf(w, "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n")
		fmt.Fprintf(w, text)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}

}
