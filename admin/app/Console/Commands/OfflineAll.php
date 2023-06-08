<?php

namespace App\Console\Commands;

use App\Admin\Services\SpuService;
use App\Models\Product\SpuModel;
use Illuminate\Console\Command;

class OfflineAll extends Command
{
    /**
     * The name and signature of the console command.
     *
     * @var string
     */
    protected $signature = 'offline:all';

    /**
     * The console command description.
     *
     * @var string
     */
    protected $description = '下架所有商品';

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
        $ids = SpuModel::query()->where('status', SpuModel::STATUS_ONLINE)->pluck('id');
        foreach ($ids as $id) {
            SpuService::offline($id);
            SpuModel::query()->where('id', $id)->update(['status' => SpuModel::STATUS_OFFLINE]);
        }
    }
}
