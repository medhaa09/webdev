import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Grid from '@mui/material/Grid';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Header from './Header';
import MainFeaturedPost from './MainFeaturedPost';
import FeaturedPost from './FeaturedPost';
import Main from './Main';
import { useEffect, useState } from 'react'
import axios from 'axios';
const sections = [];
const mainFeaturedPost = {
  title: 'Codeforces Recent Actions',
  description:
    "Here are the various blog posts with their corresponding comments.",
};

const defaultTheme = createTheme();
export default function Blog() { 
  const [featuredPosts, setFeaturedPosts] = useState([])

  const fetchProtectedData = async () => {
    axios.get('http://localhost:8080/activity/recent-actions-grouped', {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`
      }
    });
  }
  fetchProtectedData();
  useEffect(() => {
    const token = localStorage.getItem('token'); // Retrieve the token from localStorage
    axios.get('http://localhost:8080/activity/recent-actions-grouped', {
        headers: {
            'Authorization': `Bearer ${token}` // Include the token in the Authorization header
        }
    })
    .then((res) => {
        const posts = res.data.groupedComments.map(item => ({
            title: item.title,
            author: item.authorHandle,
            description: `comments for blog post ${item.id}`,
            path: `/blogs/:${item.id}`
        }));
        setFeaturedPosts(posts); // Update your state with the fetched posts
    })
    .catch((error) => {
        console.error("Error fetching data:", error);
    });
}, []);
  return (
    <ThemeProvider theme={defaultTheme}>
      <CssBaseline />
      <Container maxWidth="lg">
        <Header title="Blog" />
        <main>
          <MainFeaturedPost post={mainFeaturedPost} />
          <Grid container spacing={4}>
            {featuredPosts.map((post, index) => (
              <FeaturedPost key={index} post={post} />
            ))}
          </Grid>
        </main>
      </Container>
    </ThemeProvider>
  );
}