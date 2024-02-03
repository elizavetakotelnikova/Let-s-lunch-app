import React, {useState, useEffect} from "react";
import './styleMain.css';
import Card from "./Card";
import axios from "axios";
import NavBar from "../NavBar/NavBar";

function MainPage() {
    const [searchBarActive, setSearchBarActive] = useState(false);
    const [cards, setCards] = useState([{
        url:"https://voyagist.ru/wp-content/uploads/2017/09/pekarni-sankt-peterburga-9.jpg",
        name:"Люди Любят",
        description:"Описания нет, придумайте сами"}]);

    useEffect(() => {
       fetchCard()
        }, []
    )

    async function fetchCard() {
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
            <h2>Куда отправляемся?</h2>

            <div className="list">
                {cards.map(card =>
                <Card url={card.url} name={card.name}>{card.description}</Card>)}
            </div>

        </div>
    );
}

export default MainPage;


