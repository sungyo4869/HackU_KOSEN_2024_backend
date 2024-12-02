# API設計

## /healthz
- **メソッド :** `GET`
- **役割 :** サーバーが生きているか確認する
- **パラメータ :** なし
- **レスポンス :**
    ```json
    {
        "message": "ok"
    }
    ```

## /login
- **メソッド :** `GET`
- **役割 :** パラメータで与えられたusernameとpasswordをDBと照会し、tokenを返す。
- **パラメータ :**
| **Name** | **Description** |
| --- | --- |
| username | ユーザーの名前 |
| password | パスワード |
- **レスポンス :**
    ```json
    {
    "token": JWTtoken
    }
    ```

## /cards
- **メソッド :** `GET`
- **役割 :** ヘッダーのトークンからuser-idを取得し、そのユーザーが所持しているカード一覧を返す。
- **パラメータ :** なし
- **レスポンス :**
    ```json
    {
    "cards": [
        {"Id":2,"UserId":2,"Picture":"card2.jpg","Name":"カード2"},
        {"Id":5,"UserId":2,"Picture":"card5.jpg","Name":"カード5"}
    ]
    }
    ```

- **メソッド :** `POST`
- **役割 :** カードをDBに新規登録する
- リクエスト
    ```json
    {
        "user-id":2,
        "picture":"/card1.png",
        "name":"おとうさん"
    }
    ```
- レスポンス
    ```json
    {
        "card-id": 1
        "user-id":2,
        "picture":"/card1.png",
        "name":"おとうさん"
    }
    ```

## /ws/matching

- **役割**: マッチングのためのWebSocket通信を確立し、マッチング完了時に情報を返す。
- **サーバー側から送信するjsonの形 :**

    ```json
    {
    "room-id": 7,
    "players": [
        {
        "user-id": 5,
        "cards": [
            {
            "card-id": 6,
            "name": "おかあさん",
            "attribute": "fire",
            "picture":"/card1.jpg"
            },
            {
            "card-id": 4,
            "name": "おとうさん",
            "attribute": "water",
            "picture":"/card2.jpg"
            },
            // ほかのカードも同様
        ]
        },
        {
        "user-id": 5,
        "cards": [
            {
            "card-id": 6,
            "name": "おかあさん",
            "attribute": "fire",
            "picture":"/card1.jpg"
            },
            {
            "card-id": 4,
            "name": "おとうさん",
            "attribute": "water",
            "picture":"/card2.jpg"
            },
            // ほかのカードも同様
        ]
        }
    ]
    }
    ```

## /ws/game

```
対戦画面に遷移した後、画面表示前に確立させる。
```

- **役割 :** ゲーム中のWebSocket通信を確立し、選択したカードやターンごとのリザルトをやり取りする。
- **クライアントが投げるjson :**

    ```json
    {
        "room-id":2,
        "user-id":2,
        "select-hand-card-id":1
    }
    ```

- **サーバーが投げるjson :**

    ```json
    {
        "battle":[
            {
                "user-id":2,
                "hp":2,
                "select-hand-card-id":1,
                "result":"loss",
                "red-card-id":null,
                "blue-card-id":6,
        "green-card-id":6,
        "kamekame-card-id":6,
        "nankuru-card-id":6
            },
            {
                "user-id":3,
                "hp":3,
                "select-hand-card-id":2,
                "result":"win",
                "red-card-id":null,
                "blue-card-id":6,
        "green-card-id":6,
        "kamekame-card-id":6,
        "nankuru-card-id":6
            }
        ]
    }
    ```
