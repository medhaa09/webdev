import * as React from 'react';
import { useEffect, useState } from 'react'
import axios from 'axios'

const token = localStorage.getItem('token');
const fetchProtectedData = async () => {
   
  try {
    const token = localStorage.getItem('token');
    const response = await axios.get('http://localhost:8080/activity/recent-actions-grouped', {
      headers: {
        'Authorization': `Bearer ${token}`
      }
    });
   // console.log(response.data);
  } catch (error) {
    console.error("Error fetching protected data:", error);
  }
}

fetchProtectedData();
const Axios = () => {
  const [rawJson, setRawJson] = useState('');
  useEffect(() => {
    const token = localStorage.getItem('token');
    axios.get('http://localhost:8080/activity/recent-actions-grouped', {
        headers: {
            'Authorization': `Bearer ${token}` // Include the token in the Authorization header
        }
    })
      .then((res) => {
        // Convert the JSON object to a string and store it in state
        const jsonString = JSON.stringify(res.data, null, 2); // null and 2 are for formatting
        setRawJson(jsonString);
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
        setRawJson(`Error fetching data: ${error.message}`);
      });
  }, []);
    return (
<div>
      <div/>
      <pre>{rawJson}</pre>
    </div>
    )
      }
  export default Axios;