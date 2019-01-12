package entity

import (
	"github.com/galaco/Lambda-Core/core/model"
	"testing"
)

func TestPropBase_GetModel(t *testing.T) {
	sut := PropBase{}
	if sut.GetModel() != nil {
		t.Error("model was set, but should not be")
	}

	mod := &model.Model{}
	sut.SetModel(mod)

	if sut.GetModel() != mod {
		t.Errorf("set mode l does not match expected")
	}
}

func TestPropBase_SetModel(t *testing.T) {
	sut := PropBase{}
	if sut.GetModel() != nil {
		t.Error("model was set, but should not be")
	}

	mod := &model.Model{}
	sut.SetModel(mod)

	if sut.GetModel() != mod {
		t.Errorf("set mode l does not match expected")
	}
}
