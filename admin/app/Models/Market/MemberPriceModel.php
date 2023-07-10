<?php


namespace App\Models\Market;


use App\Admin\Common\CastAmount;

/**
 * 商品会员价格模型
 * Class MemberPriceModel
 * @package App\Models\Coupon
 */
class MemberPriceModel extends BaseModel
{
    protected $table = 'sms_member_price';

    public $timestamps = false;

    protected $casts = [
        'price' => CastAmount::class
    ];
}
