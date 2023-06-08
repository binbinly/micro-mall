<?php

namespace App\Admin\Controllers\Ware;

use App\Admin\Actions\Post\PurchaseSubmit;
use App\Admin\Common\Search;
use App\Admin\Controllers\BaseController;
use App\Models\Ware\PurchaseDetailModel;
use App\Models\Ware\WareModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class PurchaseDetailController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '采购需求';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new PurchaseDetailModel());
        $grid->model()->orderBy('id', 'desc');
        $grid->disableExport();

        $grid->batchActions(function ($batch) {
            $batch->add(new PurchaseSubmit());
        });

        $grid->column('id', 'ID');
        $grid->column('purchase_id', '采购单ID');
        $grid->column('sku_id', '商品id');
        $grid->column('sku_num', '商品数量');
        $grid->column('sku_price', '商品价格');
        $grid->column('ware_id', '仓库ID');
        $grid->column('status', '状态')->using(PurchaseDetailModel::$statusLabel)->filter(PurchaseDetailModel::$statusLabel)->label();
        $this->setGridTimeView($grid);

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new PurchaseDetailModel());

        $form->number('sku_id', '采购商品id');
        $form->number('sku_num', '采购数量');
        $form->select('ware_id', '仓库')->options(WareModel::getAll());
        $form->select('status', '状态')->options(PurchaseDetailModel::$statusLabel);

        return $form;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('sku_id', '商品id');
        $filter->equal('ware_id', '仓库')->select(WareModel::getAll());
    }
}
