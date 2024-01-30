import React from 'react'
import EntranceButton from "../Button/EntranceButton";
import InWebsite from "../Button/InWebsite";
import Validation from "../Validation/Validation";
import './Registration.css'

function Registration() {
    const {blurHandler,
        phoneNumberHandler,
        dataHandler,
        phoneNumber,
        phoneNumberDirty,
        phoneNumberError,
        formValid,
        focusedPhoneNumber,
        focusedData,
        dataDirty,
        dataError,
        data} = Validation()

    return (
        <div className="registration">
            <form style={{textAlign: 'center', display: 'flex-block'}}>
                <h1>Let's Lunch</h1>
                <h2>Регистрация</h2>
                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Номер телефона</div>}

                <input className="Input"
                       onChange={e => phoneNumberHandler(e)}
                       value={phoneNumber}
                       onBlur={e => blurHandler(e)}
                       focused={focusedPhoneNumber.toString()}
                       name='phone'
                       type="text"
                       placeholder='Enter your phone number...'
                       style={{border: (focusedPhoneNumber && phoneNumberError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(phoneNumberDirty && phoneNumberError) && <div style={{color: 'red'}}>{phoneNumberError}</div>}

                {<div className="numberPhone">Дата Рождения</div>}
                <input className="Input"
                       onChange={e => dataHandler(e)}
                       value={data}
                       onBlur={e => blurHandler(e)}
                       type="date"
                       focused={focusedData.toString()}
                       style={{border: (focusedData && dataError) ? '1px solid red' : '1px solid #ccc'}}/>
                {(dataDirty && dataError) && <div style={{color: 'red'}}>{dataError}</div>}


                {<div className="numberPhone" style={{color: '#2a1100', textAlign: 'left'}}>Пол</div>}
                <label style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block"}}><input  style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block", marginLeft: '10px', marginRight: '5px'}} type="radio" name="gender" value="male"/>Мужской </label>
                <label style={{ color: '#2a1100',   fontSize: '1.5rem',  display: "inline-block"}}><input  style={{ color: '#2a1100',  fontSize: '1.5rem', display: "inline-block",  marginLeft: '10px',  marginRight: '5px'}} type="radio" name="gender" value="female"/>Женский</label>

                <div className="reg" style={{textAlign: 'center'}}>
                <InWebsite disabled={!formValid}>Зарегистрироваться</InWebsite>
                <EntranceButton className="login-button" >Войти</EntranceButton>
                </div>
            </form>
        </div>
    );
}

export default Registration;