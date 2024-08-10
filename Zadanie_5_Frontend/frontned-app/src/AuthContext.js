import React, { createContext, useContext, useState, useEffect } from 'react';

const AuthContext = createContext();

export const useAuth = () => useContext(AuthContext);

export const AuthProvider = ({ children }) => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const login = (token) => {
    setIsAuthenticated(true);
    localStorage.setItem('token', token);
  };

  const googleLogin = (token) => {
    setIsAuthenticated(true);
    localStorage.setItem('token', token);
  };

  const logout = () => {
    setIsAuthenticated(false);
    localStorage.removeItem('token');
  };

  return (
    <AuthContext.Provider value={{ isAuthenticated, googleLogin, logout, login }}>
      {children}
    </AuthContext.Provider>
  );
};

export default AuthProvider;