import UIKit
import XCTest

class GildedRose
{
    /**
     * Type of current rose
     */
    private var name: String

    /**
     * Current quality
     */
    private var quality: Int

    /**
     * Remaining days till sell date
     */
    private var sellIn: Int

    /**
     * main.java.GildedRose constructor
     */
    init(name: String, quality: Int, sellIn: Int)
    {
        self.name    = name
        self.quality = quality
        self.sellIn  = sellIn
    }

    /**
     * Return current quality
     */
    public func getQuality() -> Int
    {
        return quality;
    }

    /**
     * Return days remaining till sell date
     */
    public func getDaysRemaining() -> Int
    {
        return sellIn;
    }

    /**
     * Executes a day tick
     */
    public func tick()
    {
        if (name != "Aged Brie" && name != "Backstage passes to a TAFKAL80ETC concert") {
            if (quality > 0) {
                if (name != "Sulfuras, Hand of Ragnaros") {
                    quality -= 1;
                }
            }
        } else {
            if (quality < 50) {
                quality += 1;
                if (name == "Backstage passes to a TAFKAL80ETC concert") {
                    if (sellIn < 11) {
                        if (quality < 50) {
                            quality += 1;
                        }
                    }
                    if (sellIn < 6) {
                        if (quality < 50) {
                            quality += 1;
                        }
                    }
                }

            }
        }
        if (name != "Sulfuras, Hand of Ragnaros") {
            sellIn -= 1;
        }
        if (sellIn < 0) {
            if (name != "Aged Brie") {
                if (name != "Backstage passes to a TAFKAL80ETC concert") {
                    if (quality > 0) {
                        if (name != "Sulfuras, Hand of Ragnaros") {
                            quality -= 1;
                        }
                    }
                } else {
                    quality = quality - quality;
                }
            } else {
                if (quality < 50) {
                    quality += 1;
                }
            }
        }
    }
}

public class GildedRoseTest: XCTestCase {

    func testNormalBeforeSellDate() {
        let rose = GildedRose(name: "normal", quality: 10, sellIn: 5);
        rose.tick();

        assert(9 == rose.getQuality());
        assert(4 == rose.getDaysRemaining());
    }

    func testNormalOnSellDate() {
        let rose = GildedRose(name: "normal", quality: 10, sellIn: 0);
        rose.tick();

        assert(8 == rose.getQuality());
        assert(-1 == rose.getDaysRemaining());
    }

    func testNormalAfterSellDate() {
        let rose = GildedRose(name: "normal", quality: 10, sellIn: -1);
        rose.tick();

        assert(8 == rose.getQuality());
        assert(-2 == rose.getDaysRemaining());
    }


    func testNormalOfZeroQuality() {
        let rose = GildedRose(name: "normal", quality: 0, sellIn: 5);
        rose.tick();

        assert(0 == rose.getQuality());
        assert(4 == rose.getDaysRemaining());
    }


    func testBrieBeforeSellDate() {
        let rose = GildedRose(name: "Aged Brie", quality: 10, sellIn: 5);
        rose.tick();

        assert(11 == rose.getQuality());
        assert(4 == rose.getDaysRemaining());
    }


    func testBrieBeforeSellDateWithMaxQuality() {
        let rose = GildedRose(name: "Aged Brie", quality: 50, sellIn: 5);
        rose.tick();

        assert(50 == rose.getQuality());
        assert(4 == rose.getDaysRemaining());
    }


    func testBrieOnSellDate() {
        let rose = GildedRose(name: "Aged Brie", quality: 10, sellIn: 0);
        rose.tick();

        assert(12 == rose.getQuality());
        assert(-1 == rose.getDaysRemaining());
    }


    func testBrieOnSellDateNearMaxQuality() {
        let rose = GildedRose(name: "Aged Brie", quality: 49, sellIn: 0);
        rose.tick();

        assert(50 == rose.getQuality());
        assert(-1 == rose.getDaysRemaining());
    }


    func testBrieOnSellDateWithMaxQuality() {
        let rose = GildedRose(name: "Aged Brie", quality: 50, sellIn: 0);
        rose.tick();

        assert(50 == rose.getQuality());
        assert(-1 == rose.getDaysRemaining());
    }


    func testBrieAfterSellDate() {
        let rose = GildedRose(name: "Aged Brie", quality: 40, sellIn: -1);
        rose.tick();

        assert(42 == rose.getQuality());
        assert(-2 == rose.getDaysRemaining());
    }


    func testBrieAfterSellDateWithMaxQuality() {
        let rose = GildedRose(name: "Aged Brie", quality: 50, sellIn: -1);
        rose.tick();

        assert(50 == rose.getQuality());
        assert(-2 == rose.getDaysRemaining());
    }


    func testSulfurasBeforeSellDate() {
        let rose = GildedRose(name: "Sulfuras, Hand of Ragnaros", quality: 10, sellIn: 5);
        rose.tick();

        assert(10 == rose.getQuality());
        assert(5 == rose.getDaysRemaining());
    }


    func testSulfurasOnSellDate() {
        let rose = GildedRose(name: "Sulfuras, Hand of Ragnaros", quality: 10, sellIn: 0);
        rose.tick();

        assert(10 == rose.getQuality());
        assert(0 == rose.getDaysRemaining());
    }


    func testSulfurasAfterSellDate() {
        let rose = GildedRose(name: "Sulfuras, Hand of Ragnaros", quality: 10, sellIn: -1);
        rose.tick();

        assert(10 == rose.getQuality());
        assert(-1 == rose.getDaysRemaining());
    }

    func testBackstageLongBeforeSellDate() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 20)
        rose.tick()

        assert(11 == rose.getQuality())
        assert(19 == rose.getDaysRemaining())
    }


    func testBackstageMediumCloseToSellDateUpperBound() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 10)
        rose.tick()

        assert(12 == rose.getQuality())
        assert(9 == rose.getDaysRemaining())
    }


    func testBackstageMediumCloseToSellDateUpperBoundAtMaxQuality() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 50, sellIn: 10)
        rose.tick()

        assert(50 == rose.getQuality())
        assert(9 == rose.getDaysRemaining())
    }


    func testBackstageMediumCloseToSellDateLowerBound() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 6)
        rose.tick()

        assert(12 == rose.getQuality())
        assert(5 == rose.getDaysRemaining())
    }


    func testBackstageMediumCloseToSellDateLowerBoundAtMaxQuality() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 50, sellIn: 6)
        rose.tick()

        assert(50 == rose.getQuality())
        assert(5 == rose.getDaysRemaining())
    }


    func testBackstageVeryCloseToSellDateUpperBound() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 5)
        rose.tick()

        assert(13 == rose.getQuality())
        assert(4 == rose.getDaysRemaining())
    }


    func testBackstageVeryCloseToSellDateUpperBoundAtMaxQuality() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 50, sellIn: 5)
        rose.tick()

        assert(50 == rose.getQuality())
        assert(4 == rose.getDaysRemaining())
    }


    func testBackstageVeryCloseToSellDateLowerBound() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 1)
        rose.tick()

        assert(13 == rose.getQuality())
        assert(0 == rose.getDaysRemaining())
    }


    func testBackstageVeryCloseToSellDateLowerBoundAtMaxQuality() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 50, sellIn: 1)
        rose.tick()

        assert(50 == rose.getQuality())
        assert(0 == rose.getDaysRemaining())
    }

    func testBackstageOnSellDate() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: 0)
        rose.tick()

        assert(0 == rose.getQuality())
        assert(-1 == rose.getDaysRemaining())
    }


    func testBackstageAfterSellDate() {
        let rose = GildedRose(name: "Backstage passes to a TAFKAL80ETC concert", quality: 10, sellIn: -1)
        rose.tick()

        assert(0 == rose.getQuality())
        assert(-2 == rose.getDaysRemaining())
    }
    /*
     func testConjuredBeforeSellDate() {
     let rose = GildedRose(name: "Conjured", quality: 10, sellIn: 10)
     rose.tick()

     assert(8 == rose.getQuality())
     assert(9 == rose.getDaysRemaining())
     }

     func testConjuredOnSellDate() {
     let rose = GildedRose(name: "Conjured", quality: 10, sellIn: 0)
     rose.tick()

     assert(6 == rose.getQuality())
     assert(-1 == rose.getDaysRemaining())
     }

     func testConjuredAfterSellDate() {
     let rose = GildedRose(name: "Conjured", quality: 10, sellIn: -1)
     rose.tick()

     assert(6 == rose.getQuality())
     assert(-2 == rose.getDaysRemaining())
     }

     func testConjuredOfZeroQuality() {
     let rose = GildedRose(name: "Conjured", quality: 0, sellIn: 5)
     rose.tick()

     assert(0 == rose.getQuality())
     assert(4 == rose.getDaysRemaining())
     }
     */
}

GildedRoseTest.defaultTestSuite.run()
