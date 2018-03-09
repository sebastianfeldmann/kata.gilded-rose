<?php

namespace SebastianFeldmann\Kata\GildedRose;

class GildedRose
{
    /**
     * Type of current rose.
     *
     * @var string
     */
    private $name;

    /**
     * Current quality.
     *
     * @var int
     */
    private $quality;

    /**
     * Remaining days till sell date.
     *
     * @var int
     */
    private $sellIn;

    /**
     * GildedRose constructor.
     *
     * @param string $name
     * @param int    $quality
     * @param int    $sellIn
     */
    public function __construct(string $name, int $quality, int $sellIn)
    {
        $this->name    = $name;
        $this->quality = $quality;
        $this->sellIn  = $sellIn;
    }

    /**
     * Return current quality.
     *
     * @return int
     */
    public function getQuality() : int
    {
        return $this->quality;
    }

    /**
     * Return days remaining till sell date.
     *
     * @return int
     */
    public function getDaysRemaining() : int
    {
        return $this->sellIn;
    }

    /**
     * Executes a day tick.
     */
    public function tick() : void
    {
        if ($this->name != 'Aged Brie' && $this->name != 'Backstage passes to a TAFKAL80ETC concert') {
            if ($this->quality > 0) {
                if ($this->name != 'Sulfuras, Hand of Ragnaros') {
                    $this->quality -= 1;
                }
            }
        } else {
            if ($this->quality < 50) {
                $this->quality += 1;
                if ($this->name == 'Backstage passes to a TAFKAL80ETC concert') {
                    if ($this->sellIn < 11) {
                        if ($this->quality < 50) {
                            $this->quality += 1;
                        }
                    }
                    if ($this->sellIn < 6) {
                        if ($this->quality < 50) {
                            $this->quality += 1;
                        }
                    }
                }

            }
        }
        if ($this->name != 'Sulfuras, Hand of Ragnaros') {
            $this->sellIn -= 1;
        }
        if ($this->sellIn < 0) {
            if ($this->name != 'Aged Brie') {
                if ($this->name != 'Backstage passes to a TAFKAL80ETC concert') {
                    if ($this->quality > 0) {
                        if ($this->name != 'Sulfuras, Hand of Ragnaros') {
                            $this->quality -= 1;
                        }
                    }
                } else {
                    $this->quality = $this->quality - $this->quality;
                }
            } else {
                if ($this->quality < 50) {
                    $this->quality += 1;
                }
            }
        }
    }
}
