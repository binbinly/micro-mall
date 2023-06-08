<?php

namespace App\Admin\Controllers\Product;

use App\Models\Product\SkuAttrModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class SkuAttrController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = 'SkuAttrModel';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SkuAttrModel());

        $grid->column('id', __('Id'));
        $grid->column('sku_id', __('Sku id'));
        $grid->column('attr_id', __('Attr id'));
        $grid->column('attr_name', __('Attr name'));
        $grid->column('attr_value', __('Attr value'));
        $grid->column('sort', __('Sort'));

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new SkuAttrModel());

        $form->number('sku_id', __('Sku id'));
        $form->number('attr_id', __('Attr id'));
        $form->text('attr_name', __('Attr name'));
        $form->text('attr_value', __('Attr value'));
        $form->number('sort', __('Sort'))->default(50);

        return $form;
    }
}
