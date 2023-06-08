<?php


namespace App\Models\Product;


/**
 * sku模型
 * Class SkuModel
 * @package App\Models\Product
 * @property int id
 * @property int spu_id
 * @property int cat_id
 * @property int brand_id
 * @property string name
 * @property string desc
 * @property string cover
 * @property string title
 * @property string subtitle
 * @property int price
 * @property int sale_count
 */
class SkuModel extends BaseModel
{
    protected $table = 'pms_sku';

    public $timestamps = false;
}