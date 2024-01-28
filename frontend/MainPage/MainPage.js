import React, { useState, useEffect } from "react";
import './styleMain.css';
import Modal from "../Modal/Modal";
import Button from "../Button/Button";
import Form from "../Modal/Form";

function MainPage() {
    const [searchBarActive, setSearchBarActive] = useState(false);
    const [modal, setModal] = useState(false);

    const handleSearchClick = () => {
        setSearchBarActive(!searchBarActive);
    };

    useEffect(() => {
        const searchBtn = document.querySelector('#search-btn');
        const searchBar = document.querySelector('.search-bar-container');

        const handleScroll = () => {
            searchBtn.classList.remove('fa-times');
            searchBar.classList.remove('active');
        };

        window.addEventListener('scroll', handleScroll);

        return () => {
            window.removeEventListener('scroll', handleScroll);
        };
    }, []); // Здесь нет зависимостей, так как searchBtn и searchBar получаются один раз при монтировании



    return (
        <div className="Main">
            <header>
                <link rel="stylesheet"
                      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/5.15.3/css/all.min.css"/>
                <link rel="stylesheet" href="css/style.css"/>

                <a href="#" className="logo"><h1>L</h1>et's Lunch</a>
                <nav className="navbar">
                    <a href="/">Выход</a>
                    <a href="/mainpage/account/"> Аккаунт</a>
                    <a href="/mainpage/contact/"> Контакты</a>
                </nav>
                <div className="icons">
                    <i className="fas fa-search" id="search-btn" onClick={handleSearchClick}></i>
                    <i className="fas fa-user" id="login-btn"></i>
                </div>
                <form action="" className={`search-bar-container ${searchBarActive ? 'active' : ''}`}>
                    <input type="search" id="search-bar" placeholder="Search"/>
                    <label htmlFor="search-bar" className="fas fa-search"></label>
                </form>
            </header>
            <h2>Куда отправляемся?</h2>

            <div className="list">
                <div className="card">
                    <img src="https://i1.photo.2gis.com/images/branch/38/5348024604197800_8158.jpg" alt="Теремок"
                         className="card_img"/>
                    <div className="card_body">
                        <h3 className="card_title">Теремок</h3>
                        <p className="card_desc">Описание теремка пока нет, поэтому придумайте сами</p>
                        <Button onClick={() => setModal(true)} className="event">Назначить встречу</Button>
                        <Modal visible={modal} setVisible={setModal}>
                            <Form />
                        </Modal>

                    </div>
                </div>


                <div className="card">
                    <img src="https://i3.photo.2gis.com/images/branch/38/5348024573203103_a167.jpg" alt="Теремок"
                         className="card_img"/>
                    <div className="card_body">
                        <h3 className="card_title">Вольчика</h3>
                        <p className="card_desc">Описание вольчика пока нет, поэтому придумайте сами</p>
                        <Button onClick={() => setModal(true)} className="event">Назначить встречу</Button>
                        <Modal visible={modal} setVisible={setModal}>
                            <Form />
                        </Modal>

                    </div>
                </div>

                <div className="card">
                    <img src="https://api.gzktour.ru/storage/media/435318/caption.jpg" alt="Пельменная"
                         className="card_img"/>
                    <div className="card_body">
                        <h3 className="card_title">Пельменная</h3>
                        <p className="card_desc">Описание пельменной пока нет, поэтому придумайте сами</p>
                        <Button onClick={() => setModal(true)} className="event">Назначить встречу</Button>
                        <Modal visible={modal} setVisible={setModal}>
                            <Form />
                        </Modal>
                    </div>
                </div>

                <div className="card">
                    <img src="https://voyagist.ru/wp-content/uploads/2017/09/pekarni-sankt-peterburga-9.jpg" alt="Люди Любят"
                         className="card_img"/>
                    <div className="card_body">
                        <h3 className="card_title">Люди любят</h3>
                        <p className="card_desc">Описание люди любят пока нет, поэтому придумайте сами</p>
                        <Button onClick={() => setModal(true)} className="event">Назначить встречу</Button>
                        <Modal visible={modal} setVisible={setModal}>
                            <Form />
                        </Modal>
                    </div>
                </div>

            </div>
        </div>
    );
}

export default MainPage;
