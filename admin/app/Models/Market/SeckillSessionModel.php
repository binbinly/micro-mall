<?php


namespace App\Models\Market;


/**
 * 秒杀场次
 * Class SeckillSessionModel
 * @package App\Models\Market
 */
class SeckillSessionModel extends BaseModel
{
    protected $table = 'sms_seckill_session';

    public static function getAll()
    {
        return self::query()->pluck('name', 'id')->toArray();
    }
}