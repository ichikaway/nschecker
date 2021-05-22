# NSchecker
DNS record changing detection tool with slack notification.  
It is easy to use as we provide an executable files for Linux, MacOS.

In some cases, NS records have been tampered with by unauthorized access to DNS registrars and the domain has been hijacked.
In order to detect such cases, this tool checks if the NS records are correct.

If you run cron regularly, it will only notify you of slack when there is a problem, so you can be aware of unintentional NS/MX changes before they happen.


## Lookup with No dns cache
The NS records are retrieved from the DNS Root servers(authority servers) and not affected by dns cache.

MX records refer to the local DNS cache server(full resolver).


## Usage
#### Download exe binary
Download linux/macOS binary from here.   
https://github.com/ichikaway/nschecker/releases/

Download latest linux binary directly.  
https://github.com/ichikaway/nschecker/releases/latest/download/nschecker-linux-64bit

```
wget https://github.com/ichikaway/nschecker/releases/latest/download/nschecker-linux-64bit
chmod 700 nschecker-linux-64bit
```

#### for Linux user
```
./nschecker-linux-64bit -type Type(NS/MX) -domain <your domain> -expect <NS records with comma> 
```

#### for Mac user
```
./nschecker-macOS-64bit -type Type(NS/MX) -domain <your domain> -expect <NS records with comma> 
```

#### or, go run command.
```
go run NSchecker.go -type Type(NS/MX) -domain <your domain> -expect <NS records with comma> 
```

### Example
```
./nschecker-linux-64bit -type NS -domain "vaddy.net"  -expect "ns-1151.awsdns-15.org. , ns-1908.awsdns-46.co.uk. , ns-457.awsdns-57.com. , ns-700.awsdns-23.net." 
```

## Results
Return status code 0 if there is no problem.  
Return status code 1 or higher with error message if there there are problems.

## Slack Notification
Set slack webhook settings on OS env, 
NSchecker sends error message to the slack channel when detecting errors or DNS record changing.

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

cronで定期実行すれば問題がある時のみslack通知するため、意図しないNS/MXの変更に早く気付けます。


## DNSキャッシュの影響は? 

TLD権威サーバに登録されているNSレコードについては、DNS Rootサーバからデータを取得するため、キャッシュの影響はうけません。

MXレコードはローカルDNSキャッシュサーバを参照します。