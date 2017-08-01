package location

import (
	"math"	
)

type Location struct {
	Latitude		float64
	Longitude		float64
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