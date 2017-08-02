package location

import (
   "net/http"
   "encoding/json"   
)


type Results struct {
   Results []Result `json:"results"`
   Status string `json:"status"`
}
type Result struct {
   AddressComponents []Address `json:"address_components"`
   FormattedAddress string `json:"formatted_address"`
   Geometry Geometry `json:"geometry"`
   PlaceId string `json:"place_id"`
   Types []string `json:"types"`
}

type Address struct {
   LongName string `json:"long_name"`
   ShortName string `json:"short_name"`
   Types []string `json:"types"`
}

type Geometry struct {
   Bounds Bounds `json:"bounds"`
   Location LatLng `json:"location"`
   LocationType string `json:"location_type"`
   Viewport Bounds `json:"viewport"`
}

type Bounds struct {
   Northeast LatLng `json:"northeast"`
   Southwest LatLng `json:"southwest"`
}

type LatLng struct {
   Lat float64 `json:"lat"`
   Lng float64 `json:"lng"`
}



func GetAddressByCoordinates(request *http.Request) ([]byte) {
    request.ParseForm()
    long := request.FormValue("Longitude")
    lat := request.FormValue("Latitude")    

    url := "http://maps.googleapis.com/maps/api/geocode/json?latlng=" + lat + "," + long + "&sensor=true"

    var results Results

    resp, err := http.Get(url)
    if err != nil {
      // handle error
      panic(err.Error())
    }

    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&results)
    

    out, err := json.Marshal(results)
    return out  

}

