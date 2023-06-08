<?php


namespace App\Models\Product;


/**
 * 属性值模型
 * Class AttrValueModel
 * @package App\Models\Product
 */
class AttrValueModel extends BaseModel
{
    protected $table = 'pms_attr_value';

    public function attr()
    {
        return $this->hasOne(AttrModel::class, 'id', 'attr_id');
    }
}