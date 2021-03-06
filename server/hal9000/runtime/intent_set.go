package runtime

import (
	"github.com/johnjones4/hal-9000/server/hal9000/core"
)

type IntentSet struct {
	Intents []core.Intent
}

func (h *IntentSet) ProcessRequest(req core.Inbound) (core.Outbound, error) {
	for _, commandHandler := range h.Intents {
		for command := range commandHandler.SupportedComandsForState(req.State) {
			if req.Command == command {
				res, err := commandHandler.Execute(req)
				if err != nil {
					if ferr, ok := err.(core.FeedbackError); ok {
						return core.Outbound{
							OutboundBody: core.OutboundBody{
								Body: ferr.Error(),
							},
							State: req.State,
						}, nil
					}
					return core.Outbound{}, err
				}
				return res, nil
			}
		}
	}
	return core.Outbound{
		OutboundBody: core.OutboundBody{
			Body: "I do not understand",
		},
		State: req.State,
	}, nil
}
