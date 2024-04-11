package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"log"
	"strconv"
	"time"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var greeting string
	sourceIP := request.RequestContext.Identity.SourceIP

	log.Println("コスト取得バッチ 開始")
	log.Println("セッション作成")

	svc := costexplorer.New(session.Must(session.NewSession()))

	log.Println("コスト取得 実行")

	cost := GetCost(svc)

	log.Println("コスト取得バッチ 完了")

	fmt.Println(cost)

	amout := SumCost(cost)

	fmt.Println(amout)

	if sourceIP == "" {
		greeting = fmt.Sprintf("NO IP Hello, %s!\n", cost)
	} else {
		greeting = fmt.Sprintf("YES IP Hello, %s!\n", amout)
	}

	return events.APIGatewayProxyResponse{
		Body:       greeting,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}

func SumCost(cost *costexplorer.GetCostAndUsageOutput) (total string) {

	sum := 0.0
	for _, data := range cost.ResultsByTime[0].Groups {
		amount, _ := strconv.ParseFloat(*data.Metrics["UnblendedCost"].Amount, 64)
		sum = sum + amount
	}
	total = fmt.Sprintf("%.10f", sum)
	return total
}

func GetCost(svc costexploreriface.CostExplorerAPI) (result *costexplorer.GetCostAndUsageOutput) {

	// Granularity
	granularity := aws.String("DAILY")

	// Metrics
	metric := "UnblendedCost"
	metrics := []*string{&metric}

	// TimePeriod
	// 現在時刻の取得
	jst, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().UTC().In(jst)
	dayBefore := now.AddDate(0, 0, -1)

	nowDate := now.Format("2006-01-02")
	dateBefore := dayBefore.Format("2006-01-02")

	// 昨日から今日まで
	timePeriod := costexplorer.DateInterval{
		Start: aws.String(dateBefore),
		End:   aws.String(nowDate),
	}

	// GroupBy
	group := costexplorer.GroupDefinition{
		Key:  aws.String("SERVICE"),
		Type: aws.String("DIMENSION"),
	}
	groups := []*costexplorer.GroupDefinition{&group}

	// Inputの作成
	input := costexplorer.GetCostAndUsageInput{}
	input.Granularity = granularity
	input.Metrics = metrics
	input.TimePeriod = &timePeriod
	input.GroupBy = groups

	// 処理実行
	result, err := svc.GetCostAndUsage(&input)
	if err != nil {
		log.Println(err.Error())
	}

	return result
}
