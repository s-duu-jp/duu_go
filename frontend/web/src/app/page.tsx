"use client"
import { useState, useEffect, useRef } from 'react';
import { useSearchParams } from 'next/navigation';

interface Message {
  username: string;
  message: string;
  channel: string;
}

export default function Home() {
  const searchParams = useSearchParams();
  const [username, setUsername] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const [messages, setMessages] = useState<Message[]>([]);
  const ws = useRef<WebSocket | null>(null);
  const [isConnected, setIsConnected] = useState<boolean>(false);
  const [channelId, setChannelId] = useState<string | null>(null);

  const id = searchParams.get('id');

  useEffect(() => {
    if (!id) {
      return;
    }

    setChannelId(id); // チャネルIDを状態に保存

    const connectWebSocket = () => {
      // 既に接続されている場合は何もしない
      if (ws.current) return

      // 新しいWebSocket接続を作成する
      ws.current = new WebSocket(`ws://localhost:3000/chat?id=${id}`);


      // 接続が確立されたときにsetIsConnected(true)を呼び出して接続状態を更新する
      ws.current.onopen = () => {
        setIsConnected(true);
      };

      // 受信したメッセージを処理し、setMessagesを使ってメッセージリストを更新する
      ws.current.onmessage = (event) => {
        const msg = JSON.parse(event.data);
        console.log(event)
        setMessages((prevMessages) => [...prevMessages, msg]);
      };

      // 接続が閉じられたときにsetIsConnected(false)を呼び出して接続状態を更新します。
      // 異常終了の場合（event.code !== 1000）、再接続ロジックを使用して1秒後に再接続を試みます。
      ws.current.onclose = (event) => {
        setIsConnected(false);
        if (event.code !== 1000) {
          // 再接続ロジック
          setTimeout(() => {
            connectWebSocket();
          }, 1000);
        }
      };
    };

    connectWebSocket();

    return () => {
      if (ws.current) {
        ws.current.close();
        ws.current = null;
      }
    };
  }, [id]);

  // メッセージを送信する
  const sendMessage = () => {
    if (ws.current && message) {
      const msg = { username, message, channel: id };
      ws.current.send(JSON.stringify(msg));
      setMessage('');
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <h1>WebSocket Chat</h1>
      {/* idがない場合は非表示 */}
      {channelId && (
        <>
          <h2>Room: {channelId}</h2>
          <div>
            <input
              type="text"
              placeholder="Username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              style={{ marginRight: '10px' }}
            />
          </div>
          <div>
            <textarea
              placeholder="Type a message..."
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              style={{ marginRight: '10px', marginTop: '10px', width: '300px', height: '100px' }}
            />
          </div>
          <div>
            <button onClick={sendMessage} style={{ marginTop: '10px' }} disabled={!isConnected}>
              Send
            </button>
          </div>
          <div style={{ marginTop: '20px' }}>
            <h2>Chat</h2>
            <div style={{ border: '1px solid black', padding: '10px', height: '300px', overflowY: 'scroll' }}>
              {messages.map((msg, index) => (
                <p key={index}><strong>{msg.username}:</strong> {msg.message}</p>
              ))}
            </div>
          </div>
        </>
      )}
    </div>
  );
}
