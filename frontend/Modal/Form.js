import React, {useContext, useState} from 'react'
import Button from "../Button/Button";
import Input from "../components/Input";
import axios from "axios";
import cl from "./Modal.module.css";
import visitContext from "../visitContext";
import tokenContext from "../tokenContext";

const Form =  ({setVisible, url, names, id}) => {
    const {token, setToken} = useContext(tokenContext)
    const [name, setName] = useState('')
    const [time, setTime] = useState([])
    const [modal, setModal] = useState(false);
    const [place, setPlace] = useState('')
    const {meeting, setMeeting} = useContext(visitContext)
    const config = {headers: { Authorization: `Bearer ${token}` }}


    async function fetchVisit(e) {
        e.preventDefault()
        setMeeting({ url: url,
            name: names,
            description:"Вы создали встречу"})
        const response = await axios.post('http://localhost:3333/api/meeting/create', {
            gatheringPlaceId: id,
            initiatorsId: "1671f0f0-5c27-4aa9-96c8-87ed8f0e272d",
            startTime: "2024-01-30T18:38:25.125Z",
            endTime: "2024-01-30T18:38:25.125Z",
            usersQuantity: 2,
            state: 0
        }, config)
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