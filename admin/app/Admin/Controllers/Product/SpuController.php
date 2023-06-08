<?php

namespace App\Admin\Controllers\Product;

use App\Admin\Actions\Post\Offline;
use App\Admin\Actions\Post\Online;
use App\Admin\Common\Search;
use App\Admin\Forms\Spu;
use App\Models\Product\BrandModel;
use App\Models\Product\CategoryModel;
use App\Models\Product\SpuModel;
use Encore\Admin\Grid;
use Encore\Admin\Layout\Content;
use Encore\Admin\Widgets\MultipleSteps;

/**
 * 商品管理
 * Class SpuController
 * @package App\Admin\Controllers\Product
 */
class SpuController extends BaseController
{
    use Search;

    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '商品管理';

    /**
     * Make a grid builder.
     *
     * @return Grid
     */
    protected function grid()
    {
        $grid = new Grid(new SpuModel());
        $grid->model()->orderBy('id', 'desc');

        $grid->actions(function (Grid\Displayers\Actions $actions) {
            $actions->disableView();
            in_array($actions->row->status, [SpuModel::STATUS_INIT]) && $actions->add(new Online());
            in_array($actions->row->status, [SpuModel::STATUS_ONLINE]) && $actions->add(new Offline());
        });

        $grid->column('id', 'ID');
        $this->gridCat($grid);
        $grid->column('brand_id', '品牌')->using(BrandModel::getAll())->label('info');
        $grid->column('name', '商品名称');
        $grid->column('desc', '描述');
        $grid->column('status', '状态')->using(SpuModel::$statusLabel)->label('primary');
        $grid->column('is_many', '是否多规格')->bool();
        $grid->column('original_price', '市场价')->amount();;
        $this->setGridTimeView($grid);
        $this->search($grid);

        return $grid;
    }

    public function create(Content $content)
    {
        $steps = [
            'basic' => Spu\Basic::class,
            'attr' => Spu\Attr::class,
            'sku' => Spu\Sku::class
        ];

        return $content
            ->title('添加商品')
            ->body(MultipleSteps::make($steps));
    }

    protected function filter(Grid\Filter &$filter)
    {
        $filter->equal('cat_id', '分类')->select(CategoryModel::selectOptions());
        $filter->equal('brand_id', '品牌')->select(BrandModel::getAll());
        $filter->equal('name', '商品名');
    }
}
