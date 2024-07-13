"use client"
import { useState, useEffect, useRef } from 'react';

interface Message {
  username: string;
  message: string;
}

export default function Home() {
  const [username, setUsername] = useState<string>('');
  const [message, setMessage] = useState<string>('');
  const [messages, setMessages] = useState<Message[]>([]);
  const ws = useRef<WebSocket | null>(null);

  useEffect(() => {
    ws.current = new WebSocket('ws://localhost:3000/chat');

    ws.current.onmessage = (event) => {
      const msg = JSON.parse(event.data);
      setMessages((prevMessages) => [...prevMessages, msg]);
    };

    return () => {
      if (ws.current) {
        ws.current.close();
      }
    };
  }, []);

  const sendMessage = () => {
    if (ws.current && message) {
      const msg = { username, message };
      ws.current.send(JSON.stringify(msg));
      setMessage('');
    }
  };

  return (
    <div style={{ padding: '20px' }}>
      <h1>WebSocket Chat</h1>
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
        <button onClick={sendMessage} style={{ marginTop: '10px' }}>Send</button>
      </div>
      <div style={{ marginTop: '20px' }}>
        <h2>Chat</h2>
        <div style={{ border: '1px solid black', padding: '10px', height: '300px', overflowY: 'scroll' }}>
          {messages.map((msg, index) => (
            <p key={index}><strong>{msg.username}:</strong> {msg.message}</p>
          ))}
        </div>
      </div>
    </div>
  );
}
