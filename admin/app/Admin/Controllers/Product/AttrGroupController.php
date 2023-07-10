<?php

namespace App\Admin\Controllers\Product;

use App\Admin\Actions\Post\RelAttr;
use App\Admin\Common\Search;
use App\Models\Product\AttrGroupModel;
use App\Models\Product\AttrRelGroupModel;
use App\Models\Product\CategoryModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;

/**
 * 属性分组
 * Class AttrGroupController
 * @package App\Admin\Controllers\Product
 */
class AttrGroupController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '属性分组管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new AttrGroupModel());
        $grid->disableExport();

        $grid->column('id', 'ID');
        $this->gridCat($grid);
        $grid->column('name', '分组名');
        $grid->column('desc', '描述');
        $grid->column('icon', 'ICON');
        $this->sortGrid($grid);
        $this->search($grid);

        return $grid;
    }

    public function edit($id, Content $content): Content
    {
        return $content
            ->title($this->title())
            ->description($this->description['edit'] ?? trans('admin.edit'))
            ->body($this->form()->edit($id))
            ->row($this->relAttr($id));
    }

    /**
     * Make a form builder.
     *
     * @return Form
     */
    protected function form()
    {
        $form = new Form(new AttrGroupModel());

        $this->formCat($form);
        $form->text('name', '分组名')->required();
        $form->text('desc', '描述');
        $form->text('icon', 'ICON');
        $this->sortForm($form);

        return $form;
    }

    /**
     * 关联属性
     * @param $groupId
     * @return Grid
     */
    protected function relAttr($groupId){
        $grid = new Grid(new AttrRelGroupModel());
        $grid->setTitle('关联属性');
        $grid->model()->where('group_id', $groupId);
        $this->disableGridCreateFilterExP($grid);

        $grid->actions(function (Grid\Displayers\DropdownActions $action){
            $action->disableView();
            $action->disableEdit();
        });

        $grid->tools(function (Grid\Tools $tools) use($groupId) {
            $tools->append(new RelAttr($groupId));
        });

        $grid->column('attr_id', '属性id');
        $grid->column('attr.name', '属性名');

        return $grid;
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('cat_id', '分类')->select(CategoryModel::selectOptions());
        $filter->equal('name', '分组名');
    }
}
