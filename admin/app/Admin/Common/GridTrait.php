<?php


namespace App\Admin\Common;


use Encore\Admin\Grid;

/**
 * 模型表格
 * Trait GridTrait
 * @package AdminBase\Traits
 */
trait GridTrait
{
    /**
     * 发布状态
     * @param Grid $grid
     */
    protected function releaseGrid(Grid $grid)
    {
        $grid->column('is_release', '是否发布')->switch(Constant::SWITCH);
    }

    /**
     * 排序
     * @param Grid $grid
     */
    protected function sortGrid(Grid $grid)
    {
        $grid->column('sort', '排序')->editable()->sortable();
    }

    /**
     * 禁用表格所有操作
     * @param Grid $grid
     */
    protected function disableGridAll(Grid $grid)
    {
        $grid->disableActions();
        $grid->disablePagination();
        $grid->disableCreateButton();
        $grid->disableFilter();
        $grid->disableRowSelector();
        $grid->disableColumnSelector();
        $grid->disableTools();
        $grid->disableExport();
    }

    /**
     * 禁用基础功能操作
     * @param Grid $grid
     */
    protected function disableGridAction(Grid $grid)
    {
        $grid->disableExport();
        $grid->disableFilter();
        $this->disableGridCreate($grid);
    }

    /**
     * 禁用基础操作 导出
     * @param Grid $grid
     */
    protected function disableGridExport(Grid $grid)
    {
        $grid->disableExport();
        $this->disableGridCreate($grid);
    }

    /**
     * 禁用基础操作 筛选
     * @param Grid $grid
     */
    protected function disableGridFilter(Grid $grid)
    {
        $grid->disableFilter();
        $this->disableGridCreate($grid);
    }

    /**
     * 禁用基础操作 创建
     * @param Grid $grid
     */
    protected function disableGridCreate(Grid $grid)
    {
        $grid->disableActions();
        $grid->disableRowSelector();
        $grid->disableCreateButton();
    }

    /**
     * 禁用筛选 创建 导出
     * @param Grid $grid
     */
    protected function disableGridCreateFilterExP(Grid $grid)
    {
        $grid->disableExport();
        $grid->disableCreateButton();
        $grid->disableFilter();
    }

    /**
     * 禁用 删除 和 显示 操作功能
     * @param Grid $grid
     */
    protected function disableGridDeleteAndView(Grid $grid)
    {
        $grid->actions(function (Grid\Displayers\Actions $actions) {
            $actions->disableDelete();
            $actions->disableView();
        });
    }

    /**
     * 只禁编辑删除操作
     * @param Grid $grid
     */
    protected function disableGridDeleteAndEdit(Grid $grid)
    {
        $grid->actions(function (Grid\Displayers\Actions $actions) {
            $actions->disableEdit();
            $actions->disableDelete();
        });
    }

    /**
     * 只禁删除操作
     * @param Grid $grid
     */
    protected function disableGridDelete(Grid $grid)
    {
        $grid->actions(function (Grid\Displayers\Actions $actions) {
            $actions->disableDelete();
        });
    }

    /**
     * 显示开始/修改时间
     * @param Grid $grid
     * @param $onlyCreated
     */
    protected function setGridTimeView(Grid $grid, $onlyCreated = true)
    {
        $grid->column('created_at', trans('admin.created_at'))->sortable();
        if ($onlyCreated) {
            $grid->column('updated_at', trans('admin.updated_at'))->sortable();
        }
    }

    /**
     * 创建者 / 修改者
     * @param Grid $grid
     */
    protected function setGridByView(Grid $grid) {
        $grid->column('create_by', '创建者');
        $grid->column('update_by', '修改者');
    }
}