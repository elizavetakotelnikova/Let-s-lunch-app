import React from "react";


import './account.css';
function Account() {
    return(
        <div className="account">
            <header>
                <a href="#" class="logo"><h1>L</h1>et's Lunch</a>
                <nav class="navbar">
                    <a href="/mainpage/">home </a>
                    <a href="/chat/">chat</a>
                </nav>
            </header>

            <div className="book" id="book">
                <h1 className="heading"></h1>
                <div className="row">
                    <form action="">
                        <div className="inputBox">
                            <h3>Имя</h3>
                            <input type="text" placeholder="place name"/>
                        </div>
                        <div className="inputBox">
                            <h3>О себе</h3>
                            <input type="text" placeholder="place something about you"/>
                        </div>
                        <input type="submit" className="btn" value="save"/>
                    </form>
                </div>
            </div>
        </div>
    )
}

export default Account;