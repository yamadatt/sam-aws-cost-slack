
## このリポジトリは？

以下のリポジトリを参考にSAMで動作するように変更した。

https://github.com/tetsuya28/aws-cost-report

## 何ができるの？

以下のように毎朝9:00に当月累計のAWS利用料をslackに通知する。

![](./docs/Slack-example.jpg)


## ローカルでテストしたい場合

環境変数に`SLACK_TOKEN`と`SLACK_CHANNEL`を設定する。必要に応じて`LANGUAGE`も設定する（デフォルト: ja）。

```bash
export SLACK_TOKEN=xoxb-...
export SLACK_CHANNEL=information
# 英語で通知したい場合
export LANGUAGE=en
```

また、goとして動かしたいときはMakefileを使用するのが便利。

```bash
make run
```

lambdaとして試験する場合はビルドしてからテストする。

build

```bash
sam build
```

テスト。

```bash
sam local invoke
```

## AWSにデプロイする前に

パラメータストアにslackのAPIキーを設定すること。

![](./docs/parameter.jpg)

以下のようにパラメータストアから取得している。

```yaml
      Environment:
        Variables:
            SLACK_TOKEN: '{{resolve:ssm:/SLACK_API_KEY}}'
            SLACK_CHANNEL: information
```


