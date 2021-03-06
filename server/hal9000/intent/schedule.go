package intent

import (
	"fmt"
	"strings"
	"time"

	"github.com/johnjones4/hal-9000/server/hal9000/core"
	"github.com/johnjones4/hal-9000/server/hal9000/service"
	"github.com/johnjones4/hal-9000/server/hal9000/util"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
)

type Schedule struct {
	Service *service.Google
}

func (c *Schedule) SupportedComandsForState(s string) map[string]core.CommandInfo {
	if s != core.StateDefault {
		return map[string]core.CommandInfo{}
	}
	return map[string]core.CommandInfo{
		"schedule": {
			Description:  "Add a new event to the calendar.",
			RequiresBody: true,
		},
	}
}

func (c *Schedule) Execute(req core.Inbound) (core.Outbound, error) {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	dateInfo, err := w.Parse(req.Body, time.Now()) //TODO better parsing
	if err != nil {
		return core.Outbound{}, err
	}

	event := service.Event{
		Name:  strings.TrimSpace(strings.ReplaceAll(req.Body, dateInfo.Text, "")),
		Start: dateInfo.Time,
		End:   dateInfo.Time.Add(time.Hour),
	}

	createdEvent, err := c.Service.CreateNewEvent(event)
	if err != nil {
		return core.Outbound{}, err
	}

	return core.Outbound{
		OutboundBody: core.OutboundBody{
			Body: fmt.Sprintf("Scheduled \"%s\" for %s", createdEvent.Summary, util.FormatTime(dateInfo.Time)),
			URL:  createdEvent.HtmlLink,
		},
		State: req.State,
	}, nil
}
