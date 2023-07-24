import { FunctionComponent, useEffect } from "react";
import useLocalStorageState from "use-local-storage-state";
import classes from './header.module.css'
import { CartWidget } from "../CartWidget/CartWidget";

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
    const shrinkHeader = ()=>{
        const distance_from_top = 140
        const headerElement = document.querySelector("header")
    }

    var productsCount = 10
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