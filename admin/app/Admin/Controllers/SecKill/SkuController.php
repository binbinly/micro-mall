<?php

namespace App\Admin\Controllers\SecKill;

use App\Admin\Controllers\BaseController;
use App\Admin\Selectable\Sku;
use App\Models\Market\SeckillActivityModel;
use App\Models\Market\SeckillSessionModel;
use App\Models\Market\SeckillSkuModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class SkuController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '商品管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SeckillSkuModel());
        $grid->disableExport();
        $grid->model()->orderBy('id', 'desc');

        $grid->column('id', 'ID');
        $grid->column('sku_id', 'sku_id');
        $grid->column('activity_id', '活动')->using(SeckillActivityModel::getAll());
        $grid->column('session_id', '场次')->using(SeckillSessionModel::getAll());
        $grid->column('price', '秒杀价');
        $grid->column('count', '秒杀数量');
        $grid->column('limit', '限购数量');
        $this->sortGrid($grid);

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new SeckillSkuModel());

        $form->radio('activity_id', '选择活动')->options(SeckillActivityModel::getAll())->required();
        $form->radio('session_id', '选择场次')->options(SeckillSessionModel::getAll())->required();
        $form->belongsTo('sku_id', Sku::class, '选择商品')->required();
        $form->number('price', '秒杀价')->required();
        $form->number('count', '秒杀数量')->required();
        $form->number('limit', '个人限购数量')->default(1);
        $this->sortForm($form);

        return $form;
    }
}
