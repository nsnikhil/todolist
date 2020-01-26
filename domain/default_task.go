package domain

import (
	"github.com/lib/pq"
	"todolist/util"
)

type DefaultTask struct {
	Id          string         `db:"id"`
	Title       string         `db:"title"`
	Description string         `db:"description"`
	Status      bool           `db:"status"`
	Tags        pq.StringArray `db:"tags"`
}

func (t *DefaultTask) UpdateTitle(newTitle string) {
	util.DebugLog("[DefaultTask] [UpdateTitle]")
	t.Title = newTitle
}

func (t *DefaultTask) UpdateDescription(newDescription string) {
	util.DebugLog("[DefaultTask] [UpdateDescription]")
	t.Description = newDescription
}

func (t *DefaultTask) UpdateStatus() {
	util.DebugLog("[DefaultTask] [UpdateStatus]")
	t.Status = !t.Status
}

func (t *DefaultTask) UpdateTags(tags ...string) {
	util.DebugLog("[DefaultTask] [UpdateTags]")
	t.Tags = tags
}

func (t *DefaultTask) GetTitle() string {
	util.DebugLog("[DefaultTask] [GetTitle]")
	return t.Title
}

func (t *DefaultTask) GetDescription() string {
	util.DebugLog("[DefaultTask] [GetDescription]")
	return t.Description
}

func (t *DefaultTask) GetStatus() bool {
	util.DebugLog("[DefaultTask] [GetStatus]")
	return t.Status
}

func (t *DefaultTask) GetID() string {
	util.DebugLog("[DefaultTask] [GetID]")
	return t.Id
}

func (t *DefaultTask) GetTags() []string {
	util.DebugLog("[DefaultTask] [GetTags]")
	return t.Tags
}
