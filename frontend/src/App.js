import React from 'react'
import './App.css'
import TopAppBar from './components/Appbar/TopAppBar'
import Form from './components/Form/Form'
import { BrowserRouter, Routes, Route } from 'react-router-dom'
function App() {
  return (
    <div className="App">
      <BrowserRouter>
        <TopAppBar />
        <Routes>
          <Route path='/' element={<Form />}/>
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
