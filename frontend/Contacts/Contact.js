import React from "react";


import './Contacts.css';
import NavBar from "../NavBar/NavBar";
function Contact() {
    return (
        <div className="contact">

            <NavBar list={[
                {id: 0, link: "/mainpage/", description: "На главную"},
            ]}/>

            <div className="info">
            <h2>Над приложением работали:</h2>
            </div>
        </div>
    )
}

export default Contact;