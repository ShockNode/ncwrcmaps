package fishingareas

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	/*
		"github.com/stretchr/testify/assert"
		"github.com/stretchr/testify/mock"
	*/)

func TestRequestorGetFishingAreaInfo(t *testing.T) {

	requestor := Requestor{
		Client: *http.DefaultClient,
	}

	info, err := requestor.GetFishingAreaInfo(context.TODO(), 2389)
	assert.Nil(t, err)
	assert.NotNil(t, info)
	fmt.Printf("%+v\n", info)
}
