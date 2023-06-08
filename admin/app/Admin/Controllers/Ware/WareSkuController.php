<?php

namespace App\Admin\Controllers\Ware;

use App\Admin\Common\Search;
use App\Admin\Controllers\BaseController;
use App\Models\Ware\WareModel;
use App\Models\Ware\WareSkuModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class WareSkuController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'WareSkuModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new WareSkuModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('sku_id', '商品id');
        $grid->column('ware_id', '仓库')->using(WareModel::getAll());
        $grid->column('sku_name', '商品名');
        $grid->column('stock', '库存');
        $grid->column('stock_lock', '锁定库存');
        $this->setGridTimeView($grid);
        $this->search($grid);

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new WareSkuModel());

        $form->text('sku_id', '商品id');
        $form->select('ware_id', '仓库')->options(WareModel::getAll());
        $form->text('sku_name', '商品名');
        $form->number('stock', '库存');
        $form->number('stock_lock', '锁定库存');

        return $form;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('sku_id', '商品id');
    }
}
