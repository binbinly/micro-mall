<?php


namespace App\Models\Market;


use App\Admin\Common\CastAmount;

/**
 * 商品满减优惠
 * Class SkuFullReductionModel
 * @package App\Models\Coupon
 */
class SkuFullReductionModel extends BaseModel
{
    protected $table = 'sms_sku_full_reduction';

    public $timestamps = false;

    protected $casts = [
        'full_price' => CastAmount::class,
        'reduce_price' => CastAmount::class
    ];
}
