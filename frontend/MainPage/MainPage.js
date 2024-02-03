import React, {useState, useEffect, useContext} from "react";
import './styleMain.css';
import Card from "./Card";
import MyCard from "./MyCard";
import axios from "axios";
import NavBar from "../NavBar/NavBar";
import tokenContext from "../tokenContext";
import visitContext from "../visitContext";

function MainPage() {
    const {token, setToken} = useContext(tokenContext)
    const [searchBarActive, setSearchBarActive] = useState(false);
    const {meeting, setMeeting} = useContext(visitContext)
    const [cards, setCards] = useState([{
        url:"https://voyagist.ru/wp-content/uploads/2017/09/pekarni-sankt-peterburga-9.jpg",
        name:"Люди Любят",
        description:"Описания нет, придумайте сами"}]);

    useEffect(() => {
       fetchCard()
        }, []
    )

    async function fetchCard() {
        console.log(token)
        axios.defaults.headers.common['Authorization'] = token;
        const response = await axios.get('http://localhost:3333/api/gatheringPlace/find');
        setCards(response)
    }

    return (
        <div className="Main">

            <NavBar list={[
                {id: 0, link: "/", description: "Выход"},
                {id: 1, link: "/mainpage/account/", description: "Аккаунт"},
                {id: 2, link: "/mainpage/visit/", description: "Вcтречи"},
                {id: 3, link: "/mainpage/contact/", description: "Контакты"}
            ]}/>
            {meeting ?
                <div>
                    <h2>Моя встреча</h2>
                    <div  style={{display: 'flex', justifyContent: 'center'}}>
                    <MyCard url={meeting.url} name={meeting.name}>{meeting.description}</MyCard>
                    </div>
                    <h2>Куда отправляемся?</h2>
                </div>
                : <h2>Куда отправляемся?</h2>
            }

            <div className="list">
                {cards.map(card =>
                <Card url={card.url} name={card.name}>{card.description}</Card>)}
            </div>

        </div>
    );
}

export default MainPage;


