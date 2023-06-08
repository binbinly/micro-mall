<?php


namespace App\Models\Market;


use App\Admin\Common\Format;

/**
 * 商品满减优惠
 * Class SkuFullReductionModel
 * @package App\Models\Coupon
 */
class SkuFullReductionModel extends BaseModel
{
    protected $table = 'sms_sku_full_reduction';

    public $timestamps = false;

    public function getFullPriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setFullPriceAttribute($value)
    {
        $this->attributes['full_price'] = Format::amountToPenny($value);
    }

    public function getReducePriceAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setReducePriceAttribute($value)
    {
        $this->attributes['reduce_price'] = Format::amountToPenny($value);
    }
}