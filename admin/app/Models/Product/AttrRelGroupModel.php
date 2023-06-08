<?php


namespace App\Models\Product;

/**
 * 属性关联属性分组模型
 * Class AttrRelGroupModel
 * @package App\Models\Product
 */
class AttrRelGroupModel extends BaseModel
{
    protected $table = 'pms_attr_rel_group';

    public $timestamps = false;

    public function attr()
    {
        return $this->hasOne(AttrModel::class, 'id', 'attr_id');
    }

    /**
     * 分组关联的所有属性
     * @param array $ids
     * @return array
     */
    public static function getAttrByGroupIds(array $ids){
        if (count($ids) == 0) return [];
        return self::query()->whereIn('group_id', $ids)->pluck('attr_id')->toArray();
    }

    /**
     * 获取当前分组下的所有规格属性
     * @param $ids
     * @return array
     */
    public static function getBaseAttrs($ids){
        return self::query()->whereIn('group_id', $ids)->with('attr')
            ->whereHas('attr', function ($query){
                $query->where('type', AttrModel::TYPE_BASE);
            })->get()->toArray();
    }
}