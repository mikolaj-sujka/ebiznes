import React, { useState } from 'react';
import Payments from '../src/components/Payment';
import Cart from '../src/components/Cart';
import Products from '../src/components/Product';


const App = () => {
  const [cartItems, setCartItems] = useState([]);

  const addToCart = (product) => {
    setCartItems([...cartItems, product]);
  };

  const removeFromCart = (id) => {
    setCartItems(cartItems.filter(item => item.id !== id));
  };

  return (
    <div>
      <Products addToCart={addToCart} />
      <Cart cartItems={cartItems} removeFromCart={removeFromCart} />
      <Payments cartItems={cartItems} />
    </div>
  );
};

export default App;
