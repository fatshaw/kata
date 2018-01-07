package gildedrose_test

import (
	. "kata/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gilded Rose", func() {
	Context("when a normal item's sell in -1", func() {
		It("quality should -1", func() {
			item := Create("normal", 2, 5)
			item.Change()
			item.GetQuality()
			Expect(item.GetQuality()).To(Equal(4))
		})
	})
})
