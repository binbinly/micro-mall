<?php


namespace App\Admin\Selectable;


use App\Models\Product\BrandModel;
use App\Models\Product\CategoryModel;
use App\Models\Product\SkuModel;
use Encore\Admin\Grid\Filter;
use Encore\Admin\Grid\Selectable;

class Sku extends Selectable
{
    public $model = SkuModel::class;

    public function make()
    {
        $this->column('id', 'ID');
        $this->column('title', '标题');
        $this->column('cover', '封面')->image();
        $this->column('price', '价格')->amount();
        $this->column('attr_value', '规格');

        $this->filter(function (Filter $filter) {
            $filter->equal('cat_id', '分类')->select(CategoryModel::selectOptions());
            $filter->equal('brand_id', '品牌')->select(BrandModel::getAll());
            $filter->like('title', '商品名');
        });
    }
}