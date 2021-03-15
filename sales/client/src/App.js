import ProductsList from './catalog/productsList'
import './App.css';
import Addition from "./addition/addition";
import Mock from "./data-sources/mock";
import React, {useState} from 'react';

const dataSource = new Mock();
const products = dataSource.getProducts();


function App() {
    const [chosenItems, setItems] = useState(new Map());
    const [search, setSearch] = useState("");

    const updateItems = (name, price, amount) => {
        setItems(new Map(chosenItems.set(name, {price: price, amount: amount})));
    }

    const filteredProducts = products.filter(product => product.name.toLowerCase().includes(search.toLowerCase()));

    return (
        <div className="App">
            <div className="SearchBar"><SearchBar update={setSearch}/></div>
            <ProductsList products={filteredProducts} update={updateItems}/>
            <Addition items={chosenItems}/>
        </div>
    );
}

function SearchBar(props) {
    return <form>
        <label htmlFor="search-bar">Search: </label>
        <input type="search" id="search-bar" onChange={(component)=>props.update(component.target.value)}/>
    </form>
}

export default App;
