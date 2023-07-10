<?php


namespace App\Models\Market;


use App\Admin\Common\CastAmount;

/**
 * 秒杀商品
 * Class SeckillSku
 * @package App\Models\Market
 */
class SeckillSkuModel extends BaseModel
{
    protected $table = 'sms_seckill_sku';

    public $timestamps = false;

    protected $casts = [
        'price' => CastAmount::class
    ];
}
