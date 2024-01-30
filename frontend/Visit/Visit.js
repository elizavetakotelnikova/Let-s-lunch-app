import React from "react";


import './Visit.css';
function Visit() {
    return (
        <div className="account">
            <header>
                <a href="#" className="logo"><h1>L</h1>et's Lunch</a>
                <nav className="navbar">
                    <a href="/mainpage/">На главную </a>
                    <a href="/mainpage/account/">Аккаунт</a>
                    <a href="/">Выход </a>
                </nav>
            </header>
        </div>
    );
}

export default Visit;