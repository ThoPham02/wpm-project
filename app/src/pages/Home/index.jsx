import React from "react";

import PostList from "./Components/PostList";

function Home() {
  const posts = [
    {
      id: 1,
      title: "Hello World!",
      descriptions: "test",
      content: "test test test",
      age: "20:20 11/01/2023",
    },
  ];
  return (
    <div className="home">
      <div className="home-title">
        <h2>List Post</h2>
      </div>
      <div className="home-search">
        <input type="text" placeholder="Search by name" />
      </div>
      <div className="home-table">
        <PostList posts={posts} />
      </div>
    </div>
  );
}

export default Home;