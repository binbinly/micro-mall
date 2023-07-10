<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Controllers\BaseController;
use App\Admin\Forms\PageSet;
use App\Models\Market\AppSettingModel;
use App\Models\Market\ConfigModel;
use Encore\Admin\Form;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;
use Illuminate\Support\Facades\Redis;

class AppSettingController extends BaseController
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '页面设置';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new AppSettingModel());
        $grid->disableExport();
        $grid->disableFilter();

        $grid->column('id', 'ID');
        $grid->column('page', '页面')->using(AppSettingModel::$pageLabel)->label();
        $grid->column('type', '类型')->using(AppSettingModel::$typeLabel)->label();
        $grid->column('cat_id', '分类')->using(ConfigModel::cat())->default('全部')->label('info');
        $grid->column('data', '数据');
        $grid->column('sort', '排序值')->editable()->sortable();
        $this->setGridByView($grid);
        $this->setGridTimeView($grid);

        return $grid;
    }

    /**
     * Edit interface.
     *
     * @param mixed $id
     * @param Content $content
     *
     * @return Content
     */
    public function edit($id, Content $content): Content
    {
        session()->put(PageSet::SESSION_ID, $id);
        return $content
            ->title($this->title())
            ->description($this->description['edit'] ?? trans('admin.edit'))
            ->row($this->form()->edit($id));
    }

    /**
     * Make a form builder.
     * @return Form
     */
    protected function form()
    {
        $id = session()->get(PageSet::SESSION_ID);
        $form = new Form(new AppSettingModel());

        $cats = ConfigModel::cat();
        $cats[0] = '全部';
        $form->select('page', '页面')->options(AppSettingModel::$pageLabel)->required();
        $form->select('cat_id', '分类')->default(0)->options($cats)->required();

        if ($form->isEditing()) {
            $form->radio('type', '类型')->options(AppSettingModel::$typeLabel)->disable();
            $info = AppSettingModel::query()->findOrFail($id);
            switch ($info['type']) {
                case AppSettingModel::TYPE_SWIPER:
                case AppSettingModel::TYPE_THREE_ADV:
                    $form->multipleImage('data', '图片上传')->removable()->uniqueName()->sortable();
                    break;
                case AppSettingModel::TYPE_ONE_ADV:
                    $form->embeds('data', '数据', function (Form\EmbeddedForm $form) {
                        $form->text('title', '标题')->required();
                        $form->image('cover', '封面图')->required()->uniqueName();
                    })->required();
                    break;
                case AppSettingModel::TYPE_NAV:
                    $form->table('data', '数据', function (Form\NestedForm $form) use($info) {
                        $form->text('title', '标题')->required();
                        $form->image('icon', 'ICON')->required()->uniqueName();
                    })->required();
                    break;
                case AppSettingModel::TYPE_GOODS_LIST:
                    $form->embeds('data', '数据', function (Form\EmbeddedForm $form) {
                        $form->text('router', '路由')->required();
                    })->required();
                    break;
            }
            $form->saved(function (Form $form){
                //删除缓存
                Redis::del('cache:app_page:'.$form->page.'_'.$form->cat_id);
            });
        } else {
            $form->radio('type', '类型')->options(AppSettingModel::$typeLabel);
        }
        $form->number('sort', '排序')->default(50);
        $this->updateByForm($form);

        return $form;
    }
}
