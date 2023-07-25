import { FunctionComponent } from "react";
import classes from "./products.module.css"

const API_URL = 'http://localhost:8080/api/newproduct'

export type Product ={
    id : number
    title : string
    price : number
    thumbnail : string
    image : string
    quantity : number
}

export interface CartProps {
    [productId: string] : Product
}