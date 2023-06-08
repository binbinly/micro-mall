<?php

namespace App\Admin\Controllers\SecKill;

use App\Admin\Controllers\BaseController;
use App\Models\Market\SeckillActivityModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class ActivityController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '活动管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SeckillActivityModel());
        $grid->disableExport();
        $grid->model()->orderBy('id', 'desc');

        $grid->column('id', 'ID');
        $grid->column('title', '标题');
        $grid->column('cover', '封面')->image();
        $grid->column('start_at', '开始时间');
        $grid->column('end_at', '结束时间');
        $this->releaseGrid($grid);
        $this->setGridByView($grid);
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
        $form = new Form(new SeckillActivityModel());

        $form->text('title', '标题')->required();
        $form->image('cover', '封面');
        $form->datetime('start_at', '开始时间')->required();
        $form->datetime('end_at', '结束时间')->required();
        $this->releaseForm($form);
        $this->updateByForm($form);

        return $form;
    }
}
