package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Betriebsstelle struct {
	Bahnhof              bool     `json:"bahnhof"`
	Betriebsstellentypen []string `json:"betriebsstellentypen"`
	Bundesland           string   `json:"bundesland"`
	DS100                string   `json:"ds100"`
	GeoKoordinaten       struct {
		Breite float64 `json:"breite"`
		Laenge float64 `json:"laenge"`
	} `json:"geo_koordinaten"`
	ID                 int    `json:"id"`
	LangnameStammdaten string `json:"langname_stammdaten"`
	PrimaryCode        int    `json:"primary_code"`
}

type Infrastructure struct {
	Anzeigename    string `json:"anzeigename"`
	Fahrplanjahr   int    `json:"fahrplanjahr"`
	GueltigBis     string `json:"gueltig_bis"`
	GueltigVon     string `json:"gueltig_von"`
	ID             int    `json:"id"`
	Ordnungsrahmen struct {
		Betriebsstellen []Betriebsstelle `json:"betriebsstellen"`
	} `json:"ordnungsrahmen"`
}

func GetInfrastructure(infrastructure_id int) (*Infrastructure, error) {
	resp, err := http.Get(fmt.Sprintf("https://trassenfinder.de/api/web/infrastrukturen/%d", infrastructure_id))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	infrastructure := Infrastructure{}
	json.Unmarshal(body, &infrastructure)
	return &infrastructure, nil
}
