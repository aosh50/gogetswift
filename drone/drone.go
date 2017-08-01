package drone

import (
	l "getswift/location"
	p "getswift/parcel"
	"fmt"
	"io/ioutil"
	"net/http"	
	// "github.com/gorilla/schema"
	"encoding/json"
)

type Drone struct {
	DroneId 	int `json:"droneId"`
	Location 	l.Location `json:"location"`
	Packages 	[]p.Parcel `json:"packages"`
	DistanceDepot float64
}
type Drones struct {
	drones 		[]Drone
}

func ListDrones() {
	resp, err := http.Get("https://codetest.kube.getswift.co/drones")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func GetDrones(r *http.Request) ([]byte){
	resp, err := http.Get("https://codetest.kube.getswift.co/drones")
	if err != nil {
		// handle error
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	dat := make([]Drone,0)	

    if err := json.Unmarshal([]byte(body), &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)
    fmt.Printf("%#v", dat)

    for i := 0; i < len(dat); i++ {
    	
    	dat[i].DistanceFromDepot
	}

    out, _ := json.Marshal(dat)
    return out    
}

func (drone *Drone) DistanceFromDepot()(float64) {

	var depotLocation l.Location
	depotLocation.Longitude = -37.816480 //Taken from Google Maps
	depotLocation.Latitude = 144.963844

	drone.DistanceDepot = depotLocation.DistanceBetween(drone.Location)

}
