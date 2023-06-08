<?php

namespace App\Admin\Controllers;

use App\Admin\Common\FormTrait;
use App\Admin\Common\GridTrait;
use Encore\Admin\Controllers\AdminController;
use Encore\Admin\Facades\Admin;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;
use Illuminate\Http\RedirectResponse;
use Illuminate\Support\MessageBag;

/**
 * 定义全局的一些方法，比如：开关样式、全局提示语，全局功能禁用、重写
 * Class BaseController
 * @package App\Admin\Controllers
 */
class BaseController extends AdminController
{
    use GridTrait;
    use FormTrait;

    /**
     * 列表自定义参数
     * @var array
     */
    protected $params;

    /**
     * 修改时的主键ID
     * @var int
     */
    protected $id;

    //编辑权限slug
    protected $editSlug = '';

    //新增权限slug
    protected $addSlug = '';

    /**
     * 列表
     * @param Content $content
     * @return Content
     */
    public function index(Content $content)
    {
        $this->params = request()->all();
        /** @var Grid $grid */
        $grid = $this->grid();
        if($grid->option('show_exporter')){//导出功能开启，则权限判断是否有权限
            if(!Admin::user()->can($grid->model()->getTable() . '_export')){//导出权限为数据库表 + '_export'
                $grid->disableExport();
            }
        }
        return $content
            ->title($this->title())
            ->description($this->description['index'] ?? trans('admin.list'))
            ->body($grid);
    }

    /**
     * 编辑
     * @param mixed $id
     * @param Content $content
     * @return Content
     */
    public function edit($id, Content $content)
    {
        $this->id = $id;
        return parent::edit($id, $content);
    }

    /**
     * 检查新增权限
     * @param $grid
     */
    protected function checkAddPermission(Grid $grid)
    {
        if (!Admin::user()->can($this->addSlug)) {
            $grid->disableCreateButton();
        }
    }

    /**
     * 检查编辑权限
     * @param $grid
     */
    protected function checkEditPermission(Grid $grid)
    {
        if (Admin::user()->can($this->editSlug)) {
            $this->disableGridDeleteAndView($grid);
        } else {
            $grid->disableActions();
        }
    }

    /**
     * larave-admin 全局失败提示
     * @param string $msg
     * @return RedirectResponse
     */
    protected function alertError($msg = '操作失败')
    {
        $error = new MessageBag([
            'title' => $msg,
            'message' => $msg,
        ]);
        return back()->with(compact('error'));
    }

    /**
     * 全局成功提示
     * @param string $msg
     * @return RedirectResponse
     */
    protected function alertSuccess($msg = '操作成功')
    {
        admin_success($msg);
        return back();
    }

    /**
     * 显示一个按钮
     * @param $url
     * @param $title
     * @return string
     */
    protected function showButton($url, $title)
    {
        return '<a href="' . $url . '" class="btn btn-primary"  target="_blank" rel="noopener noreferrer" style="margin:0 20px">' . $title . '</a>';
    }
}