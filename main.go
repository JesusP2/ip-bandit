package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ip2location "github.com/ip2location/ip2location-go/v9"
)

func main() {
	db, err := ip2location.OpenDB("./iplocation/ip2location.bin")
	if err != nil {
		panic(err)
	}
	results, err := db.Get_all("189.203.182.35")
	if err != nil {
		panic(err)
	}
	fmt.Printf("country_short: %s\n", results.Country_short)
	fmt.Printf("country_long: %s\n", results.Country_long)
	fmt.Printf("region: %s\n", results.Region)
	fmt.Printf("city: %s\n", results.City)
	fmt.Printf("latitude: %f\n", results.Latitude)
	fmt.Printf("longitude: %f\n", results.Longitude)
	fmt.Printf("domain: %s\n", results.Domain)
	fmt.Printf("zipcode: %s\n", results.Zipcode)
	fmt.Printf("timezone: %s\n", results.Timezone)
	fmt.Printf("elevation: %f\n", results.Elevation)
	r := gin.Default()
	r.Static("/static", "./static/css")
	r.LoadHTMLGlob("static/**/*")
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home", gin.H{
			"Title": "Main website",
			"IP":    c.RemoteIP(),
		})
	})
	r.Run(":8080")
}
