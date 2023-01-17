package main

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestReserve(t *testing.T) {
	assert.Equal(t, reverse("hello world"), "dlrow olleh")
}
