import React, {useEffect, useState} from 'react'


function Validation() {
    const [phoneNumber, setphoneNumber] = useState('')
    const [phoneNumberDirty, setphoneNumberDirty] = useState(false)
    const [phoneNumberError, setphoneNumberError] = useState('Номер не может быть пустым')
    const [formValid, setFormValid] = useState(false)
    const [focusedPhoneNumber, setFocusedPhoneNumber] = useState(false)
    const phoneNumberHandler = (e) => {
        setphoneNumber(e.target.value)
        const re = /^[\\+]?[(]?[0-9]{3}[)]?[-\s\\.]?[0-9]{3}[-\s\\.]?[0-9]{4,6}$/i;
        if (!re.test(String(e.target.value).toLowerCase())) {
            setphoneNumberError('Некорректный номер')
        } else {
            setphoneNumberError('')
        }
    }

    useEffect(() => {
        if (phoneNumberError) {
            setFormValid(false)
        } else {
            setFormValid(true)
        }
    }, [phoneNumberError])
    const blurHandler = (e) => {
        switch (e.target.name) {
            case 'phone':
                setphoneNumberDirty(true)
                setFocusedPhoneNumber(true)
                break
            default:
                break
        }
    }

    return {blurHandler, phoneNumberHandler, phoneNumber, setphoneNumber, phoneNumberDirty, setphoneNumberDirty, phoneNumberError, setphoneNumberError, formValid, setFormValid, focusedPhoneNumber, setFocusedPhoneNumber};
}

export default Validation;