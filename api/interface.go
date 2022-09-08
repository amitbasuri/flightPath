package api

type FlightJourneyInterface interface {
	GetFlightStartingAndEndingAirportCode(data [][]string) (string, string, error)
}
