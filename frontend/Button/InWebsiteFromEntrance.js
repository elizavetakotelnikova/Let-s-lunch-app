import React from "react";
import ReactDOM from "react-dom";
import  "../components/formInput.css";

class EntranceButtonFromEntrance extends React.Component {
    onclick () {
        window.location.assign('http://localhost:3000/mainpage');
    }

    render() {
        return (<button className="in-website-from-entrance"
            onClick={(e) => this.onclick(e.preventDefault())}
        >Войти через телеграмм</button>);
    }
}

export default EntranceButtonFromEntrance;