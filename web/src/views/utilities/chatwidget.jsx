import React, { useEffect } from 'react';
import { Widget, addResponseMessage } from 'react-chat-widget';

import 'react-chat-widget/lib/styles.css';

function ChatWidget() {
  useEffect(() => {
    addResponseMessage('Welcome to our **support** chat! How can I assist you today?');
  }, []);

  const handleNewUserMessage = (newMessage) => {
    console.log(`New message incoming! ${newMessage}`);
    const response = generateResponse(newMessage);
    addResponseMessage(response);
  };

  const generateResponse = (message) => {
    if (message.toLowerCase().includes('hello') || message.toLowerCase().includes('hi')) {
      return 'Hi there! How can I assist you today?';
    } else if (message.toLowerCase().includes('help')) {
      return 'Sure, I can help you. What do you need assistance with?';
    } else {
      return "I'm sorry, I didn't understand that. Could you please rephrase?";
    }
  };

  const handleQuickButtonClicked = (value) => {
    console.log(`Quick button clicked! Value: ${value}`);
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
        handleQuickButtonClicked={handleQuickButtonClicked}
        showTimeStamp={true}
        resizable={false}
        emojis={false}
        showBadge={true}
      />
    </div>
  );
}

export default ChatWidget;
