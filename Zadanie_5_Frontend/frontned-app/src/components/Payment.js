import React, { useState } from 'react';
import axios from 'axios';

const Payments = ({ cartItems }) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(false);

  const handlePayment = async () => {
    setLoading(true);
    try {
      await axios.post('http://localhost:8080/payments', { items: cartItems });
      setSuccess(true);
    } catch (error) {
      setError(error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <h1>Payments</h1>
      <button onClick={handlePayment} disabled={loading || cartItems.length === 0}>
        {loading ? 'Processing...' : 'Pay Now'}
      </button>
      {error && <p>Error processing payment: {error.message}</p>}
      {success && <p>Payment successful!</p>}
    </div>
  );
};

export default Payments;
