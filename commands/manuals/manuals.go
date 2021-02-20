package manuals

// Manual represents the coupling of the one-line summary and detailed
// explaination of a cli command or general Lawyer topic.
type Manual struct {
	Summary string
	Content string
}

// ManualMap represents a one-to-one mapping between a cli help keyword
// (ex. lawyer help "indict") and the corresponding Manual.
type ManualMap = map[string]Manual

// CommandManualMap is a list of all cli command Manuals mapped by their
// corresponding help keywords.
var CommandManualMap = ManualMap{
	"help":   Manual{"print this help page", "stop repeating yourself"},
	"indict": indictManual,
}

// TopicManualMap is a list of all help topics mapped by their corresponding
// help keywords.
var TopicManualMap = ManualMap{
	"law": lawManual,
}

// EntireManualMap is a list of all cli command Manuals AND help topic mapped
// by their corresponding help keywords.
var EntireManualMap = merge(CommandManualMap, TopicManualMap)

func merge(a, b ManualMap) ManualMap {
	for key, value := range a {
		b[key] = value
	}
	return b
}
