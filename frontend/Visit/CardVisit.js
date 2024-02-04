import Button from "../Button/Button";
import Modal from "../Modal/Modal";
import Form from "../Modal/Form";
import React, {useState} from "react";

const CardVisit = ({url, name, children}) => {
    const [modal, setModal] = useState(false);


    return (
        <div className="card">
            <img src={url} alt={name}
                 className="card_img"/>
            <div className="card_body">
                <h3 className="card_title">{name}</h3>
                <p className="card_desc">{children}</p>
                <Button onClick={() => setModal(true)} className="event">Присоединиться</Button>
                <Modal visible={modal} setVisible={setModal}>
                    <h1>Вы присоеденились!</h1>
                </Modal>

            </div>
        </div>
    );
};

export default CardVisit;