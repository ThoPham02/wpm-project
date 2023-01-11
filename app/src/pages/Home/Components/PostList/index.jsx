import React from "react";
import Table from "react-bootstrap/Table";
import PropTypes from "prop-types";
import PostItem from "../PostItem";

PostList.propTypes = {
  posts: PropTypes.array,
};

function PostList(props) {
  const { posts } = props;
  return (
    <Table striped hover className="postList">
      <thead>
        <tr className="postList-head">
          <th className="postList-id">ID</th>
          <th className="postList-name">Name</th>
          <th className="postList-length">Length</th>
          <th className="postList-age">Age</th>
        </tr>
      </thead>
      {posts.map((item, index) => {
        return (
          <PostItem
            key={index}
            id={item.id}
            title={item.title}
            descriptions={item.descriptions}
            content={item.content}
            age={item.age}
          />
        );
      })}
    </Table>
  );
}

export default PostList;
