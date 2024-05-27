type Name =
    | 'normal'
    | 'Aged Brie'
    | 'Backstage passes to a TAFKAL80ETC concert'
    | 'Sulfuras, Hand of Ragnaros'
    | 'Conjured';

export class GildedRose {
    /**
     * Type of current rose
     */
    private name: Name;

    /**
     * Current quality
     */
    private quality: number;

    /**
     * Remaining days till sell date
     */
    private sellIn: number;

    public constructor(name: Name, quality: number, sellIn: number) {
        this.name = name;
        this.quality = quality;
        this.sellIn = sellIn;
    }

    /**
     * Return current quality
     */
    public getQuality(): number {
        return this.quality;
    }

    /**
     * Return days remaining till sell date
     */
    public getDaysRemaining(): number {
        return this.sellIn;
    }

    /**
     * Executes a day tick
     */
    public tick(): void {
        if (this.name != 'Aged Brie' && this.name != 'Backstage passes to a TAFKAL80ETC concert') {
            if (this.quality > 0) {
                if (this.name != 'Sulfuras, Hand of Ragnaros') {
                    this.quality -= 1;
                }
            }
        } else {
            if (this.quality < 50) {
                this.quality += 1;
                if (this.name == 'Backstage passes to a TAFKAL80ETC concert') {
                    if (this.sellIn < 11) {
                        if (this.quality < 50) {
                            this.quality += 1;
                        }
                    }
                    if (this.sellIn < 6) {
                        if (this.quality < 50) {
                            this.quality += 1;
                        }
                    }
                }
            }
        }
        if (this.name != 'Sulfuras, Hand of Ragnaros') {
            this.sellIn -= 1;
        }
        if (this.sellIn < 0) {
            if (this.name != 'Aged Brie') {
                if (this.name != 'Backstage passes to a TAFKAL80ETC concert') {
                    if (this.quality > 0) {
                        if (this.name != 'Sulfuras, Hand of Ragnaros') {
                            this.quality -= 1;
                        }
                    }
                } else {
                    this.quality = this.quality - this.quality;
                }
            } else {
                if (this.quality < 50) {
                    this.quality += 1;
                }
            }
        }
    }
}
