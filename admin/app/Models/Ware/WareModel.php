<?php


namespace App\Models\Ware;


class WareModel extends BaseModel
{
    protected $table = 'wms_ware';

    public static function getAll(): array
    {
        return self::query()->pluck('name', 'id')->toArray();
    }
}
