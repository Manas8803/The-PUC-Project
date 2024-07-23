package main

import (
	"fmt"
	"os"

	event "github.com/Manas8803/puc-detection/deploy-scripts/events"
	"github.com/Manas8803/puc-detection/deploy-scripts/roles"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	dynamodb "github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

var stack_name = "PUC-Detection"

type PucDetectionStackProps struct {
	awscdk.StackProps
}

func NewPucDetectionStack(scope constructs.Construct, id string, props *PucDetectionStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	//^ Vehicle-TABLE
	vehicle_table := dynamodb.NewTable(stack, jsii.String(fmt.Sprintf("%s-Vehicle-Table", stack_name)), &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("reg_no"),
			Type: dynamodb.AttributeType_STRING,
		},
		TableName: jsii.String(fmt.Sprintf("%s-Vehicle_Table", stack_name)),
	})

	//^ User-TABLE
	user_table := dynamodb.NewTable(stack, jsii.String(fmt.Sprintf("%s-User-Table", stack_name)), &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("email"),
			Type: dynamodb.AttributeType_STRING,
		},
		TableName: jsii.String(fmt.Sprintf("%s-User_Table", stack_name)),
	})

	//^ RTO-Office-TABLE
	rto_office_table := dynamodb.NewTable(stack, jsii.String(fmt.Sprintf("%s-Rto-Office-Table", stack_name)), &dynamodb.TableProps{
		PartitionKey: &dynamodb.Attribute{
			Name: jsii.String("office_name"),
			Type: dynamodb.AttributeType_STRING,
		},
		TableName: jsii.String(fmt.Sprintf("%s-RTO-Office_Table", stack_name)),
	})

	//^ Log group of vrc handler
	logGroup_vrc := awslogs.NewLogGroup(stack, jsii.String("VRC-LogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s-VRC", stack_name)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	//^ VRC handler
	vrc_handler := awslambda.NewFunction(stack, jsii.String("VRC-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../vrc-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Role:    roles.CreateVRCHandlerRole(stack, logGroup_vrc, vehicle_table),
		Environment: &map[string]*string{
			"REGION":               jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"API_KEY":              jsii.String(os.Getenv("API_KEY")),
			"VEHICLE_TABLE_ARN":    jsii.String(*vehicle_table.TableArn()),
			"REPORT_WEBSOCKET_URL": jsii.String(os.Getenv("REPORT_WEBSOCKET_URL")),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-VRC-Lambda", stack_name)),
		LogGroup:     logGroup_vrc,
	})

	//^ Log group of fetch-vehicle handler
	logGroup_fetch_vehicle := awslogs.NewLogGroup(stack, jsii.String("Fetch-Vehicle-LogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s-Fetch-Vehicle", stack_name)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	//^ Fetch-Vehicle handler
	fetch_vehicle_handler := awslambda.NewFunction(stack, jsii.String("Fetch-Vehicle-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../fetch_vehicle-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Role:    roles.CreateFetchVehicleHandlerRole(stack, logGroup_fetch_vehicle, vehicle_table),
		Environment: &map[string]*string{
			"REGION":            jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"VEHICLE_TABLE_ARN": jsii.String(*vehicle_table.TableArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-Fetch-Vehicle-Lambda", stack_name)),
		LogGroup:     logGroup_fetch_vehicle,
	})

	//^ Fetch API Gateway
	awsapigateway.NewLambdaRestApi(stack, jsii.String("Puc_Detection_Fetch_Vehicle"), &awsapigateway.LambdaRestApiProps{
		Handler: fetch_vehicle_handler,
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
		},
	})

	//^ Log group of reg_ex_cron_job handler
	logGroup_reg_ex := awslogs.NewLogGroup(stack, jsii.String("RegExpCronJob-LogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s-RegExpCronJob", stack_name)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	//^ Registration expiration cron job handler
	reg_exp_cron_job := awslambda.NewFunction(stack, jsii.String("RegExpCronJob-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../reg_expiration_job-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Role:    roles.CreateRegExpCronJobRole(stack, logGroup_reg_ex, vrc_handler, vehicle_table),
		Environment: &map[string]*string{
			"REGION":            jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"VEHICLE_TABLE_ARN": jsii.String(*vehicle_table.TableArn()),
			"VRC_HANDLER_ARN":   jsii.String(*vrc_handler.FunctionArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-RegExp-Lambda", stack_name)),
		LogGroup:     logGroup_reg_ex,
	})

	event.CreateDailyScheduler(stack, reg_exp_cron_job)

	//^ Log group of reg_renewal_reminder handler
	logGroup_reg := awslogs.NewLogGroup(stack, jsii.String("Reg_Renewal_Reminder-LogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s-Reg_Renewal_Reminder", stack_name)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	//^ Reg_Renewal_Reminder handler
	reg_renewal_reminder_handler := awslambda.NewFunction(stack, jsii.String("Reg_Renewal_Reminder-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../reg_renewal_reminder-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Role:    roles.CreateRegReminderHandlerRole(stack, logGroup_reg, vehicle_table, vrc_handler),
		Environment: &map[string]*string{
			"REGION":            jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"VRC_HANDLER_ARN":   jsii.String(*vrc_handler.FunctionArn()),
			"VEHICLE_TABLE_ARN": jsii.String(*vehicle_table.TableArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-Reg_Renewal_Reminder-Lambda", stack_name)),
		LogGroup:     logGroup_reg,
	})

	//^ Log group of ocr handler
	logGroup_ocr := awslogs.NewLogGroup(stack, jsii.String("OCR-Lambda-LogGroup"), &awslogs.LogGroupProps{
		LogGroupName:  jsii.String(fmt.Sprintf("/aws/lambda/%s-OCR", stack_name)),
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	//^ Ocr handler
	awslambda.NewFunction(stack, jsii.String("OCR-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../ocr-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Role:    roles.CreateOCRHandlerRole(stack, logGroup_ocr, vrc_handler, reg_renewal_reminder_handler),
		Environment: &map[string]*string{
			"REGION":          jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"REG_RENEWAL_ARN": jsii.String(*reg_renewal_reminder_handler.FunctionArn()),
		},
		LogGroup:     logGroup_ocr,
		FunctionName: jsii.String(fmt.Sprintf("%s-OCR-Lambda", stack_name)),
	})

	//^ Auth handler
	auth_handler := awslambda.NewFunction(stack, jsii.String("auth-service"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../auth-service"), nil),
		Runtime: awslambda.Runtime_PROVIDED_AL2023(),
		Handler: jsii.String("main"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"JWT_SECRET_KEY": jsii.String(os.Getenv("JWT_SECRET_KEY")),
			"JWT_LIFETIME":   jsii.String(os.Getenv("JWT_LIFETIME")),
			"EMAIL":          jsii.String(os.Getenv("EMAIL")),
			"PASSWORD":       jsii.String(os.Getenv("PASSWORD")),
			"RELEASE_MODE":   jsii.String(os.Getenv("RELEASE_MODE")),
			"ADMIN":          jsii.String(os.Getenv("ADMIN")),
			"USER_TABLE_ARN": jsii.String(*user_table.TableArn()),
		},
		Role:         roles.CreateDbRole(stack, user_table),
		FunctionName: jsii.String(fmt.Sprintf("%s-Auth-Lambda", stack_name)),
	})

	//^ Auth API Gateway
	awsapigateway.NewLambdaRestApi(stack, jsii.String("Puc_Detection_Auth"), &awsapigateway.LambdaRestApiProps{
		Handler: auth_handler,
		DefaultCorsPreflightOptions: &awsapigateway.CorsOptions{
			AllowOrigins: awsapigateway.Cors_ALL_ORIGINS(),
			AllowMethods: awsapigateway.Cors_ALL_METHODS(),
			AllowHeaders: awsapigateway.Cors_DEFAULT_HEADERS(),
		},
	})

	//~ WEBSOCKET API :
	//^ Connect Route
	awslambda.NewFunction(stack, jsii.String("Connect-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../websocket/connect"), nil),
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("index.handler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"REGION":               jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"RTO_OFFICE_TABLE_ARN": jsii.String(*rto_office_table.TableArn()),
			"USER_TABLE_ARN":       jsii.String(*user_table.TableArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-Connect-Lambda", stack_name)),
		Role:         roles.CreateWebSocketLambdaRole(stack, "Connect"),
	})

	//^ Disconnect Route
	awslambda.NewFunction(stack, jsii.String("Disconnect-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../websocket/disconnect"), nil),
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("index.handler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"REGION":               jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"RTO_OFFICE_TABLE_ARN": jsii.String(*rto_office_table.TableArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-Disconnect-Lambda", stack_name)),
		Role:         roles.CreateWebSocketLambdaRole(stack, "Disconnect"),
	})

	//^ Report Authority Route
	awslambda.NewFunction(stack, jsii.String("Report-Authority-Lambda"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("../websocket/report-authority"), nil),
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("index.handler"),
		Timeout: awscdk.Duration_Seconds(jsii.Number(10)),
		Environment: &map[string]*string{
			"REGION":               jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
			"RTO_OFFICE_TABLE_ARN": jsii.String(*rto_office_table.TableArn()),
		},
		FunctionName: jsii.String(fmt.Sprintf("%s-Report-Authority-Lambda", stack_name)),
		Role:         roles.CreateWebSocketLambdaRole(stack, "Report-Authority"),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewPucDetectionStack(app, "PucDetectionStack", &PucDetectionStackProps{
		awscdk.StackProps{
			StackName: jsii.String("PucDetectionStack"),
			Env:       env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {

	// err := godotenv.Load("../.env")
	// if err != nil {
	// 	log.Fatalln("Error loading .env file : ", err)
	// }

	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}

}
