package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestFluvius(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fluvius Suite")
}
