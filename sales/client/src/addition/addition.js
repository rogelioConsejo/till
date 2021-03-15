import './addition.css'

function Addition(props) {
    const items = props.items;
    let itemsList = [];
    let addition = 0;
    items.forEach((total, product) => {
        addition += total.amount * total.price;
        if (total.amount > 0) {
            itemsList.push(
                <li>{product}: ${total.price.toFixed(2)}c/u x {total.amount} = ${(total.amount * total.price).toFixed(2)}</li>
            );
        }
    });

    return (
        <div id="addition">
            <ul>
                {itemsList}
                <li>TOTAL: ${addition.toFixed(2)}</li>
            </ul>

        </div>
    )
        ;
}

export default Addition;