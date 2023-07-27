import React from 'react';
import { Operation } from './Operation' // Assuming this is the correct path to the Operation type.

interface QuantifierProps {
  removeProductCallback: () => void;
  productId: number;
  handleUpdateQuantity: (productId: number, operation: Operation) => void;
}

const Quantifier: React.FC<QuantifierProps> = ({
  removeProductCallback,
  productId,
  handleUpdateQuantity,
}) => {
  return (
    <div>
      <button onClick={() => handleUpdateQuantity(productId, 'decrease')}>-</button>
      <button onClick={() => handleUpdateQuantity(productId, 'increase')}>+</button>
      <button onClick={removeProductCallback}>Remove</button>
    </div>
  );
};

export default Quantifier;
