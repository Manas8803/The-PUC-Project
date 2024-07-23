package event

import (
	"github.com/Manas8803/puc-detection/deploy-scripts/roles"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/jsii-runtime-go"
)

func CreateDailyScheduler(stack awscdk.Stack, invoke_func awslambda.Function) {

	role := roles.CreateScheduler_InvokeRole(stack, invoke_func)

	awscdk.NewCfnResource(stack, jsii.String("PUC-Detection-CRON-JOB"), &awscdk.CfnResourceProps{
		Type: jsii.String("AWS::Scheduler::Schedule"),
		Properties: &map[string]any{
			"Description": jsii.String("Cron Job scheduler for PUC Detection"),
			"FlexibleTimeWindow": &map[string]any{
				"Mode": jsii.String("OFF"),
			},
			"ScheduleExpression": jsii.String("cron(0 0 */2 3 ? 2025)"),
			"Target": &map[string]any{
				"Arn":     invoke_func.FunctionArn(),
				"RoleArn": role.RoleArn(),
			},
		},
	})
}
