# HackU_KOSEN_2024_backend

## あとでなおしたいな
- `/cards`に存在しないユーザーIDのパラメータをつけると、
    ``` json
    {
        "cards": null
    }
    ```
    が返ってきてしまう、あとから404かえるようにしたい
- `/cards`のやつがたぶんrow.Scanのときの順番が間違っているので直す
- /wsのやつ認証機能つけてないので時間あったらつける。
- コードが汚いので時間があったらきれいにする

## API一覧

### /healthz
- **メソッド:** `GET`
- **役割:** サーバーが生きているか確認する
- **パラメータ:** なし
- **レスポンス:**
  ```json
  {
    "message": "ok"
  }
    ```

### /login
- **メソッド:** `GET`
- **役割:** パラメータで与えられたusernameをDBと照会し、user-idを返す
- **パラメータ:**

     | **name**              | **Description**   |
     |:---------------------:|:-------------------:|
     | username              | ユーザーの名前       |

- **レスポンス:**
  ```json
  {
    "user-id": 2
  }
    ```

### /cards
- **メソッド:** `GET`
- **役割:** パラメータで与えられたuser-idから照会し、ユーザーが所持しているカード一覧を返す
- **パラメータ:**

     | **name**              | **Description**   |
     |:---------------------:|:-------------------:|
     | user-id               | ユーザー固有のID       |

- **レスポンス:**
  ```json
  {
    "cards": [
        {"Id":2,"UserId":2,"Picture":"カード2","Name":"card2.jpg"},
        {"Id":5,"UserId":2,"Picture":"カード5","Name":"card5.jpg"}
    ]
  }
    ```
### /selected-cards
- **メソッド:** `GET`
- **役割:** パラメータで与えられたuser-idから照会し、ユーザーが所持しているバトルに使うときのカード一覧を返す
- **パラメータ:**

     | **name**              | **Description**   |
     |:---------------------:|:-------------------:|
     | user-id               | ユーザー固有のID       |

- **レスポンス:**
  ```json
  {
    "selected-cards": [
        {"Id":2,"UserId":2,"Atttribute":"red"},
        {"Id":3,"UserId":5,"Atttribute":"water"},
    ]
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
      "username": "user1",
      "selected-cards": [
        {
          "card-id": 6,
          "attribute": "red",
          "name": "おかあさん",
          "picture":"/card1.jpg"
        },
        {
          "card-id": 4,
          "attribute": "blue",
          "name": "おとうさん",
          "picture":"/card2.jpg"
        },
        // ほかのカードも同様
      ]
    },
    {
      "username": "user2",
      "selected-cards": [
        {
          "card-id": 6,
          "attribute": "fire",
          "name": "おかあさん",
          "picture":"/card1.jpg"
        },
        {
          "card-id": 4,
          "attribute": "water",
          "name": "おとうさん",
          "picture":"/card2.jpg"
        },
        // ほかのカードも同様
      ]
    }
  ]
}
```
