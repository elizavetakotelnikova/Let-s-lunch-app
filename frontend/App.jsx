import './App.css';
import  "./components/formInput.css";
import React from ".";
import Entrance from "./Entrance/Entrance";
import Registration from "./Refistration/Registration";
import MainPage from "./MainPage/MainPage"
import Account from "./AccountPage/Account"
import {Routes, Route} from "react-router-dom";
import Contacts from "./Contacts/Contact"

function App() {
  return (
      <div className="App">
        <Routes>
          <Route path="/" element={<Entrance/>}/>
          <Route path="/registration" element={<Registration/>}/>
          <Route path="/mainpage" element={<MainPage/>}/>
          <Route path="/mainpage/account" element={<Account/>}/>
          <Route path="/mainpage/contact/" element={<Contacts/>}/>
        </Routes>
      </div>
  );
}

export default App;