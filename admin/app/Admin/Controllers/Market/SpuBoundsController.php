<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Common\Constant;
use App\Admin\Controllers\BaseController;
use App\Models\Market\SpuBoundsModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class SpuBoundsController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '商品积分';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SpuBoundsModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('spu_id', 'spu_id');
        $grid->column('grow_bounds', '成长值');
        $grid->column('buy_bounds', '购物积分');
        $grid->column('work', '是否赠送')->bool();

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new SpuBoundsModel());

        $form->number('spu_id', 'spu_id');
        $form->number('grow_bounds', '成长值');
        $form->number('buy_bounds', '购物积分');
        $form->switch('work', '是否赠送')->states(Constant::SWITCH)->help('优惠生效情况[0 - 无优惠，成长积分是否赠送');

        return $form;
    }
}
