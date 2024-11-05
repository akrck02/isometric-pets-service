package main

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	http.HandleFunc("GET /pets/{uuid}", getPet)
	http.HandleFunc("GET /pets/{uuid}/image", getPetImage)
	http.ListenAndServe(":8081", nil)
}

/**
* Get pet
 */
func getPet(w http.ResponseWriter, r *http.Request) {

	parserError := r.ParseForm()

	if parserError != nil {
		sendResponseText(w, "Error parsing form")
		return
	}

	var uuid string = r.PathValue("uuid")
	w.Header().Set("Content-Type", "application/json")
	sendResponseText(w, "{ \"uuid\": \""+uuid+"\", \"name\": \"Teko\", \"age\": 5 }")
}

/**
* Get pet image
 */
func getPetImage(w http.ResponseWriter, r *http.Request) {

	parserError := r.ParseForm()

	if parserError != nil {
		sendResponseText(w, "Error parsing form")
		return
	}

	var uuid string = r.PathValue("uuid")

	log.Println("uuid: " + uuid)
	var image []byte = readImage(uuid)

	if image == nil {
		sendResponseText(w, "Image not found")
		return
	}

	sendResponsePng(w, image)
}

/**
* Send response text
 */
func sendResponseText(w http.ResponseWriter, text string) {
	w.Write([]byte(text))
}

/**
* Send response PNG
 */
func sendResponsePng(w http.ResponseWriter, image []byte) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(image)))
	if _, err := w.Write(image); err != nil {
		log.Println("unable to write image.")
	}
}

func readImage(uuid string) []byte {

	f, err := os.Open("resources/" + uuid + ".png")
	if err != nil {
		return nil
	}

	defer f.Close()
	image, _, err := image.Decode(f)

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, image); err != nil {
		log.Println("unable to encode image.")
	}

	if err != nil {
		return nil
	}

	return buffer.Bytes()
}
