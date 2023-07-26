import { FunctionComponent, useEffect } from "react";
import useLocalStorageState from "use-local-storage-state";
import classes from './header.module.css'
import { CartWidget } from "../CartWidget/CartWidget";
import { CartProps } from "../Products/Products"

import {
    BrowserRouter,
    Link
} from "react-router-dom"


export const Header: FunctionComponent = () => {
    useEffect(() => {
        window.addEventListener("scroll", () => shrinkHeader(), false)
        return () => {
            window.removeEventListener("scroll", () => shrinkHeader())
        }
    }, []) //empty array := only once executed
    const shrinkHeader = () => {
        const distance_from_top = 140
        const headerElement = document.querySelector("header") as HTMLElement
        const logoElement = document.querySelectorAll("img")[0] as HTMLElement
        const cartWidgetElement = document.querySelectorAll("img")[1] as HTMLElement
        const productsCountElement = document.querySelector("span") as HTMLElement
        const scrollY = document.body.scrollTop || document.documentElement.scrollTop
        if (scrollY > distance_from_top) {
            headerElement.style.transition = "height 200ms ease-in"
            headerElement.style.height = "80px"
            logoElement.style.transition = "height 200ms ease-in"
            logoElement.style.height = "4rem"
            cartWidgetElement.style.transition = "height 200ms ease-in"
            cartWidgetElement.style.height = "2rem"
            productsCountElement.style.transition = "height 200ms ease-in"
            productsCountElement.style.fontSize = "20px"
        } else {
            headerElement.style.height = "150px"
            logoElement.style.height = "6rem"
            cartWidgetElement.style.height = "3rem"
            productsCountElement.style.fontSize = "2em"
        }
    }
    const [cart, ] = useLocalStorageState<CartProps>('cart', {})

    var productsCount : number = Object.keys(cart || {}).length
    return (
        <header className={classes.header}>
            <div>
                <Link to="/">
                    <img src="logo192.png" className={classes.logo} alt="Shopping Cart Application" />
                </Link>
            </div>
            <div>
                <CartWidget productsCount={productsCount} />
            </div>
        </header>
    )
}