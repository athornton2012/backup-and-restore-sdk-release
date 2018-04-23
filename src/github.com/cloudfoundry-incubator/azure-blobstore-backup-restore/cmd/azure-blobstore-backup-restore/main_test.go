package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"os/exec"
)

var _ = Describe("azure-blobstore-backup-restore binary", func() {
	It("no flags", func() {
		cmd := &exec.Cmd{
			Path: binPath,
			Args: []string{},
		}

		// action
		session, err := gexec.Start(cmd, GinkgoWriter, GinkgoWriter)
		Expect(err).NotTo(HaveOccurred())

		Expect(session).To(gexec.Exit(1))
	})
})
