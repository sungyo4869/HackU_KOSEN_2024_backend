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
     | password              | ユーザーのパスワード  |
- **レスポンス:**
  ```json
  {
    "user-id": 1,
    "token": "jwt-token"
  }
    ```

### /cards
- **メソッド:** `GET`
- **役割:** ユーザーが所持しているカード一覧を返す
- **レスポンス:**
  ```json
  {
    "cards": [
      {"card-id":1,"user-id":1,"picture":"pochi1.jpg","name":"カード1"},
      {"card-id":2,"user-id":1,"picture":"pochi2.jpg","name":"カード2"},
      // ほかのカードも返す
    ]
  }
    ```
- **メソッド :** `POST`
- **役割 :** カードをDBに新規登録する
- **リクエスト**
  ```json
  {
    "user-id":2,
    "picture":"/card1.png",
    "name":"おとうさん"
  }
  ```
- **レスポンス**

  ```json
  {
    "card-id": 1,
    "user-id":2,
    "picture":"/card1.png",
    "name":"おとうさん"
  }
  ```
### /select
- **メソッド:** `GET`
- **役割:** ユーザーのuser_selectedテーブルの中身を返す
- **レスポンス :**
  ```json
  {
      "cards": [
          {
            "user-selected-id":2,
            "card-id": 6,
            "attribute": "red",
          },
          {
          "user-selected-id":3,
            "card-id": 4,
            "attribute": "blue",
          },
          // ほかのカードも返す
      ]
  }
  ```
- **メソッド:** `PUT`
- **役割:** 手札を変更し、格納する
- **リクエスト**
  ```json
  {
    "selected-cards":[
	  {
		  "attribute":"red",
	    "card-id":1,
	  },
	  {
		  "attribute":"blue",
	    "card-id":2,
	  },
	  // ほかのカードも
	]
  }
  ```
- **レスポンス :**
  ```json
  {
    "selected-cards":[
      {
        "selected-card-id":2,
        "card-id":1,
        "attribute":"red"
      },
      {
        "selected-card-id":2,
        "card-id":1,
        "attribute":"blue"
      },
      
      // ほかのカードも
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
### /ws/shogun
- **役割:** ゲーム開始時に将軍ぽちを選択し情報を返す
- **投げてもらうJSONの形:**
```json
{
  "room-id": 20,
  "user-id": 1,
  "shogun-id": 1
}
```
- **サーバから送る形:**
```json
{
  "players": [
    {
      "room-id": 33,
      "user-id": 1,
      "shogun-id": 2
    },
    {
      "room-id": 33,
      "user-id": 2,
      "shogun-id": 3
    }
  ]
}
```