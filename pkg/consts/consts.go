// Copyright 2022 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package consts

import "time"

const (
	VideoTableName       = "video"
	UserTableName        = "user"
	SecretKey            = "secret key"
	IdentityKey          = "id"
	Total                = "total"
	Notes                = "notes"
	ApiServiceName       = "api"
	VideoServiceName     = "video"
	InteractServiceName  = "interact"
	UserServiceName      = "user"
	RedisAddr            = "localhost:6379"
	RedisPsw             = "123456"
	RedisDB              = 0
	TokenExpireFormat    = time.RFC3339
	TokenExpireDuration  = time.Duration(24 * time.Hour)
	MySQLDefaultDSN      = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP                  = "tcp"
	UserServiceAddr      = ":7777"
	InteractServiceAddr  = ":7778"
	VideoServiceAddr     = ":7779"
	MinioEndpoint        = "localhost:9000"
	MinioAccessURL       = "192.168.0.127:9000"
	MinioLocalAccessURL  = "localhost:9000"
	MinioAccessKeyId     = "tiktokMinio"
	MinioSecretAccessKey = "tiktokMinio"
	MinioUseSSL          = false
	MinioVideoBucketName = "tiktok-video"
	MinioPolicy          = "{\"Version\":\"2012-10-17\"," +
		"\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":" +
		"{\"AWS\":[\"*\"]},\"Action\":[\"s3:ListBucket\",\"s3:ListBucketMultipartUploads\"," +
		"\"s3:GetBucketLocation\"],\"Resource\":[\"arn:aws:s3:::" + MinioVideoBucketName +
		"\"]},{\"Effect\":\"Allow\",\"Principal\":{\"AWS\":[\"*\"]},\"Action\":[\"s3:PutObject\",\"s3:AbortMultipartUpload\",\"s3:DeleteObject\",\"s3:GetObject\",\"s3:ListMultipartUploadParts\"],\"Resource\":[\"arn:aws:s3:::" +
		MinioVideoBucketName +
		"/*\"]}]}"
	ExportEndpoint = ":4317"
	ETCDAddress    = "127.0.0.1:2379"
	DefaultLimit   = 10
)
