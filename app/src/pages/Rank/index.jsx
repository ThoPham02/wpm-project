import React from "react";
import RankList from "./components/RankList";

function Rank() {
  const rankList = [
    {
      id: 1,
      name: "Tho Pham",
      wpm: 40,
      age: "20:20 11/01/2023",
    },
  ];
  return (
    <div className="rank">
      <div className="rank-title">
        <h2>Rank</h2>
      </div>
      <RankList rankList={rankList} />
    </div>
  );
}

export default Rank;
