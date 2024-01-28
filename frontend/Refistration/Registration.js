import React, {useEffect, useState} from 'react'
import EntranceButton from "../Button/EntranceButton";
import InWebsite from "../Button/InWebsite";
import Validation from "../Validation/Validation";

function Registration() {
    const {blurHandler,  phoneNumberHandler, phoneNumber, phoneNumberDirty, phoneNumberError, formValid,  focusedPhoneNumber} = Validation()

    return (
        <div className="app">
            <form>
                <h1>Let's Lunch</h1>
                <h2>Регистрация</h2>
                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'} }>Номер телефона</div>}
                <input onChange={e => phoneNumberHandler(e)} value={phoneNumber} onBlur={e => blurHandler(e)}  focused={focusedPhoneNumber.toString()} name='phone' type="text"
                       placeholder='Enter your phone number...' style={{ border: (focusedPhoneNumber && phoneNumberError) ? '1px solid red' : '1px solid #ccc' }}/>
                {(phoneNumberDirty && phoneNumberError) && <div style={{color: 'red'}}>{phoneNumberError}</div>}
                <InWebsite disabled={!formValid}>Зарегистрироваться</InWebsite>
                <EntranceButton>Войти</EntranceButton>
            </form>
        </div>
    );
}

export default Registration;