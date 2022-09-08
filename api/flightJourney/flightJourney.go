package flightJourney

import (
	"errors"
	"flightPath/api"
)

type flightJourneySvc struct{}

func NewFlightJourneyService() api.FlightJourneyInterface {
	return &flightJourneySvc{}
}

// GetFlightStartingAndEndingAirportCode searches for possible flight path and return starting and ending airport code.
func (svc *flightJourneySvc) GetFlightStartingAndEndingAirportCode(tickets [][]string) (string, string, error) {

	h, err := NewFlightPathSearchHelper(tickets)
	if err != nil {
		return "", "", err
	}
	for source, _ := range h.destinations {
		h.Search(source, []string{source})
		if h.found {
			break
		}
	}

	if !h.found {
		return "", "", errors.New("flight path not found")
	}

	return h.sol[0], h.sol[len(h.sol)-1], nil

}

type flightPathSearchHelper struct {
	n            int
	destinations map[string][]string // source to multiple destinations mapping
	found        bool
	sol          []string
}

func NewFlightPathSearchHelper(tickets [][]string) (*flightPathSearchHelper, error) {
	destinations := make(map[string][]string)
	for _, ticket := range tickets {
		if len(ticket) != 2 {
			return nil, errors.New("invalid input")
		}
		destinations[ticket[0]] = append(destinations[ticket[0]], ticket[1])
	}

	return &flightPathSearchHelper{
		destinations: destinations,
		n:            len(tickets),
		found:        false,
		sol:          make([]string, 0),
	}, nil
}

// Search searches for possible paths from a given source
func (h *flightPathSearchHelper) Search(source string, trip []string) {
	if len(trip) == h.n+1 {
		h.sol = append([]string{}, trip...)
		h.found = true
		return
	}

	// no destinations available from source
	if len(h.destinations[source]) == 0 {
		return
	}

	destinationsFromSource := len(h.destinations[source])

	for i := 0; i < destinationsFromSource; i++ {

		// visit h destination
		destination := h.destinations[source][0]

		// remove from destinations map
		h.destinations[source] = h.destinations[source][1:]

		// add destination to trip
		trip = append(trip, destination)

		// recurse
		h.Search(destination, trip)
		if h.found {
			return
		}

		// if not found remove destination from trip
		trip = trip[:len(trip)-1]

		// if not found add destination to destinations map
		h.destinations[source] = append(h.destinations[source], destination)
	}
}
