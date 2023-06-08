<?php


namespace App\Models\Product;

use Encore\Admin\Traits\ModelTree;

/**
 * 商品分类
 * Class CategoryModel
 * @package App\Models\Goods
 */
class CategoryModel extends BaseModel
{
    use ModelTree;

    public $timestamps = false;

    protected $table = 'pms_category';

    public function __construct(array $attributes = [])
    {
        parent::__construct($attributes);

        $this->setParentColumn('parent_id');
        $this->setOrderColumn('sort');
        $this->setTitleColumn('name');
    }

    /**
     * 所有一级分类
     * @return array
     */
    public static function parentAll()
    {
        return self::query()->where('parent_id', 0)->pluck('name', 'id')->toArray();
    }

    public static function getAll()
    {
        return self::query()->pluck('name', 'id')->toArray();
    }

}