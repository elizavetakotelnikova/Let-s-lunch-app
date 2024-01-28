import React from "react";


import './Contacts.css';
function Contact() {
    return (
        <div className="contact">
            <header>
                <a href="#" className="logo"><h1>L</h1>et's Lunch</a>
                <nav className="navbar">
                    <a href="/mainpage/">На главную </a>
                    <a href="/">Выход </a>
                </nav>
            </header>
            <div className="info">
            <h2>Над приложением работали:</h2>
            </div>
        </div>
    )
}

export default Contact;