import React from "react";
import ReactDOM from "react-dom";
import  "../components/formInput.css";

class EntranceButton extends React.Component {
    onclick () {
        window.location.assign('http://localhost:3000/');
    }
    render() {
        return (<button  className="login-button" onClick={(e) => this.onclick(e.preventDefault())}>Войти</button>);
    }
}
export default EntranceButton;