<?php


namespace App\Models\Product;


/**
 * 属性模型
 * Class AttrModel
 * @package App\Models\Product
 */
class AttrModel extends BaseModel
{
    const TYPE_SALE = 0;
    const TYPE_BASE = 1;
    const TYPE_COMMON = 2;

    public static $typeLabel = [
        self::TYPE_SALE => '销售属性',
        self::TYPE_BASE => '基本属性',
//        self::TYPE_COMMON => '公共属性' //既是销售属性又是基本属性
    ];

    protected $table = 'pms_attr';

    public function relGroup()
    {
        return $this->hasOne(AttrRelGroupModel::class, 'attr_id');
    }

    /**
     * 获取分类下的所有属性map结构
     * @param $catId
     * @param int $type
     * @return array
     */
    public static function getAttrMapByCatID($catId, $type = self::TYPE_BASE) {
        return self::query()->when($catId>0, function ($query) use($catId) {
            return $query->where('cat_id', $catId);
        })->where('type', $type)->pluck('name', 'id')->toArray();
    }

    /**
     * 获取分类下的所有属性列表
     * @param $catId
     * @param int $type
     * @return array
     */
    public static function getAttrByCatID($catId, $type = self::TYPE_BASE) {
        return self::query()->where('cat_id', $catId)->where('type', $type)->get()->toArray();
    }
}