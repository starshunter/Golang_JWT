import React from 'react';

const Home = (props) => {
    return (
        <div>
            {props.name ? "Welcome back " + props.name : "You are not logged in"}
        </div>
    );
};

export default Home;