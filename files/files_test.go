package files

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Files", func() {
	var (
		filesDir string
		err      error
	)

	BeforeEach(func() {
		filesDir, err = os.MkdirTemp("", "files")
		Expect(err).Should(Succeed())
		DeferCleanup(func() { os.RemoveAll(filesDir) })
	})

	When("are present", func() {
		BeforeEach(func() {
			_, err = os.Create(filepath.Join(filesDir, IndexFileName))
			Expect(err).Should(Succeed())
		})

		It("has index.html", func() {
			_, hi := getDirContent(filesDir)
			Expect(hi).Should(Equal(true))
		})

		It("has files", func() {
			_, err = os.Create(filepath.Join(filesDir, "test1.txt"))
			Expect(err).Should(Succeed())

			_, err = os.Create(filepath.Join(filesDir, "test2.txt"))
			Expect(err).Should(Succeed())

			f, hi := getDirContent(filesDir)
			Expect(hi).Should(Equal(true))
			Expect(f).Should(ContainSubstring("test1.txt"))
			Expect(f).Should(ContainSubstring("test2.txt"))
		})
	})

	When("are not present", func() {
		It("has no index.html", func() {
			_, err = os.Create(filepath.Join(filesDir, "test1.txt"))
			Expect(err).Should(Succeed())

			_, hi := getDirContent(filesDir)
			Expect(hi).Should(Equal(false))
		})

		It("has no files", func() {
			f, hi := getDirContent(filesDir)
			Expect(hi).Should(Equal(false))
			Expect(f).Should(HaveLen(0))
		})
	})
})
