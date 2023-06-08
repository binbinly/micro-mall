<?php

namespace App\Admin\Controllers\Product;

use App\Admin\Common\Search;
use App\Models\Product\BrandModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

/**
 * 品牌管理
 * Class BrandController
 * @package App\Admin\Controllers\Product
 */
class BrandController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '品牌管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new BrandModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $grid->column('name', '品牌名');
        $this->gridCat($grid, 'relCat.cat_id');
        $grid->column('logo', 'LOGO')->image('', 50);
        $grid->column('cover', '封面')->image();
        $grid->column('desc', '描述');
        $this->releaseGrid($grid);
        $this->sortGrid($grid);
        $this->search($grid);
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
        $form = new Form(new BrandModel());

        $form->text('name', '品牌名')->required();
        $this->formCat($form, 'relCat.cat_id');
        $form->image('logo', '品牌logo');
        $form->image('cover', '品牌封面');
        $form->textarea('desc', '品牌描述');
        $this->sortForm($form);
        $this->releaseForm($form);

        return $form;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('name', '品牌名');
    }
}
