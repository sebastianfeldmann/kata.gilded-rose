import { GildedRose } from '../src';

describe('normal rose', () => {
    test('before sell date', () => {
        const rose = new GildedRose('normal', 10, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(9);
        expect(rose.getDaysRemaining()).toBe(4);
    });

    test('on sell date', () => {
        const rose = new GildedRose('normal', 10, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(8);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test('after sell date', () => {
        const rose = new GildedRose('normal', 10, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(8);
        expect(rose.getDaysRemaining()).toBe(-2);
    });

    test('zero quality', () => {
        const rose = new GildedRose('normal', 0, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(0);
        expect(rose.getDaysRemaining()).toBe(4);
    });
});

describe('brie rose', () => {
    test('before sell date', () => {
        const rose = new GildedRose('Aged Brie', 10, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(11);
        expect(rose.getDaysRemaining()).toBe(4);
    });

    test('before sell date with max quality', () => {
        const rose = new GildedRose('Aged Brie', 50, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(4);
    });

    test('on sell date', () => {
        const rose = new GildedRose('Aged Brie', 10, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(12);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test('on sell date near max quality', () => {
        const rose = new GildedRose('Aged Brie', 49, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test('on sell date with max quality', () => {
        const rose = new GildedRose('Aged Brie', 50, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test('after sell date', () => {
        const rose = new GildedRose('Aged Brie', 40, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(42);
        expect(rose.getDaysRemaining()).toBe(-2);
    });

    test('after sell date with max quality', () => {
        const rose = new GildedRose('Aged Brie', 50, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(-2);
    });
});

describe('sulfuras', () => {
    test('before sell date', () => {
        const rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(10);
        expect(rose.getDaysRemaining()).toBe(5);
    });

    test('on sell date', () => {
        const rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(10);
        expect(rose.getDaysRemaining()).toBe(0);
    });

    test('after sell date', () => {
        const rose = new GildedRose('Sulfuras, Hand of Ragnaros', 10, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(10);
        expect(rose.getDaysRemaining()).toBe(-1);
    });
});

describe('backstage', () => {
    test('long before sell date', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 20);

        rose.tick();

        expect(rose.getQuality()).toBe(11);
        expect(rose.getDaysRemaining()).toBe(19);
    });

    test('medium close to sell date upper bound', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 10);

        rose.tick();

        expect(rose.getQuality()).toBe(12);
        expect(rose.getDaysRemaining()).toBe(9);
    });

    test('medium close to sell date upper bound at max quality', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 10);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(9);
    });

    test('medium close to sell date lower bound', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 6);

        rose.tick();

        expect(rose.getQuality()).toBe(12);
        expect(rose.getDaysRemaining()).toBe(5);
    });

    test('medium close to sell date lower bound at max quality', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 6);

        rose.tick();

        expect(rose.getQuality()).toBe(12);
        expect(rose.getDaysRemaining()).toBe(5);
    });

    test('very close to sell date upper bound', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(13);
        expect(rose.getDaysRemaining()).toBe(4);
    });

    test('very close to sell date upper bound at max quality', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(4);
    });

    test('very close to sell date lower bound', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 1);

        rose.tick();

        expect(rose.getQuality()).toBe(13);
        expect(rose.getDaysRemaining()).toBe(0);
    });

    test('very close to sell date lower bound at max quality', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 50, 1);

        rose.tick();

        expect(rose.getQuality()).toBe(50);
        expect(rose.getDaysRemaining()).toBe(0);
    });

    test('on sell date', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(0);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test('after sell date', () => {
        const rose = new GildedRose('Backstage passes to a TAFKAL80ETC concert', 10, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(0);
        expect(rose.getDaysRemaining()).toBe(-2);
    });
});

describe('Conjured', () => {
    test.skip('before sell date', () => {
        const rose = new GildedRose('Conjured', 10, 10);

        rose.tick();

        expect(rose.getQuality()).toBe(8);
        expect(rose.getDaysRemaining()).toBe(9);
    });

    test.skip('on sell date', () => {
        const rose = new GildedRose('Conjured', 10, 0);

        rose.tick();

        expect(rose.getQuality()).toBe(6);
        expect(rose.getDaysRemaining()).toBe(-1);
    });

    test.skip('after sell date', () => {
        const rose = new GildedRose('Conjured', 10, -1);

        rose.tick();

        expect(rose.getQuality()).toBe(6);
        expect(rose.getDaysRemaining()).toBe(-2);
    });

    test.skip('zero quality', () => {
        const rose = new GildedRose('Conjured', 0, 5);

        rose.tick();

        expect(rose.getQuality()).toBe(0);
        expect(rose.getDaysRemaining()).toBe(4);
    });
});
