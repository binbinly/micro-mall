package resource

import (
	"encoding/json"

	"github.com/binbinly/pkg/util"
	"go-micro.dev/v4/logger"

	"market/config"
	"market/model"
	cpb "pkg/proto/common"
)

// AppSettingResource 转换页面配置数据
func AppSettingResource(list []*model.AppSettingModel) (res []*cpb.AppSetting) {
	if len(list) == 0 {
		return []*cpb.AppSetting{}
	}

	for _, setting := range list {
		switch setting.Type {
		case model.AppTypeSwiper, model.AppTypeThreeAdv:
			var images []string
			if err := json.Unmarshal([]byte(setting.Data), &images); err != nil {
				logger.Warnf("[idl.app] images json unmarshal err: %v", err)
				continue
			}
			res = append(res, &cpb.AppSetting{
				Type: int32(setting.Type),
				Data: &cpb.AppSetting_Images{Images: &cpb.SettingImages{List: buildImages(images)}},
			})
		case model.AppTypeNav:
			var navs []*cpb.SettingNav
			if err := json.Unmarshal([]byte(setting.Data), &navs); err != nil {
				logger.Warnf("[idl.app] navs json unmarshal err: %v", err)
				continue
			}
			for _, nav := range navs {
				nav.Icon = util.FormatResUrl(config.Cfg.DFS.Endpoint, nav.Icon)
			}
			res = append(res, &cpb.AppSetting{
				Type: int32(setting.Type),
				Data: &cpb.AppSetting_Navs{Navs: &cpb.SettingNavs{List: navs}},
			})
		case model.AppTypeOneAdv:
			ads := &cpb.SettingAds{}
			if err := json.Unmarshal([]byte(setting.Data), ads); err != nil {
				logger.Warnf("[idl.app] ads json unmarshal err: %v", err)
				continue
			}
			ads.Cover = util.FormatResUrl(config.Cfg.DFS.Endpoint, ads.Cover)
			res = append(res, &cpb.AppSetting{
				Type: int32(setting.Type),
				Data: &cpb.AppSetting_Ads{Ads: ads},
			})
		case model.AppTypeProduct:
			product := &cpb.SettingProduct{}
			if err := json.Unmarshal([]byte(setting.Data), product); err != nil {
				logger.Warnf("[idl.app] product json unmarshal err: %v", err)
				continue
			}
			res = append(res, &cpb.AppSetting{
				Type: int32(setting.Type),
				Data: &cpb.AppSetting_Product{Product: product},
			})
		}
	}
	return
}

// AppNoticeResource 转换公告数据
func AppNoticeResource(list []*model.AppNoticeModel) (res []*cpb.Notice) {
	if len(list) == 0 {
		return []*cpb.Notice{}
	}

	for _, notice := range list {
		res = append(res, &cpb.Notice{
			Id:        int32(notice.ID),
			Title:     notice.Title,
			Content:   notice.Content,
			CreatedAt: int32(notice.CreatedAt),
		})
	}
	return
}

// PayResource 转换支付配置数据
func PayResource(list []*model.ConfigPayList) (res []*cpb.Pay) {
	if len(list) == 0 {
		return []*cpb.Pay{}
	}

	for _, p := range list {
		res = append(res, &cpb.Pay{
			Id:      int32(p.ID),
			Name:    p.Name,
			Address: p.Address,
		})
	}
	return
}

// buildImages 构建图片组路径
func buildImages(images []string) (res []string) {
	for _, image := range images {
		res = append(res, util.FormatResUrl(config.Cfg.DFS.Endpoint, image))
	}
	return
}
