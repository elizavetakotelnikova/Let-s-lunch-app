import React, {useEffect, useState} from 'react'


function Validation() {
    const [phoneNumber, setphoneNumber] = useState('')
    const [data, setData] = useState('')
    const [phoneNumberDirty, setphoneNumberDirty] = useState(false)
    const [dataDirty, setDataDirty] = useState(false)
    const [phoneNumberError, setphoneNumberError] = useState('Номер не может быть пустым')
    const [dataError, setDataError] = useState('Дата не может быть пустой')
    const [formValid, setFormValid] = useState(false)
    const [focusedPhoneNumber, setFocusedPhoneNumber] = useState(false)
    const [focusedData, setFocusedData] = useState(false)
    const phoneNumberHandler = (e) => {
        setphoneNumber(e.target.value)
        const re = /^[\\+]?[(]?[0-9]{3}[)]?[-\s\\.]?[0-9]{3}[-\s\\.]?[0-9]{4,6}$/i;
        if (!re.test(String(e.target.value).toLowerCase())) {
            setphoneNumberError('Некорректный номер')
        } else {
            setphoneNumberError('')
        }
    }

    const dataHandler = (e) => {
        setData(e.target.value)
        if (e.target.value.length !== 0) {
            setDataError('')
        }
    }

    useEffect(() => {
        if (phoneNumberError || dataError) {
            setFormValid(false)
        } else {
            setFormValid(true)
        }
    }, [phoneNumberError, dataError])


    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'phone':
                setphoneNumberDirty(true)
                setFocusedPhoneNumber(true)
                break
            case 'date':
                setDataDirty(true)
                setFocusedData(true)
                break
            default:
                break
        }
    }

    return {blurHandler,
        phoneNumberHandler,
        phoneNumber,
        setphoneNumber,
        phoneNumberDirty,
        setphoneNumberDirty,
        phoneNumberError,
        setphoneNumberError,
        formValid,
        setFormValid,
        focusedPhoneNumber,
        setFocusedPhoneNumber,
        focusedData,
        dataDirty,
        dataError,
        dataHandler,
        data
    };
}

export default Validation;