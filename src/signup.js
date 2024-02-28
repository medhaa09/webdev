import React from 'react';
import './signup.css';
import { useState }from 'react';
import axios from 'axios';

export default function SignUpForm() {
  const [name, setName] = useState('');
  const [pass, setPass] = useState('');
  const [handle, setHandle] = useState('');
  const handlechange=(c)=>{
    c.preventDefault()  //the default behaviour will be that this function will try to navigate to the route below on click but we dont want that so we use preventDefault()
    axios.post('http://localhost:8080/user/signup',{name, handle, pass})
    .then((res) => {
     alert('Signup Successful')
    })
    .catch((error) => {
      alert('error occured:' + error.message)
    });
}
  return (
    <div className="signup-form">
      <div className="header">
        <h1>Sign Up</h1>
        <p>Enter your information to create an account</p>
      </div>
      <form className="form"  onSubmit={handlechange}>
        <div className="form-group">
          <label htmlFor="username">Username</label>
          <input id="username" placeholder="Enter your username" onChange={(e)=>setName(e.target.value)}required />
        </div>
        <div className="form-group">
          <label htmlFor="username">Codeforces Handle</label>
          <input id="cfhandle" placeholder="Enter your handle" onChange={(e)=>setHandle(e.target.value)} required />
        </div>
        <div className="form-group">
          <label htmlFor="password">Password</label>
          <input id="password" type="password" placeholder="Enter your password" onChange={(e)=>setPass(e.target.value)} required />
        </div>
        <button type="submit" className="submit-btn">Sign Up</button>
      </form>
    </div>
  );
}
