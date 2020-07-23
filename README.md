# NSchecker
Check correct DNS NS records.

## Usage
```
go run NsCheck.go Type(NS/MX) <your domain> <NS records with comma> 
```

### Example
```
go run NsCheck.go NS "vaddy.net" "ns-1151.awsdns-15.org. , ns-1908.awsdns-46.co.uk. , ns-457.awsdns-57.com. , ns-700.awsdns-23.net." 
```

### Results
Return status code 0 if there is no problem.  
Return status code 1 or higher with error message if there there are problems.


## 何のため？
DNSレジストラへの不正アクセスなどでNSレコードが改竄され、ドメインが乗っ取られるケースがあります。  
そういったケースを早く検知するため、NSレコードが正しいかチェックするツールです。  
