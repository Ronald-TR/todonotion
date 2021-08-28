package todonotion_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTodonotion(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Todonotion Suite")
}
