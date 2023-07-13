<?php

namespace App\Admin\Controllers\Product;

use App\Admin\Common\Search;
use App\Models\Product\BrandModel;
use App\Models\Product\CategoryModel;
use App\Models\Product\SkuModel;
use Encore\Admin\Grid;

class SkuController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '商品管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SkuModel());
        $grid->model()->orderBy('id', 'desc');
        $this->disableGridExport($grid);

        $grid->column('id', 'ID');
        $grid->column('spu_id', 'spu_id');
        $this->gridCat($grid);
        $grid->column('brand_id', '品牌')->using(BrandModel::getAll())->label('info');
        $grid->column('name', '商品名')->width(200);
        $grid->column('desc', '描述')->width(200);
        $grid->column('cover', '封面')->image('', 120);
        $grid->column('title', '标题')->width(200);
        $grid->column('subtitle', '副标题')->width(200);
        $grid->column('price', '价格')->amount();
        $grid->column('attr_value', '规格');
        $grid->column('sale_count', '销量');

        return $grid;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('cat_id', '分类')->select(CategoryModel::selectOptions());
        $filter->equal('brand_id', '品牌')->select(BrandModel::getAll());
        $filter->like('title', '商品名');
    }
}
