package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestServer(t *testing.T) {
    if testing.Short() {
      t.Skip("skipping test in short mode.")
      return
    }	
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}
