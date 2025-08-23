# ex - Extended Error Package

`ex`パッケージは、`error`パッケージを拡張したライブラリ

## 機能

- **エラーの種類（Kind）**: エラーを分類して適切に処理
- **スタックトレース**: エラーの発生箇所を追跡
- **追加引数**: エラーに関連する追加情報を保持
- **エラーのラッピング**: 既存のエラーを拡張
- **パニック回復**: パニックをエラーに変換

## エラーの種類

```go
const (
    InvalidArgumentError    = "InvalidArgument"
    UnauthenticatedError    = "Unauthenticated"
    AccessDeniedError       = "AccessDenied"
    NotFoundError           = "NotFound"
    TimeoutError            = "Timeout"
    ConflictError           = "Conflict"
    ResourceExhaustedError  = "ResourceExhausted"
    CanceledError           = "Canceled"
    InternalError           = "Internal"
    DependentServiceError   = "DependentService"
    FailedPreconditionError = "FailedPrecondition"
)
```

## 基本的な使用方法

### エラーの作成

```go
// 新しいエラーを作成
err := ex.NewNotFound("user %s not found", "123")

// 既存のエラーをラップ
originalErr := errors.New("database error")
err := ex.WrapAsNotFound(originalErr, "user_id", "123")
```

### エラーの種類を確認

```go
if ex.IsNotFound(err) {
    // NotFoundエラーの処理
}

// または
if ex.GetKind(err) == ex.NotFoundError {
    // NotFoundエラーの処理
}
```

### エラーの情報を取得

```go
// 追加引数を取得
args := ex.GetArgs(err)

// スタックトレースを取得
stack := ex.GetStack(err)
```

### パニックの回復

```go
defer func() {
    if err := ex.Recover(recover()); err != nil {
        log.Printf("Recovered from panic: %v", err)
    }
}()
```

## 高度な使用方法

### エラーに追加情報を付与

```go
err := ex.NewNotFound("user not found")
err = ex.WithArgs(err, "user_id", "123", "timestamp", time.Now())
```

### カスタムエラーの作成

```go
// 新しいエラーの種類を定義
const CustomError ex.ErrorKind = "Custom"

// カスタムエラーを作成
err := ex.New("custom error message", CustomError, "additional", "info")
```

### エラーの連鎖

```go
dbErr := errors.New("database connection failed")
wrappedErr := ex.WrapAsDependentService(dbErr, "database", "postgres")
finalErr := ex.WrapAsInternal(wrappedErr, "operation", "user_create")
```