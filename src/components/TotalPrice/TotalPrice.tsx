import React from 'react';

interface TotalPriceProps {
  amount: number;
}

const TotalPrice: React.FC<TotalPriceProps> = ({ amount }) => {
  // Format the amount to show two decimal places and add a currency symbol (e.g., $).
  const formattedAmount = `$${amount.toFixed(2)}`;

  return (
    <div>
      <h3>Total Price: {formattedAmount}</h3>
    </div>
  );
};

export default TotalPrice;