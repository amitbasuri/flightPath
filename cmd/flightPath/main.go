package main

import (
	"flightPath/api"
	"flightPath/api/flightJourney"
	"github.com/gin-gonic/gin"
	"net/http"
)

type restApi struct {
	flightJourneySvc api.FlightJourneyInterface
}

func main() {
	r := restApi{
		flightJourneySvc: flightJourney.NewFlightJourneyService(),
	}

	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/airport", r.GetFlightStartingAndEndingAirportCodeHandler)

	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")

}

func (r *restApi) GetFlightStartingAndEndingAirportCodeHandler(c *gin.Context) {

	var data struct {
		Flights [][]string `json:"flights" binding:"required"`
	}

	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	startingAirport, endingAirport, err := r.flightJourneySvc.GetFlightStartingAndEndingAirportCode(data.Flights)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"starting_airport": startingAirport, "ending_airport": endingAirport})

}
