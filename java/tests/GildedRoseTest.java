import org.junit.Ignore;
import org.junit.Test;

import static org.junit.Assert.*;

public class GildedRoseTest {

    @Test
    public void testNormalBeforeSellDate() {
        GildedRose rose = new GildedRose("normal", 10, 5);
        rose.tick();

        assertEquals(9, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testNormalOnSellDate() {
        GildedRose rose = new GildedRose("normal", 10, 0);
        rose.tick();

        assertEquals(8, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testNormalAfterSellDate() {
        GildedRose rose = new GildedRose("normal", 10, -1);
        rose.tick();

        assertEquals(8, rose.getQuality());
        assertEquals(-2, rose.getDaysRemaining());
    }

    @Test
    public void testNormalOfZeroQuality() {
        GildedRose rose = new GildedRose("normal", 0, 5);
        rose.tick();

        assertEquals(0, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testBrieBeforeSellDate() {
        GildedRose rose = new GildedRose("Aged Brie", 10, 5);
        rose.tick();

        assertEquals(11, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testBrieBeforeSellDateWithMaxQuality() {
        GildedRose rose = new GildedRose("Aged Brie", 50, 5);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testBrieOnSellDate() {
        GildedRose rose = new GildedRose("Aged Brie", 10, 0);
        rose.tick();

        assertEquals(12, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testBrieOnSellDateNearMaxQuality() {
        GildedRose rose = new GildedRose("Aged Brie", 49, 0);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testBrieOnSellDateWithMaxQuality() {
        GildedRose rose = new GildedRose("Aged Brie", 50, 0);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testBrieAfterSellDate() {
        GildedRose rose = new GildedRose("Aged Brie", 40, -1);
        rose.tick();

        assertEquals(42, rose.getQuality());
        assertEquals(-2, rose.getDaysRemaining());
    }

    @Test
    public void testBrieAfterSellDateWithMaxQuality() {
        GildedRose rose = new GildedRose("Aged Brie", 50, -1);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(-2, rose.getDaysRemaining());
    }

    @Test
    public void testSulfurasBeforeSellDate() {
        GildedRose rose = new GildedRose("Sulfuras, Hand of Ragnaros", 10, 5);
        rose.tick();

        assertEquals(10, rose.getQuality());
        assertEquals(5, rose.getDaysRemaining());
    }

    @Test
    public void testSulfurasOnSellDate() {
        GildedRose rose = new GildedRose("Sulfuras, Hand of Ragnaros", 10, 0);
        rose.tick();

        assertEquals(10, rose.getQuality());
        assertEquals(0, rose.getDaysRemaining());
    }

    @Test
    public void testSulfurasAfterSellDate() {
        GildedRose rose = new GildedRose("Sulfuras, Hand of Ragnaros", 10, -1);
        rose.tick();

        assertEquals(10, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageLongBeforeSellDate() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 20);
        rose.tick();

        assertEquals(11, rose.getQuality());
        assertEquals(19, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageMediumCloseToSellDateUpperBound() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 10);
        rose.tick();

        assertEquals(12, rose.getQuality());
        assertEquals(9, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageMediumCloseToSellDateUpperBoundAtMaxQuality() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 10);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(9, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageMediumCloseToSellDateLowerBound() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 6);
        rose.tick();

        assertEquals(12, rose.getQuality());
        assertEquals(5, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageMediumCloseToSellDateLowerBoundAtMaxQuality() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 6);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(5, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageVeryCloseToSellDateUpperBound() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 5);
        rose.tick();

        assertEquals(13, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageVeryCloseToSellDateUpperBoundAtMaxQuality() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 5);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageVeryCloseToSellDateLowerBound() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 1);
        rose.tick();

        assertEquals(13, rose.getQuality());
        assertEquals(0, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageVeryCloseToSellDateLowerBoundAtMaxQuality() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 50, 1);
        rose.tick();

        assertEquals(50, rose.getQuality());
        assertEquals(0, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageOnSellDate() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, 0);
        rose.tick();

        assertEquals(0, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    public void testBackstageAfterSellDate() {
        GildedRose rose = new GildedRose("Backstage passes to a TAFKAL80ETC concert", 10, -1);
        rose.tick();

        assertEquals(0, rose.getQuality());
        assertEquals(-2, rose.getDaysRemaining());
    }

    @Test
    @Ignore
    public void testConjuredBeforeSellDate() {
        GildedRose rose = new GildedRose("Conjured", 10, 10);
        rose.tick();

        assertEquals(8, rose.getQuality());
        assertEquals(9, rose.getDaysRemaining());
    }

    @Test
    @Ignore
    public void testConjuredOnSellDate() {
        GildedRose rose = new GildedRose("Conjured", 10, 0);
        rose.tick();

        assertEquals(6, rose.getQuality());
        assertEquals(-1, rose.getDaysRemaining());
    }

    @Test
    @Ignore
    public void testConjuredAfterSellDate() {
        GildedRose rose = new GildedRose("Conjured", 10, -1);
        rose.tick();

        assertEquals(6, rose.getQuality());
        assertEquals(-2, rose.getDaysRemaining());
    }

    @Test
    @Ignore
    public void testConjuredOfZeroQuality() {
        GildedRose rose = new GildedRose("Conjured", 0, 5);
        rose.tick();

        assertEquals(0, rose.getQuality());
        assertEquals(4, rose.getDaysRemaining());
    }
}