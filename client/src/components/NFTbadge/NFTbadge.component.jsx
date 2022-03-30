import axios from 'axios'
import React, {useEffect, useState} from 'react'
function NFTbadge() {
  const [ nfts, setNfts ] = useState(null) 
  useEffect(() => {
    
    axios.get("http://localhost:4000/api/topmonth")
    .then((res) => {
        console.log(res.data)
        setNfts(res.data)
    })
    .catch((err) => console.error(err))
  }, [])
  return (
    <div>
        {nfts?.map((nft) => (
            <>
                <p>{nft.nft_name}</p>
                <a href={nft.collection_url}>{nft.collection_url}</a>
            </>
        ))}
    </div>
  )
}

export default NFTbadge