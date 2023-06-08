<?php

namespace App\Admin\Controllers\Ware;

use App\Admin\Actions\Post\Assign;
use App\Admin\Actions\Post\Draw;
use App\Admin\Actions\Post\Finish;
use App\Admin\Controllers\BaseController;
use App\Models\Ware\PurchaseModel;
use App\Models\Ware\WareModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class PurchaseController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '采购单';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new PurchaseModel());
        $grid->model()->orderBy('id', 'desc');
        $grid->disableExport();

        $grid->actions(function (Grid\Displayers\Actions $actions) {
            $actions->disableView();
            in_array($actions->row->status, [PurchaseModel::STATUS_INIT]) && $actions->add(new Assign());
            in_array($actions->row->status, [PurchaseModel::STATUS_USE]) && $actions->add(new Draw());
            in_array($actions->row->status, [PurchaseModel::STATUS_RECEIVED]) && $actions->add(new Finish());
        });

        $grid->column('id', 'ID');
        $grid->column('assignee_id', '采购人id');
        $grid->column('assignee_name', '采购人');
        $grid->column('phone', '手机号');
        $grid->column('priority', '优先级')->sortable();
        $grid->column('status', '状态')->using(PurchaseModel::$statusLabel)->filter(PurchaseModel::$statusLabel)->label();
        $grid->column('ware_id', '仓库')->using(WareModel::getAll());
        $grid->column('amount', '采购金额')->amount();
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
        $form = new Form(new PurchaseModel());

        $form->number('priority', '优先级');
        $form->select('ware_id', '仓库')->options(WareModel::getAll());

        return $form;
    }
}
