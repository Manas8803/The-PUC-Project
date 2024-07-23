package roles

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/jsii-runtime-go"
)

func CreateOCRHandlerRole(stack awscdk.Stack, log_group awslogs.LogGroup, vrc_handler awslambda.Function, reg_renewal_handler awslambda.Function) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("OCR-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("textract:DetectDocumentText"),
			jsii.String("textract:AnalyzeDocument"),
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			jsii.String(*vrc_handler.FunctionArn()),
			jsii.String(*reg_renewal_handler.FunctionArn()),
			jsii.String(*log_group.LogGroupArn()),
			jsii.String("*"),
		},
	}))

	return role
}

func CreateVRCHandlerRole(stack awscdk.Stack, log_group awslogs.LogGroup, db dynamodb.Table) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("VRC-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
		},
		Resources: &[]*string{
			jsii.String(*db.TableArn()),
			jsii.String(*log_group.LogGroupArn()),
			jsii.String("*"),
		},
	}))
	return role
}
func CreateFetchVehicleHandlerRole(stack awscdk.Stack, log_group awslogs.LogGroup, db dynamodb.Table) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("Fetch-Vehicle-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
		},
		Resources: &[]*string{
			jsii.String(*db.TableArn()),
			jsii.String(*log_group.LogGroupArn()),
			jsii.String("*"),
		},
	}))
	return role
}
func CreateRegReminderHandlerRole(stack awscdk.Stack, log_group awslogs.LogGroup, db dynamodb.Table, vrc_handler awslambda.Function) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("Reg_Reminder-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			jsii.String(*db.TableArn()),
			jsii.String(*vrc_handler.FunctionArn()),
			jsii.String(*log_group.LogGroupArn()),
			jsii.String("*"),
		},
	}))
	return role
}

func CreateScheduler_InvokeRole(stack awscdk.Stack, invoke_handler awslambda.Function) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("PUCDetection-Invoke-Reg_Exp_Cron_Job"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("scheduler.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
		RoleName:  jsii.String("Reg_Exp_Cron_Job"),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			invoke_handler.FunctionArn(),
			jsii.String("*"),
		},
	}))

	return role
}

func CreateRegExpCronJobRole(stack awscdk.Stack, log_group awslogs.LogGroup, invoke_handler awslambda.Function, db dynamodb.Table) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("RegExCronJobRole-Lambda-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("lambda:InvokeFunction"),
		},
		Resources: &[]*string{
			jsii.String(*invoke_handler.FunctionArn()),
			jsii.String(*db.TableArn()),
			jsii.String(*log_group.LogGroupArn()),
			jsii.String("*"),
		},
	}))
	return role
}

func CreateDbRole(stack awscdk.Stack, db dynamodb.Table) awsiam.Role {
	role := awsiam.NewRole(stack, jsii.String("DB-Auth-Role"), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
		},
		Resources: &[]*string{
			jsii.String(*db.TableArn()),
			jsii.String("*"),
		},
	}))
	return role
}

func CreateWebSocketLambdaRole(stack awscdk.Stack, role_name string) awsiam.Role {
	role_name = fmt.Sprintf("WebSocket-Lambda-Role-%s", role_name)
	role := awsiam.NewRole(stack, jsii.String(role_name), &awsiam.RoleProps{
		AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), &awsiam.ServicePrincipalOpts{}),
	})

	role.AddToPolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions: &[]*string{
			jsii.String("logs:CreateLogGroup"),
			jsii.String("logs:PutLogEvents"),
			jsii.String("logs:DescribeLogStreams"),
			jsii.String("logs:CreateLogStream"),
			jsii.String("dynamodb:BatchGet*"),
			jsii.String("dynamodb:DescribeStream"),
			jsii.String("dynamodb:DescribeTable"),
			jsii.String("dynamodb:Get*"),
			jsii.String("dynamodb:Query"),
			jsii.String("dynamodb:Scan"),
			jsii.String("dynamodb:BatchWrite*"),
			jsii.String("dynamodb:CreateTable"),
			jsii.String("dynamodb:Delete*"),
			jsii.String("dynamodb:Update*"),
			jsii.String("dynamodb:PutItem"),
			jsii.String("execute-api:Invoke"),
			jsii.String("execute-api:ManageConnections"),
		},
		Resources: &[]*string{
			jsii.String("*"),
		},
	}))
	return role
}
