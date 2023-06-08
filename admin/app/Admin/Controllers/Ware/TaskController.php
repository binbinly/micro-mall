<?php

namespace App\Admin\Controllers\Ware;

use App\Admin\Controllers\BaseController;
use App\Models\Ware\TaskModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class TaskController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '工作单';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new TaskModel());

        $grid->column('id', __('Id'));
        $grid->column('order_id', __('Order id'));
        $grid->column('order_no', __('Order no'));
        $grid->column('consignee', __('Consignee'));
        $grid->column('phone', __('Phone'));
        $grid->column('address', __('Address'));
        $grid->column('note', __('Note'));
        $grid->column('payment_type', __('Payment type'));
        $grid->column('status', __('Status'));
        $grid->column('desc', __('Desc'));
        $grid->column('tracking_no', __('Tracking no'));
        $grid->column('ware_id', __('Ware id'));
        $grid->column('remark', __('Remark'));
        $grid->column('created_at', __('Created at'));
        $grid->column('updated_at', __('Updated at'));

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new TaskModel());

        $form->number('order_id', __('Order id'));
        $form->text('order_no', __('Order no'));
        $form->text('consignee', __('Consignee'));
        $form->mobile('phone', __('Phone'));
        $form->text('address', __('Address'));
        $form->text('note', __('Note'));
        $form->switch('payment_type', __('Payment type'));
        $form->switch('status', __('Status'));
        $form->text('desc', __('Desc'));
        $form->text('tracking_no', __('Tracking no'));
        $form->number('ware_id', __('Ware id'));
        $form->text('remark', __('Remark'));

        return $form;
    }
}
