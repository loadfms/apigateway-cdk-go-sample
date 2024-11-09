package cloud

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/jsii-runtime-go"
)

func GetRouteWorld(stack awscdk.Stack) *awsapigatewayv2.AddRoutesOptions {

	//LAMBDAS
	worldLambda := awslambda.NewFunction(stack, jsii.String("WorldLambda"), &awslambda.FunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2023(),
		Handler:      jsii.String("bootstrap"),
		Code:         awslambda.Code_FromAsset(jsii.String("bin/world/world.zip"), nil),
		Architecture: awslambda.Architecture_ARM_64(),
	})

	//INTEGRATIONS
	worldIntegration := awsapigatewayv2integrations.NewHttpLambdaIntegration(jsii.String("WorldIntegration"), worldLambda, &awsapigatewayv2integrations.HttpLambdaIntegrationProps{})

	return &awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/world"),
		Methods:     &[]awsapigatewayv2.HttpMethod{awsapigatewayv2.HttpMethod_POST},
		Integration: worldIntegration,
	}
}
