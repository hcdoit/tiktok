package consts

const (
	MinioEndpoint        = "localhost:9000"
	MinioAccessURL       = "192.168.246.19:9000"
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
)
