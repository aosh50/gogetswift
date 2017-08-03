package drone

import (
	l "getswift/location"
	p "getswift/parcel"
	"fmt"
	"io/ioutil"
	"net/http"	
	"github.com/gorilla/schema"
	"encoding/json"
)

type Drone struct {
	DroneId 			int `json:"droneId"`
	Location 			l.Location `json:"location"`
	Packages 			[]p.Parcel `json:"packages"`
	DistanceDepot 		float64
	DepotTime 			float64
	DeliveryDistance	float64
	DeliveryTime 		float64
}

func GetDrones(r *http.Request) ([]byte) {
	resp, err := http.Get("https://codetest.kube.getswift.co/drones") //Retrieve initial list of Drones
	if err != nil {
		// handle error
		panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	dat := make([]Drone,0)	

    if err := json.Unmarshal([]byte(body), &dat); err != nil { //Load HTTP Response into JSON
        panic(err)
    }

    for i := 0; i < len(dat); i++ { //Calculate starting distances
    	
    	dat[i].DistanceFromDepot() 
    	dat[i].DistanceFromDelivery()
    	// dat[i].Location.GetPlace()
    	dat[i].Location.Place = "Loading..."
    	if len(dat[i].Packages) > 0 {
    		dat[i].Packages[0].DroneDelivered = dat[i].DroneId
    	}
	}

    out, _ := json.Marshal(dat)
    return out    
}

func UpdateDronePositions(r *http.Request) ([]byte) {
	err := r.ParseForm()
    if err != nil {
            // Handle error
    }

    var drones []Drone

    decoder := schema.NewDecoder()
    decoder.Decode(&drones, r.PostForm)
    fmt.Println(drones)

	out, _ := json.Marshal(drones)
    return out  
}

func (drone *Drone) DistanceFromDepot() {

	var depotLocation l.Location
	depotLocation.Latitude = -37.816664 //Taken from Google Maps
	depotLocation.Longitude = 144.963848

	drone.DistanceDepot = depotLocation.DistanceBetween(drone.Location)
	drone.DepotTime = drone.DistanceDepot / 20 * 3600;

}

func (drone *Drone) DistanceFromDelivery() {

	if len(drone.Packages) > 0 {
		drone.DeliveryDistance = drone.Packages[0].Destination.DistanceBetween(drone.Location)
		drone.DeliveryTime = drone.DeliveryDistance / 20 * 3600;
	}

}
