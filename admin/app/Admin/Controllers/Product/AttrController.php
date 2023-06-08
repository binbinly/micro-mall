<?php

namespace App\Admin\Controllers\Product;

use App\Admin\Common\Constant;
use App\Admin\Common\Search;
use App\Models\Product\AttrGroupModel;
use App\Models\Product\AttrModel;
use App\Models\Product\CategoryModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;

/**
 * 属性管理
 * Class AttrController
 * @package App\Admin\Controllers\Product
 */
class AttrController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '属性管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new AttrModel());
        $grid->disableExport();

        $grid->selector(function (Grid\Tools\Selector $selector) {
            $selector->select('type', '属性类型', AttrModel::$typeLabel);
        });

        $grid->column('id', 'ID');
        $this->gridCat($grid);
        $grid->column('name', '属性名');
        $grid->column('icon', 'ICON');
        $grid->column('values', '属性值')->display(function ($val) {
            return explode(',', $val);
        })->label('info');
        $grid->column('type', '类型')->using(AttrModel::$typeLabel)->label();
        if (isset($this->params['_selector']['type']) && $this->params['_selector']['type'] == AttrModel::TYPE_BASE) {//销售属性不需要分组
            $grid->column('relGroup.group_id', '所属分组')->using(AttrGroupModel::getAll())->label('primary');
            $grid->column('is_many', '是否多个值')->bool();
        }
        $grid->column('is_search', '是否可搜索')->bool();
        $grid->column('is_show', '是否显示')->bool();
        $this->releaseGrid($grid);
        $this->setGridTimeView($grid);

        $this->search($grid);

        return $grid;
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new AttrModel());

        $this->formCat($form);
        $form->text('name', '属性名')->required();
        $form->radio('type', '类型')->options(AttrModel::$typeLabel)->required()->required()
            ->when(1, function (Form $form) {
                $form->select('relGroup.group_id', '所属分组')->options(AttrGroupModel::getAll());
                $form->switch('is_many', '值类型')->states(Constant::SWITCH)->help('是否多个值');
            });
        $form->tags('values', '属性值')->required();
        $form->text('icon', 'ICON');
        $form->switch('is_search', '是否可搜索')->states(Constant::SWITCH);
        $form->switch('is_show', '是否显示')->states(Constant::SWITCH);
        $this->releaseForm($form);

        return $form;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('cat_id', '分类')->select(CategoryModel::selectOptions());
        $filter->equal('name', '属性名');
        $filter->where(function ($query) {
            $query->whereRaw("find_in_set('{$this->input}', `values`)");
        }, '属性值');
    }
}
