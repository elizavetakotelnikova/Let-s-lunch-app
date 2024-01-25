import React, { useState } from "react";
import ReactDOM from "react-dom";
import "../components/formInput.css";

class InWebsite extends React.Component {
    onclick() {
        if (!this.props.disabled) {
            window.location.assign('http://localhost:3000/mainpage');
        }
    }

    render() {
        return (
            <button
                className="in-website"
                onClick={(e) => this.onclick(e.preventDefault())}
                disabled={this.props.disabled}
            >Зарегистрироваться</button>
        );
    }
}

export default InWebsite;