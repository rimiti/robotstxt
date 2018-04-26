package main

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "os"
    "testing"
    "github.com/onsi/ginkgo/reporters"
)

func TestParser(t *testing.T) {
    RegisterFailHandler(Fail)
    junitReporter := reporters.NewJUnitReporter(os.Getenv("CI_REPORT"))
    RunSpecsWithDefaultAndCustomReporters(t, "Suite", []Reporter{junitReporter})
}
