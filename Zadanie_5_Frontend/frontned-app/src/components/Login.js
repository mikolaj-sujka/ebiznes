import React, { useState } from 'react';
import axios from 'axios';
import { useAuth } from '../AuthContext';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { login } = useAuth(); // Use login from AuthContext

  const handleLogin = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post('http://localhost:8080/login', {
        email, password
      }, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      if (response.data.token) {
        login(response.data.token); // Update the authentication state with token
        localStorage.setItem('token', response.data.token); // Store the token in localStorage
        alert('Login successful!');
      } else {
        alert('Login failed: ' + response.data.error);
      }
    } catch (error) {
      alert('Login failed: ' + error.response.data.error);
    }
  };

  return (
    <div>
      <h2>Login</h2>
      <form onSubmit={handleLogin}>
        <div>
          <label>Email:</label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
    </div>
  );
};

export default Login;
