package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
)

type Document struct {
	InfrastructureId int
	Filename         string
	DS100            string
}

type DocumentResponse struct {
	Dokumente []struct {
		Anzeigename string `json:"anzeigename"`
		Dateiname   string `json:"dateiname"`
		DS100       string `json:"ds100"`
		Typ         string `json:"typ"`
	} `json:"dokumente"`
}

func GetDocument(infrastructure_id int, betriebsstelle *Betriebsstelle) (*Document, error) {
	if betriebsstelle.DS100 == "" {
		return nil, nil
	}
	resp, err := http.Get(fmt.Sprintf("https://trassenfinder.de/api/web/infrastrukturen/%d/dokumente?ds100=%s", infrastructure_id, url.QueryEscape(betriebsstelle.DS100)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, nil
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var documentResponse DocumentResponse
	json.Unmarshal(body, &documentResponse)
	for _, document := range documentResponse.Dokumente {
		if document.Typ == "apn_skizze" {
			return &Document{
				InfrastructureId: infrastructure_id,
				Filename:         document.Dateiname,
				DS100:            betriebsstelle.DS100,
			}, nil
		}
	}

	return nil, nil
}

func (document *Document) GetFilepath(directory string) string {
	return path.Join(directory, document.DS100+"_"+document.Filename)
}

func (document *Document) Download(directory string) error {
	resp, err := http.Get(fmt.Sprintf("https://trassenfinder.de/api/web/infrastrukturen/%d/dokumente/%s", document.InfrastructureId, document.Filename))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	filepath := document.GetFilepath(directory)
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	file.Close()
	resp.Body.Close()
	return nil
}
