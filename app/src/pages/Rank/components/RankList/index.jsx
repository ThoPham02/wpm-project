import React from "react";
import PropTypes from "prop-types";
import Table from "react-bootstrap/Table";

import RankItem from "../RankItem";

RankList.propTypes = {
  rankList: PropTypes.array,
};

function RankList(props) {
  const { rankList } = props;
  return (
    <Table striped hover className="rankList">
      <thead>
        <tr className="rankList-head">
          <th className="rankList-rank">Rank</th>
          <th className="rankList-name">User Name</th>
          <th className="rankList-wpm">Wpm</th>
          <th className="rankList-age">Age</th>
        </tr>
      </thead>
      {rankList.map((item, index) => {
        return (
          <RankItem
            key={index}
            rank={index + 1}
            id={item.id}
            name={item.name}
            wpm={item.wpm}
            age={item.age}
          />
        );
      })}
    </Table>
  );
}

export default RankList;
