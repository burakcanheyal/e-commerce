// ChatWidget.jsx

import React, { useEffect, useState } from 'react';
import { Widget, addResponseMessage } from 'react-chat-widget';
import 'react-chat-widget/lib/styles.css';

function ChatWidget({ messageHandler }) {
  useEffect(() => {
    addResponseMessage('Welcome to our **support** chat! How can I assist you today?');
  }, []);

  const handleNewUserMessage = (newMessage) => {
    console.log(`New message incoming! ${newMessage}`);
    messageHandler(newMessage); // Yeni mesajı ana bileşene iletiyoruz
  };

  return (
    <div className="App">
      <Widget
        handleNewUserMessage={handleNewUserMessage}
        title="Welcome to our Live Helper"
        subtitle="Ask anything to get help"
        senderPlaceHolder="Type a message..."
        showCloseButton={true}
        fullScreenMode={false}
        autofocus={true}
        launcherOpenLabel="Open chat"
        launcherCloseLabel="Close chat"
        sendButtonAlt="Send"
        showTimeStamp={true}
        resizable={false}
        emojis={false}
        showBadge={true}
      />
    </div>
  );
}

export default ChatWidget;
