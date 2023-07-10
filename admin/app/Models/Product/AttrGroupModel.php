<?php


namespace App\Models\Product;


/**
 * 属性分组模型
 * Class AttrGroupModel
 * @package App\Models\Product
 */
class AttrGroupModel extends BaseModel
{
    protected $table = 'pms_attr_group';

    public $timestamps = false;

    public static function getAll(): array
    {
        return self::query()->pluck('name', 'id')->toArray();
    }

    /**
     * 当前分类下的所有分组id
     * @param int $catId
     * @return array
     */
    public static function getGroupIdByCatId(int $catId): array
    {
        return self::query()->where('cat_id', $catId)->pluck('id')->toArray();
    }

    /**
     * 当前分类下的所有分组
     * @param int $catId
     * @return array
     */
    public static function getGroupByCatId(int $catId): array
    {
        return self::query()->where('cat_id', $catId)->pluck('name', 'id')->toArray();
    }
}
