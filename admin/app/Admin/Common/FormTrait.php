<?php


namespace App\Admin\Common;


use Encore\Admin\Facades\Admin;
use Encore\Admin\Form;

/**
 * 模型表单
 * Trait FormTrait
 * @package AdminBase\Traits
 */
trait FormTrait
{
    /**
     * 发布状态
     * @param Form $form
     */
    protected function releaseForm(Form $form){
        $form->switch('is_release', '是否发布')->states(Constant::SWITCH);
    }

    /**
     * 排序
     * @param Form $form
     */
    protected function sortForm(Form $form) {
        $form->number('sort', '排序值')->default(50)->help('值越大越靠前');
    }

    /**
     * 更新修改者
     * @param Form $form
     */
    protected function updateByForm(Form $form) {
        $form->saving(function (Form $form){
            if ($form->isCreating()) {
                $form->model()->create_by = Admin::user()->id;
            } else {
                $form->model()->update_by = Admin::user()->id;
            }
        });
    }

    /**
     * 禁用form表单底部check选项
     * @param Form $form
     */
    protected function disableFormCheck(Form &$form)
    {
        $form->footer(function (Form\Footer $footer) {
            $footer->disableViewCheck();
            $footer->disableEditingCheck();
            $footer->disableCreatingCheck();
        });
    }

    /**
     * 禁用form表单底部功能操作
     * @param Form $form
     */
    protected function disableFormFooter(Form &$form)
    {
        $form->footer(function (Form\Footer $footer) {
            $footer->disableViewCheck();
            $footer->disableEditingCheck();
            $footer->disableCreatingCheck();
            $footer->disableReset();
        });
    }

    /**
     * 禁用底部所有操作
     * @param Form $form
     */
    protected function disableFormFooterAll(Form &$form)
    {
        $form->footer(function (Form\Footer $footer) {
            $footer->disableViewCheck();
            $footer->disableEditingCheck();
            $footer->disableCreatingCheck();
            $footer->disableReset();
            $footer->disableSubmit();
        });
    }
}