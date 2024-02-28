import * as React from 'react';
import CssBaseline from '@mui/material/CssBaseline';
import Grid from '@mui/material/Grid';
import Container from '@mui/material/Container';
import { createTheme, ThemeProvider } from '@mui/material/styles';
import Header from './Header';
import MainFeaturedPost from './MainFeaturedPost';
import FeaturedPost from './FeaturedPost';
import Main from './Main';
import Footer from './Footer';

const sections = [];

const mainFeaturedPost = {
  title: 'Title',
  description:
    "Multiple lines of text that form the lede, informing new readers quickly and efficiently about what's most interesting in this post's contents.",
  image: 'https://source.unsplash.com/random?wallpapers',
  imageText: 'main image description'
  
};

const featuredPosts = [
  {
    title: 'Recent Actions',
    date: '',
    description:
      'shows all the data retrieved from the recent actions codeforces api',
    image: 'https://source.unsplash.com/random?wallpapers',
    imageLabel: 'Image Text',
    path:'/extendedblog'
  },
  {
    title: 'Blog1',
    date: '',
    description:
      'This is a wider card with supporting text below as a natural lead-in to additional content.',
    image: 'https://source.unsplash.com/random?wallpapers',
    imageLabel: 'Image Text',
    path:'/blog1'
  },
  {
    title: 'blog2',
    date: '',
    description:
      'This is a wider card with supporting text below as a natural lead-in to additional content.',
    image: 'https://source.unsplash.com/random?wallpapers',
    imageLabel: 'Image Text',
    path:'/blog2'
  },
  {
    title: 'blog3',
    date: '',
    description:
      'This is a wider card with supporting text below as a natural lead-in to additional content.',
    image: 'https://source.unsplash.com/random?wallpapers',
    imageLabel: 'Image Text',
    path:'/blog3'
  }
  
  
];

// TODO remove, this demo shouldn't need to reset the theme.
const defaultTheme = createTheme();


export default function Blog() {
  
  return (
    <ThemeProvider theme={defaultTheme}>
      <CssBaseline />
      <Container maxWidth="lg">
        <Header title="Blog" sections={sections} />
        <main>
          <MainFeaturedPost post={mainFeaturedPost} />
          <Grid container spacing={4}>
            {featuredPosts.map((post) => (
              <FeaturedPost key={post.title} post={post} />
            ))}
          </Grid>
          
        </main>
      </Container>
      <Footer
        title="Footer"
        description="The End!"
      />
    </ThemeProvider>
  );
}
