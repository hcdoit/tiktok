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
	NoteTableName       = "note"
	UserTableName       = "user"
	SecretKey           = "secret key"
	IdentityKey         = "id"
	Total               = "total"
	Notes               = "notes"
	ApiServiceName      = "api"
	VideoServiceName    = "video"
	InteractServiceName = "interact"
	UserServiceName     = "user"
	RedisAddr           = "localhost:6379"
	RedisPsw            = "123456"
	RedosDB             = 0
	TokenExpireFormat   = time.RFC3339
	TokenExpireDuration = time.Duration(24 * time.Hour)
	MySQLDefaultDSN     = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP                 = "tcp"
	UserServiceAddr     = ":9000"
	InteractServiceAddr = ":9500"
	VideoServiceAddr    = ":10000"
	ExportEndpoint      = ":4317"
	ETCDAddress         = "127.0.0.1:2379"
	DefaultLimit        = 10
)
