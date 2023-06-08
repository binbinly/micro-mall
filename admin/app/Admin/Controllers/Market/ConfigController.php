<?php

namespace App\Admin\Controllers\Market;

use App\Admin\Forms\HomeCat;
use App\Admin\Forms\Pay;
use App\Http\Controllers\Controller;
use Encore\Admin\Layout\Content;
use Encore\Admin\Widgets\Tab;

class ConfigController extends Controller
{
    /**
     * Title for current resource.
     *
     * @var string
     */
    protected $title = '配置管理';

    protected function index(Content $content)
    {
        return $content
            ->title('配置中心')
            ->body(Tab::forms([
                'cat' => HomeCat::class,
                'pay' => Pay::class
            ]));
    }
}
