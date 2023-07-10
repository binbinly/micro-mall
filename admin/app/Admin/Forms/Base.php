<?php


namespace App\Admin\Forms;


use Encore\Admin\Auth\Permission;
use Encore\Admin\Widgets\Form;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;

class Base extends Form
{
    protected $slug;

    public function handle(Request $request)
    {
        $this->slug && Permission::check($this->slug);
    }

    /**
     * 成功
     * @param string $msg
     * @return RedirectResponse
     */
    protected function success(string $msg = '操作成功'): RedirectResponse
    {
        admin_success($msg);
        return back();
    }

    /**
     * 失败
     * @param string $msg
     * @return RedirectResponse
     */
    protected function error(string $msg = '操作失败'): RedirectResponse
    {
        admin_error($msg);
        return back();
    }
}
