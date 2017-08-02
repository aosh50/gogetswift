package location

import (
	"math"	
	// "fmt"

	"net/http"	
	"strconv"
	"encoding/json"
)

type Location struct {
	Latitude		float64
	Longitude		float64
	Place 			string
}

type Place struct {
	ID 					int64
	City		        string
    Type                int
    Formatted_address   string
    Place_id            string
    Longitude           float64
    Latitude            float64   
    Creativeid  		int64
    Distance            float64
}

func (Location1 *Location) DistanceBetween(Location2 Location) (float64) {
    return Haversine(Location1.Longitude, Location1.Latitude, Location2.Longitude, Location2.Latitude)
}

func Haversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
    

    const earthRadius = float64(6371)
    
    var deltaLat = (latTo - latFrom) * (math.Pi / 180)
    var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)
    
    var a = math.Sin(deltaLat / 2) * math.Sin(deltaLat / 2) + 
        math.Cos(latFrom * (math.Pi / 180)) * math.Cos(latTo * (math.Pi / 180)) *
        math.Sin(deltaLon / 2) * math.Sin(deltaLon / 2)
    var c = 2 * math.Atan2(math.Sqrt(a),math.Sqrt(1-a))
    
    distance = earthRadius * c
    
    return
}


func (loc *Location) GetPlace() {
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng=" + strconv.FormatFloat(loc.Latitude, 'f', 6, 64) + "," + strconv.FormatFloat(loc.Longitude, 'f', 6, 64) + "&key=AIzaSyAOtA6rdfEm9KOSiQG_LHLbnF2C65fcR2s";
	// fmt.Println(url)
	
    var results Results

    resp, err := http.Get(url)
    if err != nil {
      // handle error
      panic(err.Error())
    }

    decoder := json.NewDecoder(resp.Body)
    decoder.Decode(&results)
    
    loc.Place = results.Results[0].FormattedAddress
    //out, err := json.Marshal(results)
    // fmt.Println(out)
}