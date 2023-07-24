import { FunctionComponent } from "react";
import classes from './header.module.scss'

import {
    BrowserRouter,
    Link
} from "react-router-dom"


export const Header: FunctionComponent = () => {
    return (
        <header className={classes.header}>
            <div>
                <Link to="/">
                    <img src="logo192.png" className={classes.logo} alt="Shopping Cart Application" />
                </Link>
            </div>
            <div>
                <CartWidget productsCount = {productsCount} />
            </div>
        </header>
    )
}