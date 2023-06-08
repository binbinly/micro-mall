<?php

namespace App\Admin\Controllers\SecKill;

use App\Admin\Controllers\BaseController;
use App\Models\Market\SeckillSessionModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class SessionController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '场次管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SeckillSessionModel());
        $grid->disableExport();
        $grid->model()->orderBy('id', 'desc');

        $grid->column('id', 'ID');
        $grid->column('name', '场次名');
        $grid->column('start_at', '开始时间');
        $grid->column('end_at', '结束时间');
        $this->releaseGrid($grid);
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
        $form = new Form(new SeckillSessionModel());

        $form->text('name', '场次名')->required();
        $form->datetime('start_at', '开始时间')->required();
        $form->datetime('end_at', '结束时间')->required();
        $this->releaseForm($form);

        return $form;
    }
}
