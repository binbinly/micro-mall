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
    protected $skuId = 0;

    protected $spuId = 0;

    protected $skuTitle = '';

    protected $skuPrice = 0;

    protected $skuImg = '';

    protected $saleCount = 0;

    protected $hasStock = false;

    protected $hotScore = 0;

    protected $brandId = 0;

    protected $catId = 0;

    protected $catName = '';

    protected $brandName = '';

    protected $brandImg = '';
}