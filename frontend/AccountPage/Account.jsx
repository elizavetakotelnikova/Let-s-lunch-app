import React from "react";


import './account.css';
function Account() {
    return(
        <div className="account">
            <header>
                <a href="#" className="logo"><h1>L</h1>et's Lunch</a>
                <nav className="navbar">
                    <a href="/mainpage/">На главную </a>
                    <a href="/">Выход </a>
                </nav>
            </header>

            <div className="book" id="book">
                <h1 className="heading">Профиль</h1>
                <div className="row">
                    <form className="form" action="">
                        <div className="inputBox">
                            <h3>Имя</h3>
                            <input type="text" placeholder="place name"/>
                        </div>
                        <div className="inputBox">
                            <h3>О себе</h3>
                            <input type="text" className="about" placeholder="place something about you"/>
                        </div>
                        <input type="submit" className="btn" value="Сохранить"/>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Account;