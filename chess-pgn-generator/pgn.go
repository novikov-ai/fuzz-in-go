package pgn

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const apiGameImport = "https://lichess.org/api/import"

func GetPictureURL(pgn string) (string, error) {
	respHTML, err := pgnImportRetrieveHTML(pgn)
	if err != nil {
		return "", err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(respHTML))
	pic := doc.Find(".text.position-gif")

	picURL, exists := pic.Attr("href")
	if !exists {
		return "", errors.New("picture not found")
	}

	flipBoard, err := blackMove(pgn)
	if err != nil {
		return "", err
	}

	if flipBoard {
		picURL += "&color=black"
	}

	return picURL, nil
}

func pgnImportRetrieveHTML(pgn string) ([]byte, error) {
	data := url.Values{
		"pgn": {pgn},
	}

	resp, err := http.PostForm(apiGameImport, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respHTML, nil
}

func blackMove(pgn string) (bool, error) {
	if pgn == "" {
		return false, errors.New("wrong pgn format")
	}

	split := strings.Split(pgn, " ")
	return len(split)%2 != 0, nil
}
