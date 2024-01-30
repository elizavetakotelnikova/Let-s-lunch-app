import React, {useState} from 'react'
import Button from "../Button/Button";
import Input from "../components/Input";

const Form =  () => {
    const [name, setName] = useState('')
    const [time, setTime] = useState([])
    const [modal, setModal] = useState(false);
    const [place, setPlace] = useState('')

    const addNewPost = (e) => {
        e.preventDefault()
        /* sendPostCreateVisit(name, time, placeId) */
    }

    return (
        <form>
            <Input
                value={name}
                onChange={e => setName(e.target.value)}
                type="text"
                placeholder="Ваше имя"
                />
            <Input
                value={time}
                onChange={e => setTime(e.target.value)}
                type="time"

                placeholder="Время"
            />
            <Button onClick={addNewPost}>Создать встречу</Button>
        </form>
    );
};

export default Form;