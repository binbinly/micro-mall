<?php


namespace App\Models\Market;


use App\Admin\Common\Format;

/**
 * 秒杀商品
 * Class SeckillSku
 * @package App\Models\Market
 */
class SeckillSkuModel extends BaseModel
{
    protected $table = 'sms_seckill_sku';

    public $timestamps = false;

    public function getPriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setPriceAttribute($value)
    {
        $this->attributes['price'] = Format::amountToPenny($value);
    }
}