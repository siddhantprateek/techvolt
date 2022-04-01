import axios from "axios";
import React, { useEffect, useState } from "react";
import "./NFTbadge.styles.css";
function NFTbadge() {
  const [nfts, setNfts] = useState(null);
  useEffect(() => {
    axios
      .get("http://localhost:4000/api/topmonth")
      .then((res) => {
        console.log(res.data);
        setNfts(res.data);
      })
      .catch((err) => console.error(err));
  }, []);
  return (
    <div className="nft-badge-container">
      {nfts?.map((nft, key) => (
        <a href={nft.collection_url} target="__blank">
          <div className="top-nft-badge">
            <div>
              <h1>#{key + 1}</h1>
            </div>
            <div className="nft-details">
              <p>{nft.nft_name.slice(0, 24)}</p>
              <b>{nft.price}</b>
              <p>{nft.date}</p>
            </div>
          </div>
        </a>
      ))}
    </div>
  );
}

export default NFTbadge;
