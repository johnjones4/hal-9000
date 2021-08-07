package hal9000

import "hal9000/util"

type MessageIntent struct {
	Caller  Person `json:"caller"`
	Person  Person `json:"person"`
	Message string `json:"message"`
}

func NewMessageIntent(c Person, m ParsedRequestMessage) (MessageIntent, error) {
	person, messageStart, err := GetPersonInParsedRequestMessage(m)
	if err != nil {
		return MessageIntent{}, err
	}

	sendMessage := util.ConcatTokensInRange(m.Tokens, messageStart, len(m.Tokens))

	return MessageIntent{c, person, sendMessage}, nil
}

func (i MessageIntent) Execute(lastState State) (State, ResponseMessage, error) {
	err := SendMessageToPerson(i.Caller, i.Person, i.Message)
	if err != nil {
		return nil, ResponseMessage{}, err
	}

	return lastState, MessageOk(), nil
}

func GetPersonInParsedRequestMessage(m ParsedRequestMessage) (Person, int, error) {
	for _, entity := range m.NamedEntities {
		person, err := GetPersonByName(entity.Name)
		if err != nil && err != ErrorPersonNotFound {
			return Person{}, 0, err
		} else if err != ErrorPersonNotFound {
			return person, entity.Range.End, nil
		}
	}

	nouns := util.GetContiguousUniformTokens(m.Tokens, []string{"NN", "NNP", "NNPS", "NNS"})
	for _, nounSet := range nouns {
		for i := 0; i < len(nounSet.Tokens); i++ {
			for j := len(nounSet.Tokens); j >= i+1; j-- {
				nounStr := util.ConcatTokensInRange(nounSet.Tokens, i, j)
				person, err := GetPersonByName(nounStr)
				if err != nil && err != ErrorPersonNotFound {
					return Person{}, 0, err
				} else if err != ErrorPersonNotFound {
					return person, nounSet.End, nil
				}
			}
		}
	}

	return Person{}, 0, ErrorPersonNotFound
}