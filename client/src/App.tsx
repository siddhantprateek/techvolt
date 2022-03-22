import React from 'react';
import './App.css';
import { Header } from './components';
import { Routes, Route } from "react-router-dom";
import { Explore, Collections, Stats, NoMatch } from './pages';
import 'swiper/css';
import 'swiper/css/navigation';
import 'swiper/css/pagination';
import 'swiper/css/scrollbar';
import 'swiper/css/autoplay';
// import { Layout, Menu, Breadcrumb } from 'antd';
// const { Header, Content, Footer } = Layout;
function App() {
  return (
    <div className="App">
      <Header/>
      <Routes>
        <Route path="/" element={<Explore/>} />
        <Route path="/collections" element={<Collections/>} />
        <Route path="/stats" element={<Stats/>} />
        <Route path="/*" element={<NoMatch/>} />
      </Routes>
    </div>
  );
}

export default App;
