# HackU_KOSEN_2024_backend

## あとでなおしたいな
- `/cards`に存在しないユーザーIDのパラメータをつけると、
    ``` json
    {
        "cards": null
    }
    ```
    が返ってきてしまう、あとから404かえるようにしたい

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
