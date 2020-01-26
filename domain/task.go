package domain

type Task interface {
	UpdateTitle(newTitle string)
	UpdateDescription(newDescription string)
	UpdateStatus()
	UpdateTags(tags ...string)

	GetID() string
	GetTitle() string
	GetDescription() string
	GetStatus() bool
	GetTags() []string
}
