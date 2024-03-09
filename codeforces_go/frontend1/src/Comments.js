import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';

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
                // res.data is an array of objects with blogId, blogTitle, and comments properties
                //console.log(res.data)
                const filteredComments = res.data.groupedComments.find(item => item.id === parseInt(blogIdnum));
                if (filteredComments) {
                    console.log(filteredComments)
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
        <div>
            <h2>Comments for Blog Post {blogIdnum}</h2>
            {comments && (
                <ul>
                    {comments.map((comment, index) => (
                        <li key={index} dangerouslySetInnerHTML={{__html: comment}}/>
                    ))}
                </ul>
            )}
        </div>
    );
};

export default BlogComments;
