package routing

import (

	"net/http"
	d "getswift/drone"
	p "getswift/parcel"
	gs "getswift/delivery"	
	)
func Test(w http.ResponseWriter, r *http.Request) {
	// d.GetDrones()	
}

func Packages(w http.ResponseWriter, r *http.Request) {
	result := p.GetParcels(r)	
	WriteContent(w, result)
}
func Drones(w http.ResponseWriter, r *http.Request) {
	result := d.GetDrones(r)	
	WriteContent(w, result)	
}

func Update(w http.ResponseWriter, r *http.Request) {
	// result := d.UpdateDronePositions(r)	
	// WriteContent(w, result)	
	gs.UpdateDelivery(r)
}

func LoadLocation(w http.ResponseWriter, r *http.Request) {
	result := p.ParcelPlace(r)	
	WriteContent(w, result)	
}