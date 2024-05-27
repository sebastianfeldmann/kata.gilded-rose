# -*- coding: utf-8 -*-
import unittest

from gilded_rose import GildedRose

class GildedRoseTest(unittest.TestCase):

    def testNormalBeforeSellDate(self):
        rose = GildedRose("normal", 10, 5)
        rose.tick()
        self.assertEqual(9, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())

    def testNormalOnSellDate(self):
        rose = GildedRose("normal", 10, 0)
        rose.tick()
        self.assertEqual(8, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testNormalAfterSellDate(self):
        rose = GildedRose("normal", 10, -1)
        rose.tick()
        self.assertEqual(8, rose.getQuality())
        self.assertEqual(-2, rose.getDaysRemaining())

    def testNormalOfZeroQuality(self):
        rose = GildedRose("normal", 0, 5)
        rose.tick()
        self.assertEqual(0, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())

    def testBrieBeforeSellDate(self):
        rose = GildedRose("Aged Brie", 10, 5)
        rose.tick()
        self.assertEqual(11, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())

    def testBrieBeforeSellDateWithMaxQuality(self):
        rose = GildedRose("Aged Brie", 50, 5)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())

    def testBrieOnSellDate(self):
        rose = GildedRose("Aged Brie", 10, 0)
        rose.tick()
        self.assertEqual(12, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testBrieOnSellDateNearMaxQuality(self):
        rose = GildedRose("Aged Brie", 49, 0)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testBrieOnSellDateWithMaxQuality(self):
        rose = GildedRose("Aged Brie", 50, 0)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testBrieAfterSellDate(self):
        rose = GildedRose("Aged Brie", 40, -1)
        rose.tick()
        self.assertEqual(42, rose.getQuality())
        self.assertEqual(-2, rose.getDaysRemaining())

    def testBrieAfterSellDateWithMaxQuality(self):
        rose = GildedRose("Aged Brie", 50, -1)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(-2, rose.getDaysRemaining())

    def testSulfurasBeforeSellDate(self):
        rose = GildedRose("Sulfuras, Hand of Ragnaros", 10, 5)
        rose.tick()
        self.assertEqual(10, rose.getQuality())
        self.assertEqual(5, rose.getDaysRemaining())

    def testSulfurasOnSellDate(self):
        rose = GildedRose("Sulfuras, Hand of Ragnaros", 10, 0)
        rose.tick()
        self.assertEqual(10, rose.getQuality())
        self.assertEqual(0, rose.getDaysRemaining())

    def testSulfurasAfterSellDate(self):
        rose = GildedRose("Sulfuras, Hand of Ragnaros", 10, -1)
        rose.tick()
        self.assertEqual(10, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testBackstageLongBeforeSellDate(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 20)
        rose.tick()
        self.assertEqual(11, rose.getQuality())
        self.assertEqual(19, rose.getDaysRemaining())

    def testBackstageMediumCloseToSellDateUpperBound(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 10)
        rose.tick()
        self.assertEqual(12, rose.getQuality())
        self.assertEqual(9, rose.getDaysRemaining())

    def testBackstageMediumCloseToSellDateUpperBoundAtMaxQuality(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 10)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(9, rose.getDaysRemaining())

    def testBackstageMediumCloseToSellDateLowerBound(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 6)
        rose.tick()
        self.assertEqual(12, rose.getQuality())
        self.assertEqual(5, rose.getDaysRemaining())

    def testBackstageMediumCloseToSellDateLowerBoundAtMaxQuality(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 6)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(5, rose.getDaysRemaining())

    def testBackstageVeryCloseToSellDateUpperBound(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 5)
        rose.tick()
        self.assertEqual(13, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())


    def testBackstageVeryCloseToSellDateUpperBoundAtMaxQuality(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 5)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())


    def testBackstageVeryCloseToSellDateLowerBound(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 1)
        rose.tick()
        self.assertEqual(13, rose.getQuality())
        self.assertEqual(0, rose.getDaysRemaining())


    def testBackstageVeryCloseToSellDateLowerBoundAtMaxQuality(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 1)
        rose.tick()
        self.assertEqual(50, rose.getQuality())
        self.assertEqual(0, rose.getDaysRemaining())

    def testBackstageOnSellDate(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 0)
        rose.tick()
        self.assertEqual(0, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())

    def testBackstageAfterSellDate(self):
        rose = GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, -1)
        rose.tick()
        self.assertEqual(0, rose.getQuality())
        self.assertEqual(-2, rose.getDaysRemaining())


    @unittest.skip("conjured")
    def testConjuredBeforeSellDate(self):
        rose = GildedRose("Conjured", 10, 10)
        rose.tick()
        self.assertEqual(8, rose.getQuality())
        self.assertEqual(9, rose.getDaysRemaining())

    @unittest.skip("conjured")
    def testConjuredOnSellDate(self):
        rose = GildedRose("Conjured", 10, 0)
        rose.tick()
        self.assertEqual(6, rose.getQuality())
        self.assertEqual(-1, rose.getDaysRemaining())


    @unittest.skip("conjured")
    def testConjuredAfterSellDate(self):
        rose = GildedRose("Conjured", 10, -1)
        rose.tick()
        self.assertEqual(6, rose.getQuality())
        self.assertEqual(-2, rose.getDaysRemaining())

    @unittest.skip("conjured")
    def testConjuredOfZeroQuality(self):
        rose = GildedRose("Conjured", 0, 5)
        rose.tick()
        self.assertEqual(0, rose.getQuality())
        self.assertEqual(4, rose.getDaysRemaining())

if __name__ == "__main__":
    unittest.main()
