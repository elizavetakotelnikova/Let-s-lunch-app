import './App.css';
import  "./components/formInput.css";
import React from ".";
import Entrance from "./Entrance/Entrance";
import Registration from "./Registration/Registration";
import MainPage from "./MainPage/MainPage"
import Account from "./AccountPage/Account"
import {Routes, Route} from "react-router-dom";
import Contacts from "./Contacts/Contact"
import Visit from "./Visit/Visit";

function App() {
  return (
      <div className="App">
        <Routes>
          <Route path="/" element={<Entrance/>}/>
          <Route path="/registration" element={<Registration/>}/>
          <Route path="/mainpage" element={<MainPage/>}/>
          <Route path="/mainpage/account" element={<Account/>}/>
          <Route path="/mainpage/contact/" element={<Contacts/>}/>
          <Route path="/mainpage/visit/" element={<Visit/>}/>
        </Routes>
      </div>
  );
}

export default App;