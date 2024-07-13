# プロダクト名：SAMPLE

## **概要**

SAMPLE

---

## **前提条件**

お使いの環境が利用可能でない場合は[こちら](https://code.visualstudio.com/docs/remote/containers)を参考に事前作業を行ってください。

1. VSCode Remote Containers が利用可能である事

---

## **環境構築手順**

以下の手順に従って下さい。

1. 本リポジトリを git clone します

   ```bash
   $ git clone git@github.com:s-duu-jp/sample.git
   ```

1. クローン後、以下コマンドでコンテナを起動します。

   ```bash
   $ arg="sample" \
     && cd ${arg} \
     && code --folder-uri vscode-remote://dev-container+$(echo -n $(pwd) | xxd -p)/workspace
   ```

---

## **サービス起動手順**

| システム       | 起動コマンド      | アクセス URL          |
| -------------- | ----------------- | --------------------- |
| すべて起動     | **$ npm start**   | -                     |
| バックエンド   | **$ npm run api** | http://localhost:3333 |
| フロントエンド | **$ npm run web** | http://localhost:4200 |

---

## **作業の進め方**

本プロジェクトでは与えられたタスクもしくは作業が必要なタスク毎にブランチを作成します。

### **作業の開始**

1. 既に起票中のタスクを割り当てられた場合

   1-1. ブランチを作成する前にチケット ID(XX-X)を確認するため[プロジェクトボード](https://id.atlassian.com/)へ移動します。

   1-2. 割り当てられたタスクのチケット ID (例:DC-1)を確認します。

   1-3. ブランチを作成します。

   ```bash
   $ git checkout -b ${チケットID}
   ```

   `(例：$ git checkout -b DC-1)`

   1-4. 作成したブランチをリポジトリへ反映します。

   ```bash
   $ git push --set-upstream origin ${チケットID}
   ```

   `(例：$ git push --set-upstream origin DC-1)`

1. 新たにチケット ID を発行して作業を行う場合

   2-1. [プロジェクトボード](https://id.atlassian.com/)へ移動します。

   2-2. 適所にタスクを起票してチケット ID (例:DC-1)を発行します。

   2-3. ブランチを作成します。

   ```bash
   $ git checkout -b ${チケットID}
   ```

   `(例：$ git checkout -b DC-1)`

   2-4. 作成したブランチをリポジトリへ反映します。

   ```bash
   $ git push --set-upstream origin ${チケットID}
   ```

   `(例：$ git push --set-upstream origin DC-1)`

### **Pull Request を出す**

1. Github の[プロジェクト](https://id.atlassian.com/)へアクセスして任意のブランチから`Pull request`をクリックします。

1. 修正内容を記入します。

1. 最後に`Create pull request`をクリックします。

---

## **開発するにあたって**

### コーディング規約

- `any`は極力利用しないで下さい！

  TypeScript なので`any`は極力使用せず型を指定するようにしてください。
  型を利用することでプロパティへ安全にアクセスできる、インテリセンスが利く、どんなデータを扱いたいのかわかりやすく可読性が上がる等のメリットがあります。
  any を使う理由がある場合は以下のようにコメントで理由を記載して該当箇所だけ Lint を無効にしてください。

  ```ts
  // anyを使う理由
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  ```


# 全体
```mermaid
sequenceDiagram
    participant Client as クライアント（ブラウザ）
    participant WebSocketServer as WebSocketサーバー

    Client->>WebSocketServer: WebSocket接続要求
    WebSocketServer-->>Client: WebSocket接続確立

    Client->>Client: ユーザー名とメッセージを入力
    Client->>WebSocketServer: メッセージ送信
    WebSocketServer->>WebSocketServer: メッセージを受信してブロードキャスト
    WebSocketServer->>Client: メッセージをブロードキャスト

    Client->>Client: ブロードキャストされたメッセージを受信して表示

    Client->>WebSocketServer: WebSocket切断要求
    WebSocketServer-->>Client: WebSocket切断確立
```

# フロント

```mermaid
sequenceDiagram
    participant User as ユーザー
    participant Client as クライアント（Reactコンポーネント）
    participant WebSocketServer as WebSocketサーバー

    User->>Client: ページを開く
    Client->>Client: useEffect()
    Client->>WebSocketServer: new WebSocket('ws://localhost:3000/chat')
    WebSocketServer-->>Client: WebSocket接続確立
    Client->>Client: onmessage() { setMessages() }

    User->>Client: ユーザー名を入力 (setUsername)
    User->>Client: メッセージを入力 (setMessage)
    User->>Client: 送信ボタンをクリック (sendMessage)
    Client->>WebSocketServer: ws.current.send()

    WebSocketServer->>Client: メッセージを受信
    Client->>Client: onmessage() { setMessages() }

    User->>Client: ページを閉じる
    Client->>Client: useEffect Cleanup { ws.current.close() }
    WebSocketServer-->>Client: WebSocket切断確立
```

# バックエンド

```mermaid
sequenceDiagram
    participant Client as クライアント（ブラウザ）
    participant WebSocketServer as WebSocketサーバー
    participant HandleConnections as HandleConnections
    participant HandleMessages as HandleMessages
    participant BroadcastChannel as ブロードキャストチャンネル

    Client->>WebSocketServer: WebSocket接続要求
    WebSocketServer->>HandleConnections: HandleConnections()
    HandleConnections->>Client: WebSocket接続確立
    HandleConnections->>HandleConnections: clients[ws] = true

    loop メッセージ受信
        Client->>HandleConnections: メッセージ送信
        HandleConnections->>HandleConnections: ws.ReadJSON(&msg)
        HandleConnections->>BroadcastChannel: broadcast <- msg
    end

    loop メッセージブロードキャスト
        HandleMessages->>BroadcastChannel: msg = <-broadcast
        BroadcastChannel->>HandleMessages: メッセージ受信
        HandleMessages->>HandleMessages: for client := range clients
        HandleMessages->>Clients: client.WriteJSON(msg)
        alt エラーが発生した場合
            HandleMessages->>HandleMessages: log.Printf("error: %v", err)
            HandleMessages->>HandleMessages: client.Close()
            HandleMessages->>HandleMessages: delete(clients, client)
        end
    end

    Client->>HandleConnections: WebSocket切断
    HandleConnections->>HandleConnections: delete(clients, ws)
    HandleConnections->>Client: WebSocket切断確立
```
