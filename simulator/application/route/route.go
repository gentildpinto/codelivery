package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type (
	Route struct {
		ID        string     `json:"route_id"`
		ClientID  string     `json:"client_id"`
		Positions []Position `json:"positions"`
	}

	Position struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	PartialRoutePosition struct {
		ID       string    `json:"route_id"`
		ClientID string    `json:"client_id"`
		Position []float64 `json:"position"`
		Finished bool      `json:"finished"`
	}
)

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("Route ID is empty")
	}

	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")

		latitude, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}

		longitude, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}

		r.Positions = append(r.Positions, Position{
			Latitude:  latitude,
			Longitude: longitude,
		})
	}
	return nil
}

func (r *Route) ExportJsonPositions() (result []string, err error) {
	var route PartialRoutePosition
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Latitude, v.Longitude}
		route.Finished = k == total-1

		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return
}
