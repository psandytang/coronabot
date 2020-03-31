package model

import (
	"strings"
)

type (
	Status struct {
		country    string
		confirmed      int
		newConfirmed   int
		deaths         int
		newDeaths      int
		recovered      int
		activeCases    int
		critical       int
	}

	StatusReporter interface {
		StatusPerCountry() ([]Status, error)
	}
)

func NewStatus(country string, confirmed int, newConfirmed int, deaths int, newDeaths int, recovered int, activeCases int, critical int) Status {
	return Status{
		country:       strings.TrimSpace(country),
		confirmed:     confirmed,
		newConfirmed:  newConfirmed,
		deaths:        deaths,
		newDeaths:     newDeaths,
		recovered:     recovered,
		activeCases:   activeCases,
		critical:      critical,
	}
}

func (s Status) Country() string {
	return s.country
}

func (s Status) Confirmed() int {
	return s.confirmed
}

func (s Status) NewConfirmed() int {
	return s.newConfirmed
}

func (s Status) Deaths() int {
	return s.deaths
}

func (s Status) NewDeaths() int {
	return s.newDeaths
}

func (s Status) Recovered() int {
	return s.recovered
}

func (s Status) ActiveCases() int {
	return s.activeCases
}

func (s Status) Critical() int {
	return s.critical
}
