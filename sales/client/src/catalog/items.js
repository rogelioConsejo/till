import Mock from './data-sources/mock'
import './items.css'

function Items() {
    const dataSource = new Mock()
    const products = dataSource.getProducts()
    const output = products.map((product) =>
        <li>
            <img src={product.img} alt={product.imgAlt} className="product-image"/>
            <div className="product-details">
                <h2>{product.name}</h2>
                <p>{product.description}</p>
                <p className="product-price">{product.price}</p>
            </div>
            <button>+</button>
        </li>
    );
    return(
        <ul id="product-list">
            {output}
        </ul>
    );
}

export default Items;