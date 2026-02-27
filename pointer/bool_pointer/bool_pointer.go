package main

import (
	"fmt"
)

type AppInfoScrapperResult struct {
	Country                  string   `json:"country"`
	PackageName              string   `thrift:"package_name,1,required" frugal:"1,required,string" json:"package_name"`
	BundleId                 string   `thrift:"bundle_id,2,required" frugal:"2,required,string" json:"bundle_id"`
	IconUrl                  string   `thrift:"icon_url,3,required" frugal:"3,required,string" json:"icon_url"`
	Website                  string   `thrift:"website,4,required" frugal:"4,required,string" json:"website"`
	Name                     string   `thrift:"name,5,required" frugal:"5,required,string" json:"name"`
	DownloadUrl              string   `thrift:"download_url,6,required" frugal:"6,required,string" json:"download_url"`
	Installs                 *string  `thrift:"installs,7,optional" frugal:"7,optional,string" json:"installs,omitempty"`
	Ratings                  int64    `thrift:"ratings,8,required" frugal:"8,required,i64" json:"ratings"`
	ContentRating            string   `thrift:"content_rating,9,required" frugal:"9,required,string" json:"content_rating"`
	ContentExplain           []string `thrift:"content_explain,10,optional" frugal:"10,optional,list<string>" json:"content_explain,omitempty"`
	Genre                    string   `thrift:"genre,11,required" frugal:"11,required,string" json:"genre"`
	Score                    string   `thrift:"score,12,required" frugal:"12,required,string" json:"score"`
	DeveloperName            *string  `thrift:"developer_name,13,optional" frugal:"13,optional,string" json:"developer_name,omitempty"`
	SellerName               *string  `thrift:"seller_name,14,optional" frugal:"14,optional,string" json:"seller_name,omitempty"`
	ScreenshotUrls           []string `thrift:"screenshot_urls,15,optional" frugal:"15,optional,list<string>" json:"screenshot_urls,omitempty"`
	IpadScreenshotUrls       []string `thrift:"ipad_screenshot_urls,16,optional" frugal:"16,optional,list<string>" json:"ipad_screenshot_urls,omitempty"`
	Version                  *string  `thrift:"version,17,optional" frugal:"17,optional,string" json:"version,omitempty"`
	VersionUpdated           *string  `thrift:"version_updated,18,optional" frugal:"18,optional,string" json:"version_updated,omitempty"`
	Size                     *string  `thrift:"size,19,optional" frugal:"19,optional,string" json:"size,omitempty"`
	AppId                    *int64   `thrift:"app_id,20,optional" frugal:"20,optional,i64" json:"app_id,omitempty"`
	ContentIntroduction      *string  `thrift:"content_introduction,21,optional" frugal:"21,optional,string" json:"content_introduction,omitempty"`
	DeveloperAllPackageNames []string `thrift:"developer_all_package_names,22,optional" frugal:"22,optional,list<string>" json:"developer_all_package_names"`
	Available                *bool    `thrift:"available,24,optional" frugal:"24,optional,bool" json:"available"`
}

type DetailAppInfo struct {
	PackageName string `thrift:"package_name,1,required" frugal:"1,required,string" json:"package_name"`
	// 包名
	BundleId string `thrift:"bundle_id,2,required" frugal:"2,required,string" json:"bundle_id"`
	// 图标url
	IconUrl string `thrift:"icon_url,3,required" frugal:"3,required,string" json:"icon_url"`
	// developerWebsite
	Website string `thrift:"website,4,required" frugal:"4,required,string" json:"website"`
	// 应用名称
	Name string `thrift:"name,5,required" frugal:"5,required,string" json:"name"`
	// url链接
	DownloadUrl string `thrift:"download_url,6,required" frugal:"6,required,string" json:"download_url"`
	// 下载次数 5,000,000+
	Installs *string `thrift:"installs,7,optional" frugal:"7,optional,string" json:"installs,omitempty"`
	// 评论条数 105506
	Ratings int64 `thrift:"ratings,8,required" frugal:"8,required,i64" json:"ratings"`
	// 内容分级 16岁以上
	ContentRating string `thrift:"content_rating,9,required" frugal:"9,required,string" json:"content_rating"`
	// 内容分级说明  [ '重度暴力' ]
	ContentExplain []string `thrift:"content_explain,10,optional" frugal:"10,optional,list<string>" json:"content_explain,omitempty"`
	// 分类 动作
	Genre string `thrift:"genre,11,required" frugal:"11,required,string" json:"genre"`
	// 评分 4.0
	Score string `thrift:"score,12,required" frugal:"12,required,string" json:"score"`
	// 开发者
	DeveloperName *string `thrift:"developer_name,13,optional" frugal:"13,optional,string" json:"developer_name,omitempty"`
	// 供应商
	SellerName *string `thrift:"seller_name,14,optional" frugal:"14,optional,string" json:"seller_name,omitempty"`
	// iphone上架截图
	ScreenshotUrls []string `thrift:"screenshot_urls,15,optional" frugal:"15,optional,list<string>" json:"screenshot_urls,omitempty"`
	// ipad上架截图
	IpadScreenshotUrls []string `thrift:"ipad_screenshot_urls,16,optional" frugal:"16,optional,list<string>" json:"ipad_screenshot_urls,omitempty"`
	// 版本号
	Version *string `thrift:"version,17,optional" frugal:"17,optional,string" json:"version,omitempty"`
	// 版本发布时间
	VersionUpdated *string `thrift:"version_updated,18,optional" frugal:"18,optional,string" json:"version_updated,omitempty"`
	// 大小 单位byte
	Size *string `thrift:"size,19,optional" frugal:"19,optional,string" json:"size,omitempty"`
	// ios trackId
	Id *int64 `thrift:"id,20,optional" frugal:"20,optional,i64" json:"id,omitempty"`
	// 内容简介
	ContentIntroduction *string `thrift:"content_introduction,22,optional" frugal:"22,optional,string" json:"content_introduction,omitempty"`
	// 该开发者维度下 其他应用包名
	DeveloperAllPackageNames []string `thrift:"developer_all_package_names,23,optional" frugal:"23,optional,list<string>" json:"developer_all_package_names,omitempty"`
	// 当前应用是否可用
	Available *bool `thrift:"available,24,optional" frugal:"24,optional,bool" json:"available,omitempty"`
	// 当前应用预览视频链接
	Video *string `thrift:"video,25,optional" frugal:"25,optional,string" json:"video,omitempty"`
}

func TestEmptyAppInfo() {
	appInfoEmpty := AppInfoScrapperResult{}
	ApplyPackageParse(appInfoEmpty)
}

func TestNilAvailableInAppInfo() {
	appInfoEmpty := AppInfoScrapperResult{Available: nil}
	ApplyPackageParse(appInfoEmpty)
}

func TestNilAvailableFromDetailAppInfo() {
	d1 := DetailAppInfo{}
	appInfo := AppInfoScrapperResult{
		Available: d1.Available,
	}
	ApplyPackageParse(appInfo)
}

func main() {
	// d1 := DetailAppInfo{}
	// appInfo := AppInfoScrapperResult{
	// 	Available: d1.Available,
	// }
	// ApplyPackageParse(appInfo)
	// b := false

	TestEmptyAppInfo()
	TestNilAvailableInAppInfo()
	TestNilAvailableFromDetailAppInfo()
}
func ApplyPackageParse(result AppInfoScrapperResult) {
	if result.Available != nil {
		if *result.Available {
			fmt.Println("Available is true")
		} else {
			println("Available is false")
		}
	} else {
		println("Available is nil")
	}

}
