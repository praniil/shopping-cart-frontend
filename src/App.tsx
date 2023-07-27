import React from 'react';
import logo from './logo.svg';
import './App.css';
import { CartWidget } from './components/CartWidget/CartWidget';
import { Header } from './components/Header/Header';
import { HashRouter as Router, Route, Routes } from 'react-router-dom';
import { Products } from './components/Products/Products';
import { Cart } from './components/Cart/Cart';

function App() {
  const productsCount = 10
  return (
    // <Router>
    //   <Routes>

    //     {/* <Route path="/cartwidget" element={<CartWidget productsCount={productsCount} />} /> */}
    //     <Route path="/" element={<Header />} />

    //     <Route path="/" element={<Products />} />

    //   </Routes>
    // </Router>
    <Router>
      <Header />
      <Products/>
      <Routes>

      <Route path = "/cart" element={<Cart/>} />
      </Routes>
      
    </Router>


  );
}

export default App;
