<?php


namespace App\Models\Market;


/**
 * 优惠券会员领取
 * Class CouponMemberModel
 * @package App\Models\Coupon
 */
class CouponMemberModel extends BaseModel
{
    const GET_TYPE_ADMIN = 1;
    const GET_TYPE_MY = 2;

    //使用状态[0->未使用；1->已使用；2->已过期]
    const STATUS_NOT = 0;
    const STATUS_USE = 1;
    const STATUS_EXPIRE = 2;

    protected $table = 'sms_coupon_member';

    public static array $getTypeLabel = [
        self::GET_TYPE_ADMIN => '后台赠送',
        self::GET_TYPE_MY => '主动领取'
    ];

    public static array $statusLabel = [
        self::STATUS_NOT => '未使用',
        self::STATUS_USE => '已使用',
        self::STATUS_EXPIRE => '已过期'
    ];

    protected $casts = [
        'created_at' => 'datetime',
        'updated_at' => 'datetime',
        'used_at' => 'datetime'
    ];
}
