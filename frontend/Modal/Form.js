import React, {useContext, useState} from 'react'
import Button from "../Button/Button";
import Input from "../components/Input";
import axios from "axios";
import cl from "./Modal.module.css";
import visitContext from "../visitContext";

const Form =  ({setVisible, url, names}) => {
    const [name, setName] = useState('')
    const [time, setTime] = useState([])
    const [modal, setModal] = useState(false);
    const [place, setPlace] = useState('')
    const {meeting, setMeeting} = useContext(visitContext)

    const addNewPost = (e) => {
        e.preventDefault()
        /* sendPostCreateVisit(name, time, placeId) */
    }

    async function fetchVisit(e) {
        e.preventDefault()
        setMeeting({ url: url,
            name: names,
            description:"Вы создали встречу"})
        const response = await axios.get('htpps://jsonplaceholder/typicode.com/posts')
        /*setVisit(response.data)*/
    }

    return (
        <form style={{display:"block", marginTop: "-30px"}}>
            <div style={{textAlign:"right"}} className={cl.closeModal} onClick={() => setVisible(false)}>&times;</div>
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
            <Button onClick={(e) => fetchVisit(e)}>Создать встречу</Button>
        </form>
    );
};

export default Form;