module github.com/mollgx/access-demo

go 1.15

replace github.com/mmmorris1975/ssm-session-client v0.200.0 => github.com/mollgx/ssm-session-client v0.200.1-0.20210513210517-b8157aca6315

require (
	github.com/aws/aws-sdk-go-v2 v1.4.0
	github.com/aws/aws-sdk-go-v2/config v1.1.7
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.6.0
	github.com/aws/aws-sdk-go-v2/service/ssm v1.5.0
	github.com/mmmorris1975/ssm-session-client v0.200.0
)
