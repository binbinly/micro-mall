<?php


namespace App\Admin\Bean;


//{
//    "mappings": {
//    "properties": {
//        "skuId":{
//            "type":"long"
//      },
//      "spuId":{
//            "type": "keyword"
//      },
//      "skuTitle":{
//            "type": "text",
//        "analyzer": "ik_smart"
//      },
//      "skuPrice":{
//            "type": "keyword"
//      },
//      "skuImg":{
//            "type":"keyword",
//        "index": false,
//        "doc_values": false
//      },
//      "saleCount":{
//            "type": "long"
//      },
//      "hasStock":{
//            "type":"boolean"
//      },
//      "hotScore":{
//            "type":"long"
//      },
//      "brandId":{
//            "type":"long"
//      },
//      "catId":{
//            "type":"long"
//      },
//      "catName":{
//            "type":"keyword",
//        "index": false,
//        "doc_values": false
//      },
//      "brandName":{
//            "type":"keyword",
//        "index": false,
//        "doc_values": false
//      },
//      "brandImg":{
//            "type":"keyword",
//        "index": false,
//        "doc_values": false
//      },
//      "attrs":{
//            "type":"nested",
//        "properties": {
//                "attrId":{
//                    "type":"long"
//          },
//          "attrName":{
//                    "type":"keyword",
//            "index": false,
//            "doc_values": false
//          },
//          "attrValue":{
//                    "type":"keyword"
//          }
//        }
//      }
//    }
//  }
//}
class ProductBean
{
    protected int $skuId = 0;

    protected int $spuId = 0;

    protected string $skuTitle = '';

    protected int $skuPrice = 0;

    protected string $skuImg = '';

    protected int $saleCount = 0;

    protected bool $hasStock = false;

    protected int $hotScore = 0;

    protected int $brandId = 0;

    protected int $catId = 0;

    protected string $catName = '';

    protected string $brandName = '';

    protected string $brandImg = '';
}
