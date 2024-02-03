import React from "react";
import './NavBar.css'

const NavBar = ({list}) => {
    return (
            <header>
                <a href="#" className="logo"><h1>L</h1>et's Lunch</a>
                <div className="navbar">
                    {list.map(element => <a  key={element.id} href={element.link}>{element.description}</a>)}
                </div>
            </header>
        )
}

export default NavBar;