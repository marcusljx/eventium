package client_test_test

import (
	"context"
	"testing"

	"encoding/json"

	"github.com/marcusljx/eventium/eventiumpb"
	"github.com/stretchr/testify/assert"
)

const (
	testHeadline          = "ERMAHGERD, TEST"
	testLocationX float32 = 1.2345
	testLocationY float32 = 5.4321
)

var (
	testingUUID string
)

func TestServerPost(t *testing.T) {
	e := &eventiumpb.Eventum{
		Headline: testHeadline,
		Location: &eventiumpb.Location{
			X: testLocationX,
			Y: testLocationY,
		},
	}
	response, err := et.PostEvent(context.Background(), e)
	assert.NoError(t, err)
	assert.True(t, response.Success)
	testingUUID = response.EventID

	r, _ := json.MarshalIndent(response, "", "\t")
	t.Logf("response:\n%s", r)
}

func TestServerGet(t *testing.T) {
	response, err := et.GetEventByID(context.Background(), &eventiumpb.GetEventByIDRequest{Id: testingUUID})
	assert.NoError(t, err)
	assert.Equal(t, testHeadline, response.Headline)
	assert.Equal(t, testLocationX, response.Location.X)
	assert.Equal(t, testLocationY, response.Location.Y)

	r, _ := json.MarshalIndent(response, "", "\t")
	t.Logf("response:\n%s", r)
}
