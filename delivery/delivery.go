package delivery

import (
	"net/http"		
	"fmt"
	d "getswift/drone"
	p "getswift/parcel"
	"github.com/gorilla/schema"
)

type Delivery struct {
	drones 			[]d.Drone
	packages 		[]p.Parcel
}


func UpdateDelivery(r *http.Request) {
	err := r.ParseForm()
    if err != nil {
            // Handle error
    }

    var delivery Delivery

    decoder := schema.NewDecoder()
    decoder.Decode(&delivery, r.PostForm)
    fmt.Println(delivery)
    
}