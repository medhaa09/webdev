
import './App.css';
import Blog from './Blog.js';
import BlogComments from './Comments.js';
import Extendedblog from './extendedblog.js';
import Signup from './signup.js';
import LoginForm from './LoginForm.js'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import { useState } from 'react';
const ProtectedRoute = ({ element: Element, ...rest }) => {
  const [isAuthenticated] = useState(() => localStorage.getItem('token') !== null); // Check if user is authenticated

  return isAuthenticated ? <Element {...rest} /> : <Navigate to="/user/signup" replace />;
};
function App() {
  return (
    <div className="App">
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<ProtectedRoute element={Blog} />} />
        <Route path="/extendedblog" element={<ProtectedRoute element={Extendedblog} />} />
        <Route path="/login" element={<LoginForm/>} />
        <Route path="/user/signup" element={<Signup/>} />
        <Route path="/blogs/:blogId" element={<ProtectedRoute element={BlogComments} />}/>
      </Routes>
    </BrowserRouter>
    </div>
  );
}

export default App;
