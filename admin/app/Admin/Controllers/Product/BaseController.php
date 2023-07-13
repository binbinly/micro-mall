<?php


namespace App\Admin\Controllers\Product;


use App\Models\Product\CategoryModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

/**
 * 产品服务基础控制器
 * Class BaseController
 * @package App\Admin\Controllers\Product
 */
class BaseController extends \App\Admin\Controllers\BaseController
{
    protected function formCat(Form $form, $column = 'cat_id'){
        $form->select($column, '分类')->options(CategoryModel::selectOptions(null, '选择分类'))->required()->rules('required|min:1');
    }

    protected function gridCat(Grid $grid, $column = 'cat_id') {
        $grid->column($column, '所属分类')->using(CategoryModel::getAll())->default('默认')->label();;
    }
}
