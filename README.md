# ipinfo

CIDR 表記からネットワーク情報を計算して表示する
シンプルな Go 製 CLI ツールです。

IPアドレス、サブネットマスク、ネットワークアドレス、
ブロードキャストアドレス、ホスト範囲、ホスト数を
表形式で分かりやすく出力します。

---

## 特徴

* CIDR からネットワーク情報を自動計算
* 複数CIDRを一度に処理可能
* 表形式で整形表示（tabwriter使用）
* シンプル・依存なし

---

## ビルド

```bash
make build
```

または

```bash
go build -o bin/ipinfo ./cmd/ipinfo
```

---

## 使い方

```bash
bin/ipinfo <CIDR> [<CIDR> ...]
```

---

## 使用例

### 単一CIDR

```bash
bin/ipinfo 10.10.10.10/24
```

出力:

```
IPAddress    SubnetMask     NetworkAddress  BroadcastAddress  HostRange                  Hosts
10.10.10.10  255.255.255.0  10.10.10.0      10.10.10.255      10.10.10.1 - 10.10.10.254  254
```

---

### 複数CIDR

```bash
bin/ipinfo 192.168.1.10/24 10.0.0.1/16
```

---

## 出力項目

| 項目               | 説明           |
| ---------------- | ------------ |
| IPAddress        | 入力IPアドレス     |
| SubnetMask       | サブネットマスク     |
| NetworkAddress   | ネットワークアドレス   |
| BroadcastAddress | ブロードキャストアドレス |
| HostRange        | 利用可能ホスト範囲    |
| Hosts            | 利用可能ホスト数     |

---

## 引数なしの場合

```bash
bin/ipinfo
```

```
Usage: bin/ipinfo <CIDR> [<CIDR> ...]
```

---

## Make コマンド

| コマンド         | 説明                         |
| ------------ | -------------------------- |
| `make tidy`  | go mod tidy                |
| `make build` | ビルド                        |
| `make run`   | サンプル実行（192.168.100.100/24） |
| `make test`  | テスト実行                      |
| `make lint`  | golangci-lint              |

---

## ディレクトリ構成

```
ipinfo/
 ├── cmd/ipinfo/main.go
 ├── internal/netcalc/
 ├── Makefile
 ├── go.mod
 └── README.md
```
