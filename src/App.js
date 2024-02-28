
import './App.css';
import Blog from './Blog.js';
import Extendedblog from './extendedblog.js';
import Signup from './signup.js';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import LoremIpsum from './dummtext.js'
function App() {
  return (
    <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Blog />} />
        <Route path="/extendedblog" element={<Extendedblog/>} />
        <Route path="/signup" element={<Signup/>} />
        <Route path="/blog1" element={<LoremIpsum />}/>
        <Route path="/blog2" element={<LoremIpsum />}/>
        <Route path="/blog3" element={<LoremIpsum />}/>
      </Routes>
    </BrowserRouter>
    </div>
  );
}

export default App;
