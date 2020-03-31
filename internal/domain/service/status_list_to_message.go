package service

import (
	"fmt"

	"github.com/disiqueira/coronabot/internal/domain/model"
)

type (
	StatusListToMessage struct {
		limit int
	}
)

func NewStatusListToMessage(limit int) *StatusListToMessage {
	return &StatusListToMessage{
		limit: limit,
	}
}

func (s StatusListToMessage) Convert(statusList []model.Status) model.Message {
	result := ":biohazard_sign: *|COVID-19|*\n"
	result += fmt.Sprintf("_%d Most affected countries:_\n", s.limit)
	result += "```"
	result += fmt.Sprintf("|%20s|%12s|%10s|%13s|%11s|%10s|%13s|%9s|\n", "Country", "Total Cases", "New Cases", "Total Deaths", "New Deaths", "Recovered", "Active Cases", "Critical")
	result += fmt.Sprintf("|%20s|%12s|%10s|%13s|%11s|%10s|%13s|%9s|\n", "", "", "", "", "", "", "", "")
	for i, status := range statusList {
		if i < s.limit {
			result += fmt.Sprintf("|%20s|%12d|%10d|%13d|%11d|%10d|%13d|%9d|\n", status.Country(), status.Confirmed(), status.NewConfirmed(), status.Deaths(), status.NewDeaths(), status.Recovered(), status.ActiveCases(), status.Critical())
		}
	}
	result += "```"
	return model.NewMessage(result)
}
