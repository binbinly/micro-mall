<?php

/**
 * Laravel-admin - admin builder based on Laravel.
 * @author z-song <https://github.com/z-song>
 *
 * Bootstraper for Admin.
 *
 * Here you can remove builtin form field:
 * Encore\Admin\Form::forget(['map', 'editor']);
 *
 * Or extend custom form field:
 * Encore\Admin\Form::extend('php', PHPEditor::class);
 *
 * Or require js and css assets:
 * Admin::css('/packages/prettydocs/css/styles.css');
 * Admin::js('/packages/prettydocs/js/main.js');
 *
 */

use Encore\Admin\Form;
use Encore\Admin\Grid;
use Encore\Admin\Show;

Encore\Admin\Form::forget(['map', 'editor']);

//初始化Form表单功能
Form::init(function (Form $form) {
    $form->tools(function (Form\Tools $tools) {
        $tools->disableDelete();
        $tools->disableView();
    });
    $form->footer(function (Form\Footer $footer) {
        $footer->disableViewCheck();
    });
});

//初始化grid表格功能
Grid::init(function (Grid $grid) {
    $grid->actions(function (Grid\Displayers\Actions $actions) {
//        $actions->disableDelete();
        $actions->disableView();
    });
    $grid->batchActions(function (Grid\Tools\BatchActions $batch) {
        $batch->disableDelete();
    });
});

Show::init(function (Show $show) {
    $show->panel()->tools(function (Show\Tools $tools) {
        $tools->disableDelete();
        $tools->disableEdit();
        $tools->disableList();
    });
});

/**
 * 金币转化为元,将数值缩小100倍
 */
Grid\Column::extend('amount', function ($value) {
    return \App\Admin\Common\Format::amountToYuan($value);
});

Form::extend('sku', \App\Admin\Extensions\Sku::class);
Form::extend('photo', \App\Admin\Extensions\FormMediaField::class);