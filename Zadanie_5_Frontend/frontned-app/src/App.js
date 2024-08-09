import React, { useState } from "react";
import Payments from "../src/components/Payment";
import Cart from "../src/components/Cart";
import Products from "../src/components/Product";
import GoogleLogin from "./components/GoogleLogin";
import Register from "./components/Register";
import Login from "./components/Login";
import { AuthProvider, useAuth } from "./AuthContext";

const AppContent = () => {
  const [cartItems, setCartItems] = useState([]);
  const { isAuthenticated, googleLogin, login, logout } = useAuth();

  const addToCart = (product) => {
    setCartItems([...cartItems, product]);
  };

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter((item) => item.id !== id));
  };

  return (
    <div>
      {isAuthenticated ? (
        <div>
          <button onClick={logout}>Logout</button>
          <h1>Welcome to Your Dashboard</h1>
          <Products addToCart={addToCart} />
          <Cart cartItems={cartItems} removeFromCart={removeFromCart} />
          <Payments cartItems={cartItems} />
        </div>
      ) : (
        <div>
          <h2>Login, Google login or Register</h2>
          <GoogleLogin onClick={googleLogin} />
          <Login onClick={login} />
          <Register />
        </div>
      )}
    </div>
  );
};

const App = () => {
  return (
    <AuthProvider>
      <AppContent />
    </AuthProvider>
  );
};

export default App;