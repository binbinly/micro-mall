<?php


namespace App\Models\Market;


use App\Admin\Common\Format;

/**
 * 优惠券模型
 * Class CouponModel
 * @package App\Models\Coupon
 */
class CouponModel extends BaseModel
{
    const TYPE_ALL = 0;
    const TYPE_MEMBER = 1;
    const TYPE_SHOP = 2;
    const TYPE_REGISTER = 3;

    const USE_TYPE_ALL = 0;
    const USE_TYPE_CAT = 1;
    const USE_TYPE_PRODUCT = 2;

    protected $table = 'sms_coupon';

    public static $typeLabel = [
        self::TYPE_ALL => '全场赠券',
        self::TYPE_MEMBER => '会员赠券',
        self::TYPE_SHOP => '购物赠券',
        self::TYPE_REGISTER => '注册赠券'
    ];

    public static $useTypeLabel = [
        self::USE_TYPE_ALL => '全场通用',
        self::USE_TYPE_CAT => '指定分类',
        self::USE_TYPE_PRODUCT => '指定商品'
    ];

    protected $dates = [
        'created_at',
        'updated_at',
        'deleted_at',
        'start_at',
        'end_at',
        'enable_start_at',
        'enable_end_at'
    ];

    public function getAmountAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setAmountAttribute($value)
    {
        $this->attributes['amount'] = Format::amountToPenny($value);
    }

    public function getMinPointAttribute($value)
    {
        return Format::amountToYuan($value);
    }

    public function setMinPointAttribute($value)
    {
        $this->attributes['min_point'] = Format::amountToPenny($value);
    }

}