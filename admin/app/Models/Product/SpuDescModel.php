<?php


namespace App\Models\Product;


/**
 * 商品介绍模型
 * Class SpuDescModel
 * @package App\Models\Product
 * @property int spu_id
 * @property string content
 */
class SpuDescModel extends BaseModel
{
    protected $table = 'pms_spu_desc';

    public $timestamps = false;
}