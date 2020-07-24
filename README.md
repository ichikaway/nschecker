# NSchecker
DNS record changing detection tool. 
To detect if DNS records have been tampered with by unauthorized access to DNS registrar.

## Usage
#### for Linux user
```
./bin/vaddy-linux-64bit Type(NS/MX) <your domain> <NS records with comma> 
```

#### for Mac user
```
./bin/vaddy-macosx-64bit Type(NS/MX) <your domain> <NS records with comma> 
```

#### or, go run command.
```
go run NsCheck.go Type(NS/MX) <your domain> <NS records with comma> 
```

### Example
```
./bin/vaddy-linux-64bit NS "vaddy.net" "ns-1151.awsdns-15.org. , ns-1908.awsdns-46.co.uk. , ns-457.awsdns-57.com. , ns-700.awsdns-23.net." 
```

## Results
Return status code 0 if there is no problem.  
Return status code 1 or higher with error message if there there are problems.

## Slack Notification
Set slack webhook settings on OS env, 
NsCheck sends error message to the slack channel when detecting errors or DNS record changing.

```cassandraql
export SLACK_WEBHOOK_URL="webhook url"
export SLACK_FREE_TEXT="<!channel> from AWS lambda" #optional
export SLACK_USERNAME="your user" #optional
export SLACK_CHANNEL="your channel" #optional
export SLACK_ICON_EMOJI=":smile:" #optional
export SLACK_ICON_URL="icon url" #optional
```

## 何のため？
DNSレジストラへの不正アクセスなどでNSレコードが改竄され、ドメインが乗っ取られるケースがあります。  
そういったケースを早く検知するため、NSレコードが正しいかチェックするツールです。  

実行して問題がなければステータスコード0を返して何も表示しません。  
問題があれば1以上のステータスコードを返してエラーメッセージや現在DNS登録されているレコード情報を出力します。  
