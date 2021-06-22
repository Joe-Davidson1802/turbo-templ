package turbo

import "github.com/a-h/templ"

type LoadingType int

const (
	Eager LoadingType = iota
	Lazy
)

func (lt LoadingType) String() string {
	if lt == Lazy {
		return "lazy"
	}
	return "eager"
}

type TurboFrameOptions struct {
	Id         string
	Src        string
	Loading    LoadingType
	Disabled   bool
	Target     string
	Autoscroll bool
	Contents   *templ.Component
}

type ActionType int

const (
	Unset ActionType = iota
	After
	Append
	Before
	Prepend
	Remove
	Replace
	Update
)

func (at ActionType) String() string {
	switch at {
	case After:
		return "after"
	case Append:
		return "append"
	case Before:
		return "before"
	case Prepend:
		return "prepend"
	case Remove:
		return "remove"
	case Replace:
		return "replace"
	case Update:
		return "update"
	}
	return ""
}

type TurboStreamOptions struct {
	Action   ActionType
	Target   string
	Contents *templ.Component
}
