package main

import (
	"cdk-go-sample/internal/cloud/lambda"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkGoSampleStackProps struct {
	awscdk.StackProps
}

func NewCdkGoSampleStack(scope constructs.Construct, id string, props *CdkGoSampleStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//HTTP API GATEWAY
	httpApi := awsapigatewayv2.NewHttpApi(stack, jsii.String("HttpApi"), &awsapigatewayv2.HttpApiProps{
		ApiName:     jsii.String("HelloWorldApi"),
		Description: jsii.String("HTTP API with Hello and World routes"),
	})

	//ROUTES
	lambda.RegisterHelloRoutes(stack, httpApi)
	lambda.RegisterWorldRoutes(stack, httpApi)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkGoSampleStack(app, "CdkGoSampleStack", &CdkGoSampleStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Region: jsii.String("sa-east-1"),
	}
}
