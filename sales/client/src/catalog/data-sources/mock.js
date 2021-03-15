class Mock{
    getProducts() {
        return[
            {
                id: 1,
                img: "https://picsum.photos/id/1/300",
                imgAlt: "a beautiful rendition of product 1",
                name: "Product 1",
                description: "Lorem ispum description of product bla bla bla",
                price: 0.50
            },
            {
                id: 2,
                img: "https://picsum.photos/id/2/300",
                imgAlt: "a beautiful rendition of product 2",
                name: "Product 2",
                description: "Lorem ispum description of product bla bla bla",
                price: 2.50
            },
            {
                id: 3,
                img: "https://picsum.photos/id/3/300",
                imgAlt: "a beautiful rendition of product 2",
                name: "Product 2",
                description: "Lorem ispum description of product bla bla bla",
                price: 12.50
            },
        ]
    }
}

export default Mock;