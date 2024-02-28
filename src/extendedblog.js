import * as React from 'react';
import { useEffect, useState } from 'react'
import axios from 'axios'

const Axios = () => {
  const [rawJson, setRawJson] = useState('');
  useEffect(() => {
    axios.get('http://localhost:8080/activity/recent-actions')
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