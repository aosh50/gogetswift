package parcel

import (
	l "getswift/location"
	"fmt"
	"io/ioutil"
	"net/http"	
	// "github.com/gorilla/schema"
	"encoding/json"
)

type Parcel struct {
	PackageId 		int64
	Destination 	l.Location
	Deadline		int64
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
    fmt.Println(dat)
    fmt.Printf("%#v", dat)

    out, _ := json.Marshal(dat)
    return out 
}