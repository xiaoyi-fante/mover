/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xiaoyi-fante/mover/cloud/aliyun"
)

var (
	ossBucket   string
	ossKey      string
	ossSecret   string
	ossEndpoint string
)

// ossCmd represents the oss command
var ossCmd = &cobra.Command{
	Use:   "oss",
	Short: "迁移到阿里云OSS，",
	Long: `迁移到阿里云 OSS 服务，通过 oss 命令指定将图片迁移到阿里云 OSS。比如：
			mover oss --bucket=xxx --key=xxx --secret=xxx --endpoint=xxx --path=xxxx`,
	Run: func(cmd *cobra.Command, args []string) {
		//初始化阿里云的oss client
		uploader, err := aliyun.NewAliOSS(ossBucket, ossEndpoint, ossKey, ossSecret)
		if err != nil {
			fmt.Printf("初始化OSS客户端出错了：%v\n", err)
			return
		}
		// TODO, 开始迁移了
		StartMove(uploader)
	},
}

func init() {
	rootCmd.AddCommand(ossCmd)

	ossCmd.Flags().StringVarP(&ossBucket, "bucket", "b", "", "指定阿里云 OSS Bucket")
	ossCmd.Flags().StringVarP(&ossKey, "key", "k", "", "阿里云 OSS Key")
	ossCmd.Flags().StringVarP(&ossEndpoint, "endpoint", "e", "os-cn-beijing.aliyuncs.com", "OSS Endpoint")
	ossCmd.Flags().StringVarP(&ossSecret, "secret", "s", "", "阿里云 OSS Secret")
	ossCmd.MarkFlagRequired("bucket")
	ossCmd.MarkFlagRequired("key")
	ossCmd.MarkFlagRequired("secret")

}
