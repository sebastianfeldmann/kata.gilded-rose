package main

import (
	"fmt"
)

type GildedRose struct {
	name    string
	quality int
	sellIn  int
}

func NewGildedRose(name string, quality int, sellIn int) *GildedRose {
	return &GildedRose{name: name, quality: quality, sellIn: sellIn}
}

func (g *GildedRose) GetQuality() int {
	return g.quality
}

func (g *GildedRose) GetDaysRemaining() int {
	return g.sellIn
}

func (g *GildedRose) Tick() {
	if g.name != "Aged Brie" && g.name != "Backstage passes to a TAFKAL80ETC concert" {
		if g.quality > 0 {
			if g.name != "Sulfuras, Hand of Ragnaros" {
				g.quality -= 1
			}
		}
	} else {
		if g.quality < 50 {
			g.quality += 1
			if g.name == "Backstage passes to a TAFKAL80ETC concert" {
				if g.sellIn < 11 {
					if g.quality < 50 {
						g.quality += 1
					}
				}
				if g.sellIn < 6 {
					if g.quality < 50 {
						g.quality += 1
					}
				}
			}
		}
	}
	if g.name != "Sulfuras, Hand of Ragnaros" {
		g.sellIn -= 1
	}
	if g.sellIn < 0 {
		if g.name != "Aged Brie" {
			if g.name != "Backstage passes to a TAFKAL80ETC concert" {
				if g.quality > 0 {
					if g.name != "Sulfuras, Hand of Ragnaros" {
						g.quality -= 1
					}
				}
			} else {
				g.quality = 0
			}
		} else {
			if g.quality < 50 {
				g.quality += 1
			}
		}
	}
}

func main() {
	// Example usage
	item := NewGildedRose("Aged Brie", 10, 5)
	item.Tick()
	fmt.Println("please execute the unit tests with 'go test'")
}
