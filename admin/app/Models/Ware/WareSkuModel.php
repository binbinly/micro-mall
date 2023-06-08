<?php


namespace App\Models\Ware;


/**
 * 商品库存
 * Class WareSkuModel
 * @package App\Models\Ware
 */
class WareSkuModel extends BaseModel
{
    /**
     * 可以被批量赋值的属性。
     *
     * @var array
     */
    protected $fillable = ['sku_id', 'ware_id', 'stock'];

    protected $table = 'wms_ware_sku';
}