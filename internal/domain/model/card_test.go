package model_test

import (
	"testing"

	"github.com/omegaspn/mini-blog/internal/domain/model"
)

func TestIsValidCardStatus(t *testing.T) {
	testCases := []struct {
		status      model.CardStatus
		expected    bool
		description string
	}{
		{model.CardStatusGreen, true, "Valid status GREEN"},
		{model.CardStatusViolet, true, "Valid status VIOLET"},
		{model.CardStatusBlue, true, "Valid status BLUE"},
		{model.CardStatusOrange, true, "Valid status ORANGE"},
		{"INVALID_STATUS", false, "Invalid status"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := model.ValidCardStatuses[tc.status]
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v for status %s", tc.expected, actual, tc.status)
			}
		})
	}
}

func TestIsValidCardCatetory(t *testing.T) {
	testCases := []struct {
		category    model.CardCategory
		expected    bool
		description string
	}{
		{model.CardCategoryPhy, true, "Valid status PHY"},
		{model.CardCategoryTech, true, "Valid status TECH"},
		{model.CardCategoryChem, true, "Valid status CHEM"},
		{model.CardCategorySoc, true, "Valid status SOC"},
		{"INVALID_STATUS", false, "Invalid status"},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := model.ValidCardCategories[tc.category]
			if actual != tc.expected {
				t.Errorf("Expected %v, but got %v for status %s", tc.expected, actual, tc.category)
			}
		})
	}
}
