import React, { useState } from "react";
import Payments from "../src/components/Payment";
import Cart from "../src/components/Cart";
import Products from "../src/components/Product";
import Login from "./components/Login";
import Register from "./components/Register";
import { AuthProvider, useAuth } from "./AuthContext";

const AppContent = () => {
  const [cartItems, setCartItems] = useState([]);
  const { isAuthenticated, logout } = useAuth();

  const addToCart = (product) => {
    setCartItems([...cartItems, product]);
  };

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter((item) => item.id !== id));
  };

  return (
    <div>
      {isAuthenticated ? (
        <>
          <button onClick={logout}>Logout</button>
          <Products addToCart={addToCart} />
          <Cart cartItems={cartItems} removeFromCart={removeFromCart} />
          <Payments cartItems={cartItems} />
        </>
      ) : (
        <>
          <h2>Login</h2>
          <Login />
          <Register />
        </>
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
