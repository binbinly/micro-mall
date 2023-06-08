<?php

namespace App\Console\Commands;

use App\Admin\Services\SpuService;
use App\Models\Product\SkuModel;
use App\Models\Product\SpuModel;
use App\Models\Ware\WareSkuModel;
use Illuminate\Console\Command;

class OnlineAll extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'online:all';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = '上架所有商品';

    /**
     * Create a new command instance.
     *
     * @return void
     */
    public function __construct()
    {
        parent::__construct();
    }

    /**
     * Execute the console command.
     *
     * @return mixed
     */
    public function handle()
    {
//        SpuModel::query()->where('status', SpuModel::STATUS_OFFLINE)
//            ->update(['status' => SpuModel::STATUS_INIT]);exit;
        $ids = SpuModel::query()->where('status', SpuModel::STATUS_INIT)->pluck('id');
        foreach ($ids as $id) {
            try{
                //添加库存
                SkuModel::query()->where('spu_id', $id)->chunk(100, function ($collection){
                    $skus = $collection->toArray();
                    foreach ($skus as $sku) {
                        WareSkuModel::query()->create([
                            'sku_id' => $sku['id'],
                            'ware_id' => 1,
                            'stock' => 10,
                        ]);
                    }
                });
                SpuService::online($id);
                SpuModel::query()->where('id', $id)->update(['status' => SpuModel::STATUS_ONLINE]);
            }catch (\Exception $e) {
                $this->output->error($e->getMessage());
            }
        }
        $this->output->success('finish');
    }
}
