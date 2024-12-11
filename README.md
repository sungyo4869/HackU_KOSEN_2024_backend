# HackU_KOSEN_2024_backend

## あとでなおしたいな
- `/cards`に存在しないユーザーIDのパラメータをつけると、
    ``` json
    {
        "cards": null
    }
    ```
    が返ってきてしまう、あとから404かえるようにしたい
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
- **投げてもらうJSONの形:**
  ```json
  {"user-id": 1}
  ```
- **サーバーから送信するjsonの形 :**
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

### /ws/game
- **役割:** webSocketを確立し、ターンごとの結果を返す
- **投げてもらうJSONの形:**
  ```json
  {
      "room-id": 1,
      "user-id": 1,
      "attribute": "red",
      "card-id": 4
  }

  ```
- **サーバーから送るJSONの形:**
  ```json
  {
      "Results": [
          {
              "user-id": 2,
              "hp": 3,
              "select-attribute": "green",
              "turn-result": "win",
              "red-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "blue-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "green-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "kamekame-card-id": {
                  "Int64": 16,
                  "Valid": true
              },
              "nankuru-card-id": {
                  "Int64": 18,
                  "Valid": true
              },
              "random-card-id": {
                  "Int64": 19,
                  "Valid": true
              }
          },
          {
              "user-id": 1,
              "hp": 0,
              "select-attribute": "green",
              "turn-result": "lose",
              "red-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "blue-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "green-card-id": {
                  "Int64": 0,
                  "Valid": false
              },
              "kamekame-card-id": {
                  "Int64": 7,
                  "Valid": true
              },
              "nankuru-card-id": {
                  "Int64": 9,
                  "Valid": true
              },
              "random-card-id": {
                  "Int64": 10,
                  "Valid": true
              }
          }
      ]
  }
  
  ```
