<?php


namespace App\Models\Market;


/**
 * 秒杀活动
 * Class SeckillActivityModel
 * @package App\Models\Market
 */
class SeckillActivityModel extends BaseModel
{
    protected $table = 'sms_seckill_activity';

    public static function getAll()
    {
        return self::query()->pluck('title', 'id')->toArray();
    }
}