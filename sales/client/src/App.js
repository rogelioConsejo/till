import ProductsList from './catalog/productsList'
import './App.css';
import Addition from "./addition/addition";
import Mock from "./data-sources/mock";
import React, {useState} from 'react';

const dataSource = new Mock();
const products = dataSource.getProducts();


function App() {
    const [chosenItems, setItems] = useState(new Map());

    const updateItems = (name, price, amount) => {
        setItems(new Map(chosenItems.set(name, {price: price, amount: amount})));
    }

    return (
        <div className="App">
            <ProductsList products={products} update={updateItems}/>
            <Addition items={chosenItems}/>
        </div>
    );
}

export default App;
