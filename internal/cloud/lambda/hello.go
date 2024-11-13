package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func RegisterHelloRoutes(stack awscdk.Stack, httpApi awsapigatewayv2.HttpApi) {
	//LAMBDAS
	helloLambda := awslambda.NewFunction(stack, jsii.String("HelloLambda"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/hello/hello.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	helloIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("HelloIntegration"), helloLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	httpApi.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/hello"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_GET},
		Integration: helloIntegration,
	})
}
