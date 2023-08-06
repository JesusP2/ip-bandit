package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ip2location "github.com/ip2location/ip2location-go/v9"
)

type Tab struct {
	TabContent string
	TabUrl     string
}
type Item struct {
	Title   string
	Content string
}

func main() {
	db, err := ip2location.OpenDB("./iplocation/ip2location.bin")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	r.Static("/static", "./static/css")
	r.LoadHTMLGlob("static/**/*.html")
	r.GET("/", func(c *gin.Context) {
		results, err := db.Get_all(c.RemoteIP())
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "MyIp", gin.H{
			"TabOne":      Tab{TabContent: "My IP", TabUrl: "/"},
			"TabTwo":      Tab{TabContent: "Search IP", TabUrl: "/search-ip"},
			"Title":       "Main website",
			"IP":          c.RemoteIP(),
			"Coordinates": fmt.Sprintf("%f,%f", results.Latitude, results.Longitude),
			"Data": [6]Item{
				{Title: "Country", Content: results.Country_long},
				{Title: "Region", Content: results.Region},
				{Title: "City", Content: results.City},
				{Title: "Zipcode", Content: results.Zipcode},
				{Title: "Latitude", Content: fmt.Sprint(results.Latitude)},
				{Title: "Longitude", Content: fmt.Sprint(results.Longitude)},
			},
		})
	})
	r.GET("/search-ip", func(c *gin.Context) {
		c.HTML(http.StatusOK, "SearchIP", gin.H{
			"TabOne": Tab{TabContent: "My IP", TabUrl: "/"},
			"TabTwo": Tab{TabContent: "Search IP", TabUrl: "/search-ip"},
			"Title":  "Main website",
		})
	})

	r.POST("/search-ip", func(c *gin.Context) {
		ip := c.PostForm("ip")
		results, err := db.Get_all(ip)
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "SearchIPResult", gin.H{
			"IP":          ip,
			"Coordinates": fmt.Sprintf("%f,%f", results.Latitude, results.Longitude),
			"Data": [6]Item{
				{Title: "Country", Content: results.Country_long},
				{Title: "Region", Content: results.Region},
				{Title: "City", Content: results.City},
				{Title: "Zipcode", Content: results.Zipcode},
				{Title: "Latitude", Content: fmt.Sprint(results.Latitude)},
				{Title: "Longitude", Content: fmt.Sprint(results.Longitude)},
			},
		})
	})
	r.POST("/get-map", func(c *gin.Context) {
		coordinates := c.PostForm("Coordinates")
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "Map", gin.H{
			"Coordinates": coordinates,
		})
	})
	r.Run(":8080")
}
