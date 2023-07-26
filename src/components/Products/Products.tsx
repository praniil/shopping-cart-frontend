import { FunctionComponent, useEffect, useState } from "react";
import classes from "./products.module.css"
import useLocalStorageState from "use-local-storage-state";

const API_URL = 'https://dummyjson.com/products'

export type Product = {
    id: number
    title: string
    price: number
    thumbnail: string
    image: string
    quantity: number
}
//interface is a way to define a contract or structure for an object
export interface CartProps {
    [productId: string]: Product       //key of type string
}

export const Products: FunctionComponent = () => {
    const [products, setProducts] = useState<Product[]>([])
    const [error, setError] = useState(false)
    const [cart, setCart] = useLocalStorageState<CartProps>('cart', {})


    useEffect(() => {
        fetchData(API_URL)
    }, [])

    async function fetchData(url: string) {
        try {
            const response = await fetch(url)
            if (response.ok) {
                const data = await response.json()
                setProducts(data.products)

            } else {
                setError(true)

            }
        } catch (error) {
            setError(true)

        }
    }

    const addToCart = (product: Product): void => {
        product.quantity = 1

        setCart((prevCart) => ({
            ...prevCart,
            [product.id]: product,
        }))
    }

    const isInCart = (productId: number): boolean => Object.keys(cart || {}).includes(productId.toString())
    if (error) {
        return <h3 className={classes.error}> An error occured when fetching data. Please check the API and try again.</h3>
    }

    return (
        
        <section className={classes.productPage}>
            
            <h1>Products</h1>
            <div className={classes.container}>
                {products.map(product => (
                    <div className={classes.product} key={product.id} >
                        <img width="350" height= "250" src={product.thumbnail} alt={product.title} />
                        <h3>{product.title}</h3>
                        <p> Price :  {product.price}</p>
                        <button disabled={isInCart(product.id)} onClick={() => addToCart(product)}> Add to Cart </button>
                    </div>
                ))}
            </div>
        </section>
    )
}