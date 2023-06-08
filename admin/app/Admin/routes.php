<?php

use Illuminate\Routing\Router;

Admin::routes();

Route::group([
    'prefix' => config('admin.route.prefix'),
    'namespace' => config('admin.route.namespace'),
    'middleware' => config('admin.route.middleware'),
    'as' => config('admin.route.prefix') . '.',
], function (Router $router) {

    $router->get('/', 'HomeController@index')->name('home');

    #region 产品服务
    $router->resource('product/category', 'Product\CategoryController')->names('商品分类');
    $router->resource('product/brand', 'Product\BrandController')->names('品牌管理');
    $router->resource('product/spu', 'Product\SpuController')->names('spu管理');
    $router->resource('product/sku', 'Product\SkuController')->names('商品管理');
    $router->resource('product/attr', 'Product\AttrController')->names('属性管理');
    $router->resource('product/attr_group', 'Product\AttrGroupController')->names('属性分组管理');
    $router->resource('product/sku_attr', 'Product\SkuAttrController')->names('销售属性管理');
    #endregion

    #region 仓储服务
    $router->resource('ware/list', 'Ware\WareController')->names('仓库管理');
    $router->resource('ware/sku', 'Ware\WareSkuController')->names('商品库存');
    $router->resource('ware/task', 'Ware\TaskController')->names('库存工作单');
    $router->resource('ware/purchase', 'Ware\PurchaseController')->names('库存采购单');
    $router->resource('ware/purchase_detail', 'Ware\PurchaseDetailController')->names('采购需求');
    #endregion

    #region 营销服务
    $router->resource('market/notice', 'Market\AppNoticeController')->names('公告管理');
    $router->resource('market/setting', 'Market\AppSettingController')->names('页面设置');
    $router->resource('market/config', 'Market\ConfigController')->names('页面设置');
    $router->resource('market/coupon', 'Market\CouponController')->names('优惠券');
    $router->resource('market/coupon_member', 'Market\CouponMemberController')->names('会员优惠券');
    $router->resource('market/member_price', 'Market\MemberPriceController')->names('会员价格优惠');
    $router->resource('market/full_reduction', 'Market\SkuFullReductionController')->names('满减优惠');
    $router->resource('market/bounds', 'Market\SpuBoundsController')->names('spu积分');
    #endregion

    #region 营销服务 - 秒杀管理
    $router->resource('seckill/activity', 'SecKill\ActivityController')->names('秒杀活动');
    $router->resource('seckill/session', 'SecKill\SessionController')->names('秒杀活动');
    $router->resource('seckill/sku', 'SecKill\SkuController')->names('秒杀活动');
    $router->resource('seckill/notice', 'SecKill\SkuNoticeController')->names('秒杀活动');
    #endregion

});
