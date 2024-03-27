import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardActionArea from '@mui/material/CardActionArea';
import CardContent from '@mui/material/CardContent';
import Header from './Header';

const BlogComments = () => {
    const { blogId } = useParams();
    const [comments, setComments] = useState([]);
     const blogIdnum = blogId.replace(":", "")
     useEffect(() => {
        const token = localStorage.getItem('token');
        axios.get('http://localhost:8080/activity/recent-actions-grouped', {
            headers: {
                'Authorization': `Bearer ${token}` // Include the token in the Authorization header
            }
        })
            .then((res) => {
                const filteredComments = res.data.groupedComments.find(item => item.id === parseInt(blogIdnum));
                if (filteredComments) {
                   // console.log(filteredComments)
                    setComments(filteredComments.comment);
                } else {
                    console.error("No comments found for blog ID:", blogIdnum);
                }
            })
            .catch((error) => {
                console.error("Error fetching data:", error);
            });
            console.log(comments)
    }, [blogIdnum]);
    
    return (
        <div style={{backgroundImage:'linear-gradient(12deg, rgba(193, 193, 193,0.05) 0%, rgba(193, 193, 193,0.05) 2%,rgba(129, 129, 129,0.05) 2%, rgba(129, 129, 129,0.05) 27%,rgba(185, 185, 185,0.05) 27%, rgba(185, 185, 185,0.05) 66%,rgba(83, 83, 83,0.05) 66%, rgba(83, 83, 83,0.05) 100%),linear-gradient(321deg, rgba(240, 240, 240,0.05) 0%, rgba(240, 240, 240,0.05) 13%,rgba(231, 231, 231,0.05) 13%, rgba(231, 231, 231,0.05) 34%,rgba(139, 139, 139,0.05) 34%, rgba(139, 139, 139,0.05) 71%,rgba(112, 112, 112,0.05) 71%, rgba(112, 112, 112,0.05) 100%),linear-gradient(236deg, rgba(189, 189, 189,0.05) 0%, rgba(189, 189, 189,0.05) 47%,rgba(138, 138, 138,0.05) 47%, rgba(138, 138, 138,0.05) 58%,rgba(108, 108, 108,0.05) 58%, rgba(108, 108, 108,0.05) 85%,rgba(143, 143, 143,0.05) 85%, rgba(143, 143, 143,0.05) 100%),linear-gradient(96deg, rgba(53, 53, 53,0.05) 0%, rgba(53, 53, 53,0.05) 53%,rgba(44, 44, 44,0.05) 53%, rgba(44, 44, 44,0.05) 82%,rgba(77, 77, 77,0.05) 82%, rgba(77, 77, 77,0.05) 98%,rgba(8, 8, 8,0.05) 98%, rgba(8, 8, 8,0.05) 100%),linear-gradient(334deg, hsl(247,0%,2%),hsl(247,0%,2%))'
        ,backgroundSize: '100% 100vh',backgroundRepeat: 'repeat-y', color:'white'}}>
           <Header/>
            <h1></h1>
            {comments && (
                <div>
                    <Grid container flexDirection= 'column' justifyContent="center" alignItems="center" spacing={2}>
                    {comments.map((comment, index) => (
                     
                      <Grid item xs={false} md={6} style={{ width: '100%', maxHeight:'70%' } }>
                      <CardActionArea component="a" href="#">
                        <Card sx={{ display: 'flex',flexDirection: 'column' }}>
                          <CardContent sx={{ flex: 1}}>
                            <Typography dangerouslySetInnerHTML={{__html: comment}} component="h5" variant="h5"/>
                          </CardContent>
                        </Card>
                      </CardActionArea>
                    </Grid>
                    ))}
                    </Grid>
                 </div>
            )}
        </div>
    );
};

export default BlogComments;
