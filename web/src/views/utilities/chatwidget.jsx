import React, { useEffect, useState } from 'react';
import { Widget, addResponseMessage } from 'react-chat-widget';
import 'react-chat-widget/lib/styles.css';

function generateAutomaticResponse(message) {
  const lowercaseMessage = message.toLowerCase();

  const responses = [
    { pattern: /hello|hi|hey/g, response: "Hi there! How can I assist you?" },
    { pattern: /help|assistance/g, response: "Send an e-mail to **burakcanheyal@gmail.com** to get further assistance." },
    { pattern: /thank\s*you|thanks/g, response: "You're welcome! Feel free to ask if you need further assistance." },
    { pattern: /bye|goodbye/g, response: "Goodbye! Have a great day!" }
  ];

  for (const { pattern, response } of responses) {
    if (lowercaseMessage.match(pattern)) {
      return response;
    }
  }

  return "I'm sorry, I'm not sure I understand. Could you please rephrase your question?";
}

function ChatWidget({ messageHandler }) {
  useEffect(() => {
    addResponseMessage('Welcome to our **support** chat! How can I assist you today?');
  }, []);

  const handleNewUserMessage = (newMessage) => {
    console.log(`New message incoming! ${newMessage}`);

    const automaticResponse = generateAutomaticResponse(newMessage);

    addResponseMessage(automaticResponse);

    messageHandler(newMessage);
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
