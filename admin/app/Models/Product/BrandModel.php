<?php


namespace App\Models\Product;

/**
 * 品牌模型
 * Class BrandModel
 * @package App\Models\Product
 */
class BrandModel extends BaseModel
{
    protected $table = 'pms_brand';

    public function relCat()
    {
        return $this->hasOne(CategoryRelBrandModel::class, 'brand_id');
    }

    public static function getAll()
    {
        return self::query()->pluck('name', 'id')->toArray();
    }
}