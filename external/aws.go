package external

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/account"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
)

var iconURLMap = map[string]string{
	"AWS Cost Explorer":                               "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/CloudFinancialManagement/CostExplorer.png?raw=true",
	"AWS Key Management Service":                      "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/SecurityIdentityCompliance/KeyManagementService.png?raw=true",
	"AWS Lambda":                                      "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Compute/Lambda.png?raw=true",
	"AWS X-Ray":                                       "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/DeveloperTools/XRay.png?raw=true",
	"Amazon API Gateway":                              "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ApplicationIntegration/APIGateway.png?raw=true",
	"Amazon Simple Email Service":                     "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/BusinessApplications/SimpleEmailService.png?raw=true",
	"Amazon DynamoDB":                                 "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Database/DynamoDB.png?raw=true",
	"Amazon EC2 Container Registry (ECR)":             "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Containers/ElasticContainerRegistry.png?raw=true",
	"Amazon Elastic Container Registry Public":        "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Containers/ElasticContainerRegistry.png?raw=true",
	"Amazon Elastic Container Service":                "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Containers/ElasticContainerService.png?raw=true",
	"Amazon Elastic Load Balancing":                   "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/NetworkingContentDelivery/ElasticLoadBalancingApplicationLoadBalancer.png?raw=true",
	"Amazon Relational Database Service":              "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Database/Aurora.png?raw=true",
	"Amazon Route 53":                                 "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/NetworkingContentDelivery/Route53.png?raw=true",
	"Amazon Simple Storage Service":                   "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Storage/SimpleStorageService.png?raw=true",
	"AmazonCloudWatch":                                "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ManagementGovernance/CloudWatch.png?raw=true",
	"Amazon CloudFront":                               "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/NetworkingContentDelivery/CloudFront.png?raw=true",
	"AWS Amplify":                                     "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/FrontEndWebMobile/AmplifyAWSAmplifyStudio.png?raw=true",
	"AWS Glue":                                        "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Analytics/Glue.png?raw=true",
	"Amazon Simple Notification Service":              "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ApplicationIntegration/SimpleNotificationService.png?raw=true",
	"AWS Secrets Manager":                             "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/SecurityIdentityCompliance/SecretsManager.png?raw=true",
	"Amazon Virtual Private Cloud":                    "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Groups/VPC.png?raw=true",
	"AWS WAF":                                         "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/SecurityIdentityCompliance/WAF.png?raw=true",
	"EC2 - Other":                                     "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Compute/EC2.png?raw=true",
	"AWS Step Functions":                              "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ApplicationIntegration/StepFunctions.png?raw=true",
	"Amazon Elastic Compute Cloud - Compute":          "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Compute/EC2.png?raw=true",
	"AWS CloudTrail":                                  "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ManagementGovernance/CloudTrail.png?raw=true",
	"Amazon Simple Queue Service":                     "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/ApplicationIntegration/SimpleQueueService.png?raw=true",
	"Amazon GuardDuty":                                "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/SecurityIdentityCompliance/GuardDuty.png?raw=true",
	"Amazon Elastic Container Service for Kubernetes": "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/Containers/EKSCloud.png?raw=true",
	"Amazon Cognito":                                  "https://github.com/awslabs/aws-icons-for-plantuml/blob/main/dist/SecurityIdentityCompliance/Cognito.png?raw=true",
	"Tax":                                             "",
}

func GetIconURL(service string) string {
	return iconURLMap[service]
}

func GetCost() (*costexplorer.GetCostAndUsageOutput, error) {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location()).Format("2006-01-02")
	tomorrow := now.AddDate(0, 0, 1).Format("2006-01-02")

	granularity := "MONTHLY"
	metrics := []string{
		"AmortizedCost",
		"BlendedCost",
		"UnblendedCost",
		"UsageQuantity",
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	if err != nil {
		return nil, err
	}

	svc := costexplorer.New(sess)
	result, err := svc.GetCostAndUsage(&costexplorer.GetCostAndUsageInput{
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(startOfMonth),
			End:   aws.String(tomorrow),
		},
		Granularity: aws.String(granularity),
		GroupBy: []*costexplorer.GroupDefinition{
			{
				Type: aws.String("DIMENSION"),
				Key:  aws.String("SERVICE"),
			},
		},
		Metrics: aws.StringSlice(metrics),
		Filter: &costexplorer.Expression{
			Not: &costexplorer.Expression{
				Dimensions: &costexplorer.DimensionValues{
					Key:    aws.String("RECORD_TYPE"),
					Values: aws.StringSlice([]string{"Credit"}),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetAccountFullName(ctx context.Context) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}

	client := account.NewFromConfig(cfg)

	output, err := client.GetContactInformation(ctx, &account.GetContactInformationInput{})
	if err != nil {
		return "", err
	}

	return *output.ContactInformation.FullName, nil
}
