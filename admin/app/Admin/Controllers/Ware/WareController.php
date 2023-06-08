<?php

namespace App\Admin\Controllers\Ware;

use App\Admin\Controllers\BaseController;
use App\Models\Ware\WareModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class WareController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '仓库管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new WareModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('name', '仓库名');
        $grid->column('address', '仓库地址');
        $grid->column('area_code', '仓库地区编码');
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
        $form = new Form(new WareModel());

        $form->text('name', '仓库名');
        $form->text('address', '地址');
        $form->text('area_code', '地区编号');

        return $form;
    }
}
