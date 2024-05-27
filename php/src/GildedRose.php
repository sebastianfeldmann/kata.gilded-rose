<?php declare(strict_types=1);

namespace GildedRose;

class GildedRose
{
    private string $name;

    private int $quality;

    private int $sellIn;

    public function __construct(string $name, int $quality, int $sellIn)
    {
        $this->name    = $name;
        $this->quality = $quality;
        $this->sellIn  = $sellIn;
    }

    public function getQuality() : int
    {
        return $this->quality;
    }

    public function getDaysRemaining() : int
    {
        return $this->sellIn;
    }

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
