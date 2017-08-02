package parcel

import (
	l "getswift/location"
	"fmt"
	"io/ioutil"
	"net/http"	
	"github.com/gorilla/schema"
	"encoding/json"
)

type Parcel struct {
	PackageId 		int64
	Destination 	l.Location
	Deadline		int64
	Distance 		float64
	DeliveryTime 	float64
}

func GetParcels(r *http.Request) ([]byte) {
	resp, err := http.Get("https://codetest.kube.getswift.co/packages")
	if err != nil {
		// handle error
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	dat := make([]Parcel,0)	

    if err := json.Unmarshal([]byte(body), &dat); err != nil {
        panic(err)
    }
    // fmt.Println(dat)
    // fmt.Printf("%#v", dat)
    for i := 0; i < len(dat); i++ { //Calculate starting distances    	
    	dat[i].DistanceFromDelivery()
    	// dat[i].Destination.GetPlace()
    	dat[i].Destination.Place = "Loading..."
	}

    out, _ := json.Marshal(dat)
    return out 
}

func (parcel *Parcel) DistanceFromDelivery() {
	
	var depotLocation l.Location
	depotLocation.Latitude = -37.816664 //Taken from Google Maps
	depotLocation.Longitude = 144.963848

	
	parcel.Distance = parcel.Destination.DistanceBetween(depotLocation)
	// parcel.DeliveryTime = parcel.Distance / 20 * 3600;

}

func ParcelPlace(r *http.Request) ([]byte){ 
	err := r.ParseForm()
    if err != nil {
            // Handle error
    }

    var parcel Parcel

    decoder := schema.NewDecoder()
    decoder.Decode(&parcel, r.PostForm)
    fmt.Println(parcel)

    out, _ := json.Marshal(parcel)
    return out 
}