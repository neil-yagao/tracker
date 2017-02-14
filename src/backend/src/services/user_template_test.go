package services

import "testing"

func TestAssignTemplateToUser(t *testing.T) {
	template := new(UserTemplate)
	template.UserIdentity = USER_INDENTITY
	template.TemplatesId = append(make([]int64, 0), 21)
	template.AssignTemplate()
	confirmTemplate := new(UserTemplate)
	confirmTemplate.UserIdentity = USER_INDENTITY
	confirmTemplate.GetUserTemplate()
	if len(confirmTemplate.TemplatesId) != 1 && len(confirmTemplate.Templates) != 1 {
		t.Error("unable to get correct number template")
	}
	if confirmTemplate.TemplatesId[0] == 1 && confirmTemplate.Templates[1] != "" {
		t.Error("unable to get correct template")
	}
}

func TestUnassignTemplateToUser(t *testing.T) {
	template := new(UserTemplate)
	template.UserIdentity = USER_INDENTITY
	template.TemplatesId = append(make([]int64, 0), 21)
	template.UnassignTemplate()
	confirmTemplate := new(UserTemplate)
	confirmTemplate.UserIdentity = USER_INDENTITY
	confirmTemplate.GetUserTemplate()
	var contains bool = false
	for _, template := range confirmTemplate.TemplatesId {
		if template == 21 {
			contains = true
		}
	}
	if contains {
		t.Error("unable to unassign template")
	}
}
