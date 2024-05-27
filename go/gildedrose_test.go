package main

import (
	"testing"
)

func TestNormalBeforeSellDate(t *testing.T) {
	rose := NewGildedRose("normal", 10, 5)
	rose.Tick()

	if rose.GetQuality() != 9 {
		t.Errorf("Expected quality 9 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestNormalOnSellDate(t *testing.T) {
	rose := NewGildedRose("normal", 10, 0)
	rose.Tick()

	if rose.GetQuality() != 8 {
		t.Errorf("Expected quality 8 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestNormalAfterSellDate(t *testing.T) {
	rose := NewGildedRose("normal", 10, -1)
	rose.Tick()

	if rose.GetQuality() != 8 {
		t.Errorf("Expected quality 8 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -2 {
		t.Errorf("Expected days remaining -2 but got %d", rose.GetDaysRemaining())
	}
}

func TestNormalOfZeroQuality(t *testing.T) {
	rose := NewGildedRose("normal", 0, 5)
	rose.Tick()

	if rose.GetQuality() != 0 {
		t.Errorf("Expected quality 0 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieBeforeSellDate(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 10, 5)
	rose.Tick()

	if rose.GetQuality() != 11 {
		t.Errorf("Expected quality 11 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieBeforeSellDateWithMaxQuality(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 50, 5)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieOnSellDate(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 10, 0)
	rose.Tick()

	if rose.GetQuality() != 12 {
		t.Errorf("Expected quality 12 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieOnSellDateNearMaxQuality(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 49, 0)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieOnSellDateWithMaxQuality(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 50, 0)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieAfterSellDate(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 40, -1)
	rose.Tick()

	if rose.GetQuality() != 42 {
		t.Errorf("Expected quality 42 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -2 {
		t.Errorf("Expected days remaining -2 but got %d", rose.GetDaysRemaining())
	}
}

func TestBrieAfterSellDateWithMaxQuality(t *testing.T) {
	rose := NewGildedRose("Aged Brie", 50, -1)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -2 {
		t.Errorf("Expected days remaining -2 but got %d", rose.GetDaysRemaining())
	}
}

func TestSulfurasBeforeSellDate(t *testing.T) {
	rose := NewGildedRose("Sulfuras, Hand of Ragnaros", 10, 5)
	rose.Tick()

	if rose.GetQuality() != 10 {
		t.Errorf("Expected quality 10 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 5 {
		t.Errorf("Expected days remaining 5 but got %d", rose.GetDaysRemaining())
	}
}

func TestSulfurasOnSellDate(t *testing.T) {
	rose := NewGildedRose("Sulfuras, Hand of Ragnaros", 10, 0)
	rose.Tick()

	if rose.GetQuality() != 10 {
		t.Errorf("Expected quality 10 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 0 {
		t.Errorf("Expected days remaining 0 but got %d", rose.GetDaysRemaining())
	}
}

func TestSulfurasAfterSellDate(t *testing.T) {
	rose := NewGildedRose("Sulfuras, Hand of Ragnaros", 10, -1)
	rose.Tick()

	if rose.GetQuality() != 10 {
		t.Errorf("Expected quality 10 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageLongBeforeSellDate(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 20)
	rose.Tick()

	if rose.GetQuality() != 11 {
		t.Errorf("Expected quality 11 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 19 {
		t.Errorf("Expected days remaining 19 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageMediumCloseToSellDateUpperBound(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 10)
	rose.Tick()

	if rose.GetQuality() != 12 {
		t.Errorf("Expected quality 12 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 9 {
		t.Errorf("Expected days remaining 9 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageMediumCloseToSellDateUpperBoundAtMaxQuality(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 10)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 9 {
		t.Errorf("Expected days remaining 9 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageMediumCloseToSellDateLowerBound(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 6)
	rose.Tick()

	if rose.GetQuality() != 12 {
		t.Errorf("Expected quality 12 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 5 {
		t.Errorf("Expected days remaining 5 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageMediumCloseToSellDateLowerBoundAtMaxQuality(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 6)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 5 {
		t.Errorf("Expected days remaining 5 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageVeryCloseToSellDateUpperBound(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 5)
	rose.Tick()

	if rose.GetQuality() != 13 {
		t.Errorf("Expected quality 13 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageVeryCloseToSellDateUpperBoundAtMaxQuality(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 5)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageVeryCloseToSellDateLowerBound(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 1)
	rose.Tick()

	if rose.GetQuality() != 13 {
		t.Errorf("Expected quality 13 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 0 {
		t.Errorf("Expected days remaining 0 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageVeryCloseToSellDateLowerBoundAtMaxQuality(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 1)
	rose.Tick()

	if rose.GetQuality() != 50 {
		t.Errorf("Expected quality 50 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 0 {
		t.Errorf("Expected days remaining 0 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageOnSellDate(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 0)
	rose.Tick()

	if rose.GetQuality() != 0 {
		t.Errorf("Expected quality 0 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestBackstageAfterSellDate(t *testing.T) {
	rose := NewGildedRose("Backstage passes to a TAFKAL80ETC concert", 10, -1)
	rose.Tick()

	if rose.GetQuality() != 0 {
		t.Errorf("Expected quality 0 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -2 {
		t.Errorf("Expected days remaining -2 but got %d", rose.GetDaysRemaining())
	}
}

func TestConjuredBeforeSellDate(t *testing.T) {
	t.Skip("conjured")
	rose := NewGildedRose("Conjured", 10, 10)
	rose.Tick()

	if rose.GetQuality() != 8 {
		t.Errorf("Expected quality 8 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 9 {
		t.Errorf("Expected days remaining 9 but got %d", rose.GetDaysRemaining())
	}
}

func TestConjuredOnSellDate(t *testing.T) {
	t.Skip("conjured")
	rose := NewGildedRose("Conjured", 10, 0)
	rose.Tick()

	if rose.GetQuality() != 6 {
		t.Errorf("Expected quality 6 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -1 {
		t.Errorf("Expected days remaining -1 but got %d", rose.GetDaysRemaining())
	}
}

func TestConjuredAfterSellDate(t *testing.T) {
	t.Skip("conjured")
	rose := NewGildedRose("Conjured", 10, -1)
	rose.Tick()

	if rose.GetQuality() != 6 {
		t.Errorf("Expected quality 6 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != -2 {
		t.Errorf("Expected days remaining -2 but got %d", rose.GetDaysRemaining())
	}
}

func TestConjuredOfZeroQuality(t *testing.T) {
	t.Skip("conjured")
	rose := NewGildedRose("Conjured", 0, 5)
	rose.Tick()

	if rose.GetQuality() != 0 {
		t.Errorf("Expected quality 0 but got %d", rose.GetQuality())
	}
	if rose.GetDaysRemaining() != 4 {
		t.Errorf("Expected days remaining 4 but got %d", rose.GetDaysRemaining())
	}
}
