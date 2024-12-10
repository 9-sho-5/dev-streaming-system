# dev-streaming-system

## 環境構築

### 1. ソースコードのダウンロード

```
git clone https://github.com/9-sho-5/dev-streaming-system.git
```

### 2. root ディレクトリに移動

```
cd dev-streaming-system
```

### 3. Docker イメージのビルドとコンテナの起動

以下のコマンドで、Docker コンテナをビルドし、バックグラウンドで起動します。

```
docker-compose up -d --build
```

### 4. フロントエンドのアクセス URL で表示を確認

```
http://localhost:8080
```

### 5. コンテナの終了

コンテナを停止・削除するためには、以下のコマンドを使用します。

```
docker-compose down
```
