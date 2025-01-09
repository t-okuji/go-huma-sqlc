package main

import (
	"encoding/json"
	"testing"

	"github.com/danielgtaylor/huma/v2/humatest"
	"github.com/stretchr/testify/assert"
)

func TestGetGreeting(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Get("/greeting/world")

	assert.Equal(t, 200, resp.Code)

	var expectedGreeting, actualGreeting GreetingOutput
	if err := json.NewDecoder(resp.Body).Decode(&actualGreeting.Body); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	expectedGreeting.Body.Message = "Hello, world!"

	assert.Equal(t, expectedGreeting, actualGreeting, "Unexpected response: %s", resp.Body)
}

func TestPutReview(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"author": "daniel",
		"rating": 5,
	})

	assert.Equal(t, 201, resp.Code, "Unexpected status code: %d", resp.Code)
}

func TestPutReviewError(t *testing.T) {
	_, api := humatest.New(t)

	addRoutes(api)

	resp := api.Post("/reviews", map[string]any{
		"rating": 10,
	})

	assert.Equal(t, 422, resp.Code, "Unexpected status code: %d", resp.Code)
}
