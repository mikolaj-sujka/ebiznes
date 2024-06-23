import React from 'react';

const Cart = ({ cartItems, removeFromCart }) => {
  const handleRemove = (id) => {
    removeFromCart(id);
  };

  return (
    <div>
      <h1>Cart</h1>
      {cartItems.length > 0 ? (
        <ul>
          {cartItems.map(item => (
            <li key={item.id}>
              {item.name} - ${item.price}
              <button onClick={() => handleRemove(item.id)}>Remove</button>
            </li>
          ))}
        </ul>
      ) : (
        <p>Your cart is empty</p>
      )}
    </div>
  );
};

export default Cart;
