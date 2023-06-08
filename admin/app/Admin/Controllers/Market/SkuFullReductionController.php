<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Common\Constant;
use App\Admin\Controllers\BaseController;
use App\Models\Market\SkuFullReductionModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class SkuFullReductionController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '商品满减优惠';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SkuFullReductionModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('sku_id', 'sku_id');
        $grid->column('full_price', '满多少')->sortable();
        $grid->column('reduce_price', '减多少');
        $grid->column('is_super', '是否叠加优惠')->bool();

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new SkuFullReductionModel());

        $form->number('sku_id', 'sku_id');
        $form->currency('full_price', '满多少');
        $form->currency('reduce_price', '减多少');
        $form->switch('is_super', '是否叠加优惠')->states(Constant::SWITCH);

        return $form;
    }
}
