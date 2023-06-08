<?php

namespace App\Admin\Extensions;

use Encore\Admin\Form\Field;

class Sku extends Field
{
    protected $view = 'sku';

    /**
     * 商品销售属性
     * @var array
     */
    protected $attrs = [];

    /**
     * 商品标题
     * @var string
     */
    protected $title = '';

    protected static $js = [
        'vendor/sku/sku.js'
    ];

    protected static $css = [
        'vendor/sku/sku.css'
    ];

    public function data($attrs = [])
    {
        $this->attrs = $attrs;
        return $this;
    }

    public function title($title = '') {
        $this->title = $title;
        return $this;
    }

    public function render()
    {
        $this->script = <<< EOF
window.DemoSku = new JadeKunSKU('{$this->getElementClassSelector()}')
EOF;
        $this->variables = array_merge($this->variables, [
            'attrs' => $this->attrs,
            'title' => $this->title
        ]);
        return parent::render();
    }

}
