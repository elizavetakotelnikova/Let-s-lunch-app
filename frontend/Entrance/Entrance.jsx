import React, {useState, useEffect} from 'react'
import RegistrationButton from "../Button/RegistrationButton";
import '../App.css';
import InWebsiteFromEntrance from "../Button/InWebsiteFromEntrance";
function Entrance() {
        return(
        <div className="app">
            <form>
                <h1>Let's Lunch</h1>
                <h2>Войти</h2>
                <InWebsiteFromEntrance>Войти через телеграмм</InWebsiteFromEntrance>
                <RegistrationButton>Registration</RegistrationButton>
            </form>
        </div>
    )
}

export default Entrance;