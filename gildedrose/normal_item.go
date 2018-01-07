package gildedrose

type NormalItem struct {
	sellIn  int
	quality int
}

func (i *NormalItem) GetSellIn() int {
	return i.sellIn
}

func (i *NormalItem) GetQuality() int {
	return i.quality
}

func (i *NormalItem) decreaseSellIn() {
	if i.sellIn > 0 {
		i.sellIn--
	}
}

func (i *NormalItem) decreaseQuality() {
	if i.quality > 0 {
		i.quality--
	}
}

func (i *NormalItem) Change() {
	i.decreaseQuality()
}

func init() {
	Register("default", func(sellIn, quality int) Item {
		return &NormalItem{sellIn, quality}
	})
}
