<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Actions\Href;
use App\Admin\Actions\Post\CouponRelease;
use App\Admin\Controllers\BaseController;
use App\Models\Market\CouponModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

class CouponController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '优惠券管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new CouponModel());
        $grid->model()->orderBy('id', 'desc');
        $grid->disableExport();

        $grid->actions(function (Grid\Displayers\DropdownActions $actions) {
            $actions->disableView();
            $actions->disableDelete();
            $actions->add(new Href('详情', '/coupon_member'));
            $actions->add(new CouponRelease());
        });

        $grid->column('id', 'ID');
        $grid->column('name', '名称');
        $grid->column('cover', '封面');
        $grid->column('type', '类型')->using(CouponModel::$typeLabel)->label();
        $grid->column('num', '数量');
        $grid->column('amount', '金额');
        $grid->column('per_limit', '每人限领张数');
        $grid->column('min_point', '使用门槛');
        $grid->column('start_at', '使用开始时间');
        $grid->column('end_at', '使用结束时间');
        $grid->column('use_type', '使用类型')->using(CouponModel::$useTypeLabel);
        $grid->column('note', '备注');
        $grid->column('publish_count', '发行数量');
        $grid->column('use_count', '使用数量');
        $grid->column('receive_count', '领取数量');
        $grid->column('enable_start_at', '可领取开始时间');
        $grid->column('enable_end_at', '可领取结束时间');
        $grid->column('code', '优惠码');
        $grid->column('member_level', '可领取的会员等级');
        $grid->column('is_release', '是否已发布')->bool();
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
        $form = new Form(new CouponModel());

        $form->text('name', '优惠券名称')->required();
        $form->image('cover', '封面图');
        $form->select('type', '类型')->options(CouponModel::$typeLabel)->required();
        $form->number('num', '数量')->required();
        $form->currency('amount', '金额')->required();
        $form->number('per_limit', '每人限领张数')->default(1);
        $form->currency('min_point', '使用门槛')->required();
        $form->datetime('start_at', '起始时间')->required();
        $form->datetime('end_at', '结束时间')->required();
        $form->select('use_type', '使用类型')->options(CouponModel::$useTypeLabel)->required();
        $form->text('note', '备注');
        $form->datetime('enable_start_at', '可以领取的开始时间')->required();
        $form->datetime('enable_end_at', '可以领取的结束时间')->required();
        $form->text('code', '优惠码');
        $form->number('member_level', '可以领取的会员等级')->default(0)->help('可以领取的会员等级[0->不限等级，其他-对应等级]');

        return $form;
    }
}
