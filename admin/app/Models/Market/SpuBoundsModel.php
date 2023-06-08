<?php


namespace App\Models\Market;


/**
 * 商品spu积分设置
 * Class SpuBoundsModel
 * @package App\Models\Coupon
 * @property int spu_id
 * @property int grow_bounds
 * @property int buy_bounds
 * @property int work
 */
class SpuBoundsModel extends BaseModel
{
    protected $table = 'sms_spu_bounds';

    public $timestamps = false;
}