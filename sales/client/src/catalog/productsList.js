import './products.css'
import {useState} from 'react';

function ProductsList(props) {
    const output = props.products.map((product) => <Product product={product} key={product.id} update={props.update}/>
    );
    return (
        <ul id="product-list">
            {output}
        </ul>
    );
}

function Product(props) {
    const [count, setCount] = useState(0);
    const product = props.product;
    function updateCount(newValue) {
        newValue = newValue < 0 ? 0 : newValue;
        setCount(parseInt(newValue));
        props.update(product.name,product.price, newValue);
    }
    return <li>
        <img src={product.img} alt={product.imgAlt} className="product-image"/>
        <div className="product-details">
            <h2>{product.name}</h2>
            <p>{product.description}</p>
            <p className="product-price">${product.price.toFixed(2)}</p>
        </div>
        <div className="buttons">
            <button onClick={()=>updateCount(count - 1)}>-</button>
            <button onClick={()=>updateCount(count + 1)}>+</button>
        </div>
        <input className="product-amount" type="number" step={1} value={count} onChange={(component)=>updateCount(component.target.value)}/>
    </li>;
}

export default ProductsList;