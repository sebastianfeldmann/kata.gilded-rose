<?php

namespace SebastianFeldmann\Kata;

class GildedRoseTest extends \PHPUnit\Framework\TestCase
{
    public function testNormalBeforeSellDate()
    {
        $rose = new GildedRose('normal', 10, 5);
        $rose->tick();

        $this->assertEquals(9, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testNormalOnSellDate()
    {
        $rose = new GildedRose('normal', 10, 0);
        $rose->tick();

        $this->assertEquals(8, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testNormalAfterSellDate()
    {
        $rose = new GildedRose('normal', 10, -1);
        $rose->tick();

        $this->assertEquals(8, $rose->getQuality());
        $this->assertEquals(-2, $rose->getDaysRemaining());
    }

    public function testNormalOfZeroQuality()
    {
        $rose = new GildedRose('normal', 0, 5);
        $rose->tick();

        $this->assertEquals(0, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testBrieBeforeSellDate()
    {
        $rose = new GildedRose('Aged Brie', 10, 5);
        $rose->tick();

        $this->assertEquals(11, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testBrieBeforeSellDateWithMaxQuality()
    {
        $rose = new GildedRose('Aged Brie', 50, 5);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testBrieOnSellDate()
    {
        $rose = new GildedRose('Aged Brie', 10, 0);
        $rose->tick();

        $this->assertEquals(12, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testBrieOnSellDateNearMaxQuality()
    {
        $rose = new GildedRose('Aged Brie', 49, 0);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testBrieOnSellDateWithMaxQuality()
    {
        $rose = new GildedRose('Aged Brie', 50, 0);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testBrieAfterSellDate()
    {
        $rose = new GildedRose('Aged Brie', 40, -1);
        $rose->tick();

        $this->assertEquals(42, $rose->getQuality());
        $this->assertEquals(-2, $rose->getDaysRemaining());
    }

    public function testBrieAfterSellDateWithMaxQuality()
    {
        $rose = new GildedRose('Aged Brie', 50, -1);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(-2, $rose->getDaysRemaining());
    }

    public function testSulfurasBeforeSellDate()
    {
        $rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, 5);
        $rose->tick();

        $this->assertEquals(10, $rose->getQuality());
        $this->assertEquals(5, $rose->getDaysRemaining());
    }

    public function testSulfurasOnSellDate()
    {
        $rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, 0);
        $rose->tick();

        $this->assertEquals(10, $rose->getQuality());
        $this->assertEquals(0, $rose->getDaysRemaining());
    }

    public function testSulfurasAfterSellDate()
    {
        $rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, -1);
        $rose->tick();

        $this->assertEquals(10, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testBackstageLongBeforeSellDate()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 20);
        $rose->tick();

        $this->assertEquals(11, $rose->getQuality());
        $this->assertEquals(19, $rose->getDaysRemaining());
    }

    public function testBackstageMediumCloseToSellDateUpperBound()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 10);
        $rose->tick();

        $this->assertEquals(12, $rose->getQuality());
        $this->assertEquals(9, $rose->getDaysRemaining());
    }

    public function testBackstageMediumCloseToSellDateUpperBoundAtMaxQuality()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 10);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(9, $rose->getDaysRemaining());
    }

    public function testBackstageMediumCloseToSellDateLowerBound()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 6);
        $rose->tick();

        $this->assertEquals(12, $rose->getQuality());
        $this->assertEquals(5, $rose->getDaysRemaining());
    }

    public function testBackstageMediumCloseToSellDateLowerBoundAtMaxQuality()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 6);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(5, $rose->getDaysRemaining());
    }

    public function testBackstageVeryCloseToSellDateUpperBound()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 5);
        $rose->tick();

        $this->assertEquals(13, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testBackstageVeryCloseToSellDateUpperBoundAtMaxQuality()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 5);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }

    public function testBackstageVeryCloseToSellDateLowerBound()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 1);
        $rose->tick();

        $this->assertEquals(13, $rose->getQuality());
        $this->assertEquals(0, $rose->getDaysRemaining());
    }

    public function testBackstageVeryCloseToSellDateLowerBoundAtMaxQuality()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 1);
        $rose->tick();

        $this->assertEquals(50, $rose->getQuality());
        $this->assertEquals(0, $rose->getDaysRemaining());
    }

    public function testBackstageOnSellDate()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 0);
        $rose->tick();

        $this->assertEquals(0, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testBackstageAfterSellDate()
    {
        $rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, -1);
        $rose->tick();

        $this->assertEquals(0, $rose->getQuality());
        $this->assertEquals(-2, $rose->getDaysRemaining());
    }

    public function testConjuredBeforeSellDate()
    {
        $this->markTestSkipped();

        $rose = new GildedRose('Conjured', 10, 10);
        $rose->tick();

        $this->assertEquals(8, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }

    public function testConjuredOnSellDate()
    {
        $this->markTestSkipped();

        $rose = new GildedRose('Conjured', 10, 0);
        $rose->tick();

        $this->assertEquals(6, $rose->getQuality());
        $this->assertEquals(-1, $rose->getDaysRemaining());
    }


    public function testConjuredAfterSellDate()
    {
        $this->markTestSkipped();

        $rose = new GildedRose('Conjured', 10, -1);
        $rose->tick();

        $this->assertEquals(6, $rose->getQuality());
        $this->assertEquals(-2, $rose->getDaysRemaining());
    }

    public function testConjuredOfZeroQuality()
    {
        $this->markTestSkipped();

        $rose = new GildedRose('Conjured', 0, 5);
        $rose->tick();

        $this->assertEquals(0, $rose->getQuality());
        $this->assertEquals(4, $rose->getDaysRemaining());
    }
}
