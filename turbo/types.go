package turbo

import "github.com/a-h/templ"

type LoadingType int

const (
	EagerLoading LoadingType = iota
	LazyLoading
)

func (lt LoadingType) String() string {
	if lt == LazyLoading {
		return "lazy"
	}
	return "eager"
}

type TurboFrameOptions struct {
	Id       string
	Src      string
	Loading  LoadingType
	Target   string
	Contents *templ.Component
}

type ActionType int

const (
	UnsetAction ActionType = iota
	AfterAction
	AppendAction
	BeforeAction
	PrependAction
	RemoveAction
	ReplaceAction
	UpdateAction
)

func (at ActionType) String() string {
	switch at {
	case AfterAction:
		return "after"
	case AppendAction:
		return "append"
	case BeforeAction:
		return "before"
	case PrependAction:
		return "prepend"
	case RemoveAction:
		return "remove"
	case ReplaceAction:
		return "replace"
	case UpdateAction:
		return "update"
	}
	return ""
}

type TurboStreamOptions struct {
	Action   ActionType
	Target   string
	Contents *templ.Component
}
