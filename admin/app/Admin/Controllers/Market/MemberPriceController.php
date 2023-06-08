<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Common\Constant;
use App\Admin\Controllers\BaseController;
use App\Models\Market\MemberPriceModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class MemberPriceController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '会员价格优惠';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new MemberPriceModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('sku_id', 'sku_id');
        $grid->column('level_id', '会员等级id');
        $grid->column('level_name', '会员等级');
        $grid->column('price', '会员价格');
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
        $form = new Form(new MemberPriceModel());

        $form->number('sku_id', 'sku_id');
        $form->number('level_id', '会员等级id');
        $form->text('level_name', '会员等级名');
        $form->currency('price', '会员价格');
        $form->switch('is_super', '是否叠加优惠')->states(Constant::SWITCH);

        return $form;
    }
}
