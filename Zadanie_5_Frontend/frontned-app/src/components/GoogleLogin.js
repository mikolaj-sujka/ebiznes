import React from 'react';

const GoogleLogin = () => {
    const handleLoginClick = () => {
      window.location.href = 'http://localhost:8080/login/google';
    };
  
    return (
      <div>
        <button onClick={handleLoginClick}>Login with Google</button>
      </div>
    );
  };

export default GoogleLogin;
