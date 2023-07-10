<?php


namespace App\Admin\Controllers\Product;


use App\Models\Product\CategoryModel;
use Encore\Admin\Form;
use Encore\Admin\Tree;
use Encore\Admin\Layout\Content;

/**
 * 商品分类
 * Class CategoryController
 * @package App\Admin\Controllers\Product
 */
class CategoryController extends BaseController
{
    public function index(Content $content): Content
    {
        $tree = new Tree(new CategoryModel);

        return $content
            ->header('商品分类树')
            ->body($tree);
    }

    protected function form()
    {
        $form = new Form(new CategoryModel());

        $form->select('parent_id', '父ID')->options(CategoryModel::selectOptions(null, '所属分类'))->required();
        $form->text('name', '分类名')->required();
        $this->sortForm($form);
        $this->releaseForm($form);

        return $form;
    }
}
