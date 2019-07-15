package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Catalog struct {
	CatalogUUID          string `json:"catalog_uuid"`
	CatalogType          string `json:"catalog_type"`
	TotalPrice           int64  `json:"total_price"`
	MortgageAmount       int64  `json:"mortgage_amount"`
	MonthlyRent          int64  `json:"monthly_rent"`
	FloorQty             int64  `json:"floor_qty"`
	UnitAreaExtent       int64  `json:"unit_area_extent"`
	RoomQty              int64  `json:"room_qty"`
	ParkingQty           int64  `json:"parking_qty"`
	HasParking           bool   `json:"has_parking"`
	HasWarehouse         bool   `json:"has_warehouse"`
	PropertyAge          int64  `json:"property_age"`
	HasInternalWarehouse bool   `json:"has_internal_warehouse"`
	SaloonsImg           string `json:"saloons_img"`
	KitchensImg          string `json:"kitchens_img"`
	RoomsImg             string `json:"rooms_img"`
	DistrictName         string `json:"district_name"`
	UnitVariety          string `json:"unit_variety"`
	CountVariety         int64  `json:"count_variety"`
	CatalogRequestUUID   string `json:"catalog_request_uuid"`
	ImageCount           int64  `json:"image_count"`
	CreateTimestamp      string `json:"create_timestamp"`
	Score                int64  `json:"score"`
}

type Marketing struct {
	Title           string      `json:"title"`
	Description     string      `json:"description"`
	MetaTitle       interface{} `json:"meta_title"`
	MetaDescription interface{} `json:"meta_description"`
}

type Meta struct {
	TotalCount int64     `json:"total_count"`
	Marketing  Marketing `json:"marketing"`
}

type CatalogResponse struct {
	Data []Catalog `json:"data"`
	Meta Meta      `json:"meta"`
}

type CatalogDataPage struct {
	CatalogList []Catalog
}

func main() {
	resp, err := http.Get("https://address.ir/rest/tgr/catalogs/search/0/12")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var response CatalogResponse

	_ = json.Unmarshal([]byte(string(body)), &response)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := CatalogDataPage{CatalogList: response.Data}
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, p)
	})

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	port := ":7426"
	log.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}
