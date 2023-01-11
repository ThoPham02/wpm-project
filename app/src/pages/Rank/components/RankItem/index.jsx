import React from "react";
import PropTypes from "prop-types";

RankItem.propTypes = {
  rank: PropTypes.number,
  id: PropTypes.number,
  name: PropTypes.string,
  wpm: PropTypes.number,
  age: PropTypes.string,
};

function RankItem(props) {
  const { rank, id, name, wpm, age } = props;
  return (
    <tbody>
      <tr className="rankItem">
        <td className="rankItem-rank">{rank}</td>
        <td className="rankItem-name">
          <h4>{id + ". " + name}</h4>
        </td>
        <td className="rankItem-wpm">{wpm}</td>
        <td className="rankItem-age">{age}</td>
      </tr>
    </tbody>
  );
}

export default RankItem;
