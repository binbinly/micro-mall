<?php


namespace App\Models\Market;


use App\Admin\Common\Format;

/**
 * 商品会员价格模型
 * Class MemberPriceModel
 * @package App\Models\Coupon
 */
class MemberPriceModel extends BaseModel
{
    protected $table = 'sms_member_price';

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