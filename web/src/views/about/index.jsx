import React, { useState, useEffect } from 'react';

function WalletBalance() {
  // Bakiye durumu için state tanımla
  const [balance, setBalance] = useState(null);

  // Sayfa yüklendiğinde bakiyeyi al
  useEffect(() => {
    getBalance();
  }, []);

  // GET isteği gönderme fonksiyonu
  const getBalance = () => {
    // AccessToken değerini buraya ekleyin
    const accessToken = 'your_access_token_here';

    fetch('http://localhost:8001/wallet/a', {
      headers: {
        'Authentication': `AccessToken ${accessToken}`
      }
    })
      .then(response => response.json()) // Gelen veriyi JSON formatına dönüştür
      .then(data => {
        // Gelen bakiye değerini state'e kaydet
        setBalance(data.balance);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };

  return (
    <div>
      <h1>Wallet Balance</h1>
      {/* Bakiye varsa göster */}
      {balance !== null && (
        <p>Balance: {balance}</p>
      )}
    </div>
  );
}

export default WalletBalance;
