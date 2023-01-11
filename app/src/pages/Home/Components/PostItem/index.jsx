import React from "react";
import PropTypes from "prop-types";
import { useNavigate } from "react-router-dom";

import { lengthCount, wordsCount } from "../../../../utils/function";

PostItem.propTypes = {
  id: PropTypes.number,
  title: PropTypes.string,
  descriptions: PropTypes.string,
  content: PropTypes.string,
  age: PropTypes.string,
};

function PostItem(props) {
  const { id, title, descriptions, content, age } = props;
  let words = wordsCount(content);
  let length = lengthCount(content);
  const navigate = useNavigate();

  const handleClickpost = () => {
    let path = "/typing/:" + id;
    navigate(path);
  };
  return (
    <tbody>
      <tr className="postItem" onClick={handleClickpost}>
        <td className="postItem-id">{id}</td>
        <td className="postItem-name">
          <h3>{title}</h3>
          <span>{descriptions}</span>
        </td>
        <td className="postItem-length">
          <span>{words}</span>
          <span>{length}</span>
        </td>
        <td className="postItem-age">{age}</td>
      </tr>
    </tbody>
  );
}

export default PostItem;
