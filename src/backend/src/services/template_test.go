package services

import (
	"models"
	"testing"
)

func TestSaveTemplate(t *testing.T) {
	template := new(Template)
	template.Name = "testingTemplateName"
	template.Content = "this is description for testing template."
	id := saveTemplate(*template)
	if id != -1 {
		defer models.BasicCRUD.Delete("template", id)
	}
	if id != findTemplate(template.Name).Id {
		t.Error("unable to handle save and read template!")
	}
}
